package config

var C = map[string]string{
	"username":        "",
	"password":        "",
	"mode":            "direct",
	"health_campus":   "https://drcom.tyut.edu.cn/",
	"health_internet": "https://www.baidu.com/",
	// "login":           "https://drcom.tyut.edu.cn:802/eportal/portal/login?callback=dr1003&login_method=1&user_account=%2C0%2C{{.Username}}&user_password={{.Password}}&wlan_user_ip={{.IP}}&wlan_user_ipv6=&wlan_user_mac={{.MAC}}&wlan_ac_ip=&wlan_ac_name=&jsVersion=4.1.3&terminal_type=1&lang=zh-cn&v=6688&lang=zh",
	// 升级后的
	"login": "https://drcom.tyut.edu.cn:802/eportal/portal/login?callback=130546474445&login_method=46&user_account={{.UsernameEncrypted}}&user_password={{.PasswordEncrypted}}&wlan_user_ip={{.IPEncrypted}}&wlan_user_ipv6=&wlan_user_mac=474747474747474747474747&wlan_ac_ip=&wlan_ac_name=&mac_type=47&authex_enable=&jsVersion=435944&web=47&terminal_type=46&lang=0d1f5a1419&user_agent=3a180d1e1b1b1658425947575f201e191318000457392357464759474c57201e1941434c570f41435e573607071b122012153c1e0358424440594441575f3c3f233a3b5b571b1e1c12573012141c185e57341f05181a125846454e5947594759475724161116051e58424440594441&enable_r3=47&encrypt=1&v=8816&lang=zh",
}
