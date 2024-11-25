package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"net/http"

	"github.com/rickonono3/tyut-net-connector/internal/config"
	"github.com/rickonono3/tyut-net-connector/internal/network"
	"github.com/rickonono3/tyut-net-connector/internal/silentstart"

	_ "net/http/pprof"
)

var username = flag.String("u", "", "string")
var password = flag.String("p", "", "string")
var mode = flag.String("mode", "", "string")
var silent = flag.Bool("silent", false, "silent")
var help = flag.Bool("help", false, "help")

func initConfig() {
	flag.Parse()
	if *username != "" {
		config.C["username"] = *username
	} else {
		panic("No username gotten")
	}
	if *password != "" {
		config.C["password"] = *password
	} else {
		panic("No password gotten")
	}
	if *mode != "" {
		if *mode == "direct" || *mode == "motionpro" {
			config.C["mode"] = *mode
		} else {
			panic("Invalid mode! Only supports direct/motionpro")
		}
	}
}

func showHelp() {
	fmt.Println(`HELP:
-u\tusername
-p\tpassword

... see README.md`)
}

func silentStart() {
	silentstart.SilentStart()
}

func main() {
	defer func() {
		if pan := recover(); pan != nil {
			fmt.Println("+--------------")
			fmt.Println("PANIC:", pan)
			fmt.Println("+--------------")
			panic(pan)
		}
	}()
	initConfig()
	if *help {
		showHelp()
		os.Exit(0)
	}
	if *silent {
		silentStart()
	}
	var n network.INetworker = &network.Networker{}
	fmt.Println("[CHECK] Starting check")
	if n.CheckCampus() == nil && n.CheckInternet() == nil {
		fmt.Println("[INFO] Already logged in")
	}
	go func() {
		http.ListenAndServe("0.0.0.0:13880", nil)
	}()
	var waitCount int32
	for {
		// 如果能同时连接到校园网和公网，按10s间隔重新检测
		// 如果能连接到校园网，但连接不到公网，登录一次，若成功，按3600s间隔重新检测，若失败，按120s间隔重新检测并登录
		// 如果不能连接到校园网，按10s间隔重新检测
		var sleepTime int64
		// fmt.Println("[CHECK] Check campus network")
		if campusErr := n.CheckCampus(); campusErr == nil {
			// fmt.Println("[CHECK] Check internet network")
			if internetErr := n.CheckInternet(); internetErr == nil {
				if waitCount == 0 {
					fmt.Println("[WAIT] Fully connected, waiting...")
				}
				waitCount = (waitCount + 1) % 10
				sleepTime = 10
			} else {
				fmt.Println("[LOGIN] As:", config.C["username"])
				if connectErr := n.Connect(); connectErr == nil {
					waitCount = 0
					fmt.Println("[LOGIN] Logged In")
					fmt.Println("[WAIT] 600s")
					sleepTime = 600
				} else {
					waitCount = 0
					fmt.Println("[LOGIN] Cannot login:", connectErr)
					fmt.Println("[WAIT] 120s")
					sleepTime = 120
				}
			}
		} else {
			if waitCount == 0 {
				fmt.Println("[LOGIN] No campus network (" + campusErr.Error() + "), just try to connect...")
				if connectErr := n.Connect(); connectErr == nil {
					waitCount = 0
					fmt.Println("[LOGIN] Logged In")
					fmt.Println("[WAIT] 600s")
					sleepTime = 600
				} else {
					waitCount = (waitCount + 1) % 10
				}
			} else {
				waitCount = (waitCount + 1) % 10
			}
			sleepTime = 10
		}
		time.Sleep(time.Duration(sleepTime * int64(time.Second)))
	}
}
