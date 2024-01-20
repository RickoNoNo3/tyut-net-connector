package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/rickonono3/tyut-net-connector/internal/config"
	"github.com/rickonono3/tyut-net-connector/internal/network"
	"github.com/rickonono3/tyut-net-connector/internal/silentstart"
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
			config.C["mode"] =*mode
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
			fmt.Println("ERR:", pan)
			fmt.Println("+--------------")
			showHelp()
			os.Exit(-1)
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
	var n network.INetworker
	n = &network.Networker{}
	fmt.Println("[CHECK] Starting check")
	if n.CheckCampus() && n.CheckInternet() {
		fmt.Println("[INFO] Already logged in")
	}
	for {
		// 如果能同时连接到校园网和公网，按10s间隔重新检测
		// 如果能连接到校园网，但连接不到公网，登录一次，若成功，按3600s间隔重新检测，若失败，按120s间隔重新检测并登录
		// 如果不能连接到校园网，按10s间隔重新检测
		var sleepTime int64
		fmt.Println("[CHECK] Check campus network")
		if n.CheckCampus() {
			fmt.Println("[CHECK] Check internet network")
			if n.CheckInternet() {
				fmt.Println("[WAIT] 10s: Fully connected")
				sleepTime = 10
			} else {
				fmt.Println("[LOGIN] As:", config.C["username"])
				if err := n.Connect(); err == nil {
					fmt.Println("[WAIT] 10s: Logged In")
					sleepTime = 3600
				} else {
					fmt.Println("[INFO] Cannot login:", err)
					fmt.Println("[WAIT] 120s: Cannot login")
					sleepTime = 120
				}
			}
		} else {
			fmt.Println("[WAIT] 10s: No campus network")
			sleepTime = 10
		}
		time.Sleep(time.Duration(sleepTime * int64(time.Second)))
	}
}
