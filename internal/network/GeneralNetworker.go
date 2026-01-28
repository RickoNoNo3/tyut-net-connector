package network

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	// "fmt"
	"io"
	"net"
	"net/http"
	"regexp"
	"strings"
	"text/template"

	"github.com/rickonono3/tyut-net-connector/internal/config"
)

type Networker struct {
}

var loginResponseRegexp = regexp.MustCompile(`^[A-Za-z0-9]+\((\{.*\})\)[^{}]*$`)

func (n *Networker) getCampusNetwork() (ip net.IP, mac net.HardwareAddr, I *net.Interface, err error) {
	var interfaces []net.Interface
	if interfaces, err = net.Interfaces(); err != nil {
		return nil, nil, nil, err
	}
	for _, i := range interfaces {
		var addrs []net.Addr
		if addrs, err = i.Addrs(); err != nil {
			continue
		}
		for _, addr := range addrs {
			ad := addr.String()
			slash := strings.Index(ad, "/")
			if slash != -1 {
				ad = ad[:slash]
			}
			if strings.Count(ad, ".") == 3 {
				if strings.HasPrefix(ad, "101.") {
					ip, mac = net.ParseIP(ad), i.HardwareAddr
					I = &i
				}
			}
		}
	}
	if ip == nil || mac == nil || I == nil {
		err = errors.New("campus network not found")
	}
	return
}

func (n *Networker) check(url string) error {
	// 发送请求
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: nil,
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
		Timeout: 8 * time.Second,
	}
	resp, err := client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	// 读取返回结果状态码
	if resp.StatusCode >= 200 && resp.StatusCode < 400 {
		return nil
	} else {
		return errors.New("Error: " + resp.Status)
	}
}

func (n *Networker) CheckCampus() error {
	return n.check(config.C["health_campus"])
}

func (n *Networker) CheckInternet() error {
	return n.check(config.C["health_internet"])
}

func (n *Networker) ConnectByUrl(ip net.IP, mac net.HardwareAddr) error {
	// 处理loginUrl template
	var (
		loginTemplateStr, loginUrl string
		loginTemplate              *template.Template
		err                        error
	)
	loginTemplateStr = config.C["login"]
	if loginTemplate, err = template.New("login").Parse(loginTemplateStr); err != nil {
		return err
	}
	buf := &bytes.Buffer{}
	loginTemplate.Execute(buf, map[string]interface{}{
		"Username":          config.C["username"],
		"Password":          config.C["password"],
		"IP":                ip.String(),
		"MAC":               mac.String(),
		"UsernameEncrypted": entrypt(config.C["username"]),
		"PasswordEncrypted": entrypt(config.C["password"]),
		"IPEncrypted":       entrypt(ip.String()),
		"MACEncrypted":      entrypt(mac.String()),
	})
	loginUrl = buf.String()

	// 发送login请求
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: nil,
		},
	}
	if req, err := http.NewRequest("GET", loginUrl, nil); err != nil {
		return err
	} else if resp, err := client.Do(req); err != nil {
		if resp != nil {
			body, _ := io.ReadAll(resp.Body)
			fmt.Println(string(body))
			resp.Body.Close()
		}
		return err
	} else {
		defer resp.Body.Close()
		// 处理响应并Print结果
		if resp.StatusCode >= 200 && resp.StatusCode < 300 {
			if body, err := io.ReadAll(resp.Body); err == nil {
				resMatch := loginResponseRegexp.FindSubmatch(body)
				if len(resMatch) > 1 && len(resMatch[1]) >= 2 {
					resJson := resMatch[1]
					res := map[string]interface{}{}
					if err = json.Unmarshal(resJson, &res); err == nil {
						if res["result"] == float64(1) {
							return nil
						} else {
							return errors.New(string(resJson))
						}
					} else {
						return err
					}
				} else {
					return errors.New("invalid response from server")
				}
			} else {
				body, _ := io.ReadAll(resp.Body)
				fmt.Println(string(body))
				return err
			}
		} else {
			body, _ := io.ReadAll(resp.Body)
			fmt.Println(string(body))
			return errors.New("Server Status Err" + resp.Status)
		}
	}
}

func (n *Networker) Connect() error {
	var (
		ip  net.IP
		mac net.HardwareAddr
		i   *net.Interface
		err error
	)
	if ip, mac, i, err = n.getCampusNetwork(); err != nil {
		return err
	}
	if err = n.ConnectByUrl(ip, mac); err != nil {
		fmt.Println("[Connect]ConnectByUrl err:", err)
		fmt.Println("[Connect]Trying to ConnectByPPP...")
		return n.ConnectByPPP(i)
	}
	return nil
}
