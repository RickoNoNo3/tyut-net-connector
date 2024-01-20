package config

var C = map[string]string{
	"username":        "",
	"password":        "",
	"mode":            "direct",
	"health_campus":   "http://219.226.127.250:801/",
	"health_internet": "https://www.baidu.com/",
	"login":           "http://219.226.127.250:801/eportal/portal/login?callback=dr1003&login_method=1&user_account=%2C0%2C{{.Username}}&user_password={{.Password}}&wlan_user_ip={{.IP}}&wlan_user_ipv6=&wlan_user_mac={{.MAC}}&wlan_ac_ip=&wlan_ac_name=&jsVersion=4.1.3&terminal_type=1&lang=zh-cn&v=6688&lang=zh",
}
