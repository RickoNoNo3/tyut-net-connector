package network

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net"
	"net/http"
	"regexp"
	"strings"
	"text/template"

	"github.com/rickonono3/m2obj"
	"github.com/rickonono3/tyut-net-connector/internal/config"
)

type Networker struct {
}

var loginResponseRegexp = regexp.MustCompile("^[A-Za-z0-9]+\\((\\{.*\\})\\)[^{}]*$")

func (n *Networker) getCampusNetwork() (ip net.IP, mac net.HardwareAddr, err error) {
	var interfaces []net.Interface
	if interfaces, err = net.Interfaces(); err != nil {
		return nil, nil, err
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
				}
			}
		}
	}
	if ip != nil && mac != nil {
		err = nil
	}
	return
}

func (n *Networker) check(url string) bool {
	// 发送请求
	resp, err := http.Get(url)
	if err != nil {
		return false
	}
	// 读取返回结果状态码
	return resp.StatusCode >= 200 && resp.StatusCode < 300
}

func (n *Networker) CheckCampus() bool {
	return n.check(config.C["health_campus"])
}

func (n *Networker) CheckInternet() bool {
	return n.check(config.C["health_internet"])
}

func (n *Networker) Connect() error {
	// 处理loginUrl template
	var (
		loginTemplateStr, loginUrl string
		loginTemplate              *template.Template
		ip                         net.IP
		mac                        net.HardwareAddr
		err                        error
	)
	loginTemplateStr = config.C["login"]
	if loginTemplate, err = template.New("login").Parse(loginTemplateStr); err != nil {
		return err
	}
	if ip, mac, err = n.getCampusNetwork(); err != nil {
		return err
	}
	buf := &bytes.Buffer{}
	loginTemplate.Execute(buf, m2obj.New(m2obj.Group{
		"Username": config.C["username"],
		"Password": config.C["password"],
		"IP":       ip.String(),
		"MAC":      mac.String(),
	}).Staticize())
	loginUrl = buf.String()
	// 发送login请求
	if req, err := http.NewRequest("GET", loginUrl, nil); err != nil {
		return err
	} else if resp, err := http.DefaultClient.Do(req); err != nil {
		return err
	} else {
		// 处理响应并Print结果
		if resp.StatusCode >= 200 && resp.StatusCode < 300 {
			if body, err := io.ReadAll(resp.Body); err == nil {
				resMatch := loginResponseRegexp.FindSubmatch(body)
				if resMatch != nil && len(resMatch) > 1 && len(resMatch[1]) >= 2 {
					resJson := resMatch[1]
					res := map[string]interface{}{}
					if err = json.Unmarshal(resJson, &res); err == nil {
						if res["result"] == 1 {
							return nil
						} else {
							return errors.New(string(resJson))
						}
					} else {
						return err
					}
				} else {
					return errors.New("Invalid Response From Server")
				}
			} else {
				return err
			}
		} else {
			return errors.New("Server Status Err" + resp.Status)
		}
	}
}
