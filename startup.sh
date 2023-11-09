#!/bin/bash
username=你的账号
password=你的密码

ls ./tyut-net-connector
if [[ $? -ne 0 ]]; then
  echo No executable found.
fi

chmod +x ./tyut-net-connector

./tyut-net-connector -u %username% -p %password%
