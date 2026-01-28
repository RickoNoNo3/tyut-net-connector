@echo off
rem 请将username和password配置为校园网账号和密码，silent为1则在后台静默启动，否则在命令行窗口中启动
SET username=你的账号
SET password=你的密码
SET silent=0

if %silent% EQU 1 (
  start tyut-net-connector.exe -u %username% -p %password% -silent
  goto END
)
tyut-net-connector.exe -u %username% -p %password%
pause >nul

:END
exit
