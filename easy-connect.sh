#!/bin/bash

if [ -z "$1" -o -z "$2" ]; then
    echo "Usage: $0 <username> <password>"
    exit 1
fi

chmod +x ./encrypt.sh
username=$(./encrypt.sh $1)
password=$(./encrypt.sh $2)
ip=$(ifconfig | sed -En 's/127.0.0.1//;s/.*inet (addr:)?(101\.([0-9]*\.){2}[0-9]*).*/\2/p' | head -n 1)
if [ -z "$ip" ]; then
    echo "Failed to get campus network IP"
    exit 1
else
    echo "IP: $ip"
fi
ip=$(./encrypt.sh $ip)

loginUrl="https://drcom.tyut.edu.cn:802/eportal/portal/login?callback=130546474445&login_method=46&user_account=$username&user_password=$password&wlan_user_ip=$ip&wlan_user_ipv6=&wlan_user_mac=474747474747474747474747&wlan_ac_ip=&wlan_ac_name=&mac_type=47&authex_enable=&jsVersion=435944&web=47&terminal_type=46&lang=0d1f5a1419&enable_r3=47&encrypt=1&v=8816",

wget -qO- $loginUrl
echo