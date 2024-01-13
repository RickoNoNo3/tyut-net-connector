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
		config.C.Set("username", *username)
	} else {
		panic("No username gotten")
	}
	if *password != "" {
		config.C.Set("password", *password)
	} else {
		panic("No password gotten")
	}
	if *mode != "" {
		if *mode == "direct" || *mode == "motionpro" {
			config.C.Set("mode", *mode)
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
	if n.CheckCampus() && n.CheckInternet() {
		fmt.Println("Already logged in")
	}
	for {
		// 如果能同时连接到校园网和公网，按120s间隔重新检测
		// 如果能连接到校园网，但连接不到公网，登录一次，若成功，按3600s间隔重新检测，若失败，按120s间隔重新检测并登录
		// 如果不能连接到校园网，按5s间隔重新检测
		var sleepTime int64
		if n.CheckCampus() {
			if n.CheckInternet() {
				sleepTime = 120
			} else {
				if err := n.Connect(); err == nil {
					sleepTime = 3600
				} else {
					sleepTime = 120
				}
			}
		} else {
			sleepTime = 5
		}
		time.Sleep(time.Duration(sleepTime * int64(time.Second)))
	}
}
