# TYUT-Net-Connector
本项目用于：以校内网线连接到校园网但需要网页登录，**自动完成其中的网页登录步骤**

## 食用方法
### 获取
在[Release页]()下载最新发行版本。

*对于须手工编译的系统：Go 1.17 下执行`go install github.com/rickonono3/tyut-net-connector`，此时可执行程序将自动加入PATH*

### 连接
本程序为命令行程序，启动方式如下（以Windows为例）：

```bash
./tyut-net-connector.exe -u <username> -p <password>
```

需要静默启动（不显示命令行窗口）可以添加参数`silent`，注意不要重复启动：

```bash
./tyut-net-connector.exe -u <username> -p <password> -silent
```

### 部署自启动
本程序设计为守护程序，轮询检查并自动维护网络状态，因此不要求系统开机自动联网，也可以在系统从睡眠、休眠等状态还原后自动恢复连接。

根据具体操作系统的不同，自启动的部署方式有所区别，这里给出一些建议：
- Linux系统：Debian/Ubuntu建议以systemctl服务形式部署，放置在network服务之后；其他linux系统类似
- Windows系统：建议将本程序的快捷方式放置于`开始菜单-启动`目录下，并在快捷方式属性中的目标上添加参数`-u <你的账号> -p <你的密码> -silent`；也可选择使用计划任务配置
- MacOS系统：建议手工创建启动脚本，然后在`设置-用户与群组-登录项`中配置其为登录项

### 详细参数
本程序为命令行程序，支持以`-key value`格式传递参数，其中`u`和`p`是必传参数：

| 参数   | 值     | 功能描述                         |
| ------ | ------ | -------------------------------- |
| u      | string | 你的账号, 无默认值               |
| p      | string | 你的密码, 无默认值               |
| silent | 无     | 静默启动（常驻后台），默认不启用 |
<!-- 
| mode | 从校内还是校外连接，可选`direct`或`motionpro`, 默认值`direct` |
-->

<!--
从公共网络通过MotionPro连接到校园网，**自动完成MotionPro的初始化和连接步骤**


### 校内连接
Linux/MacOS(Shell):
```bash
./tyut-net-connector -u <username> -p <password> [OPTIONS]
```

Windows类似。
### 校外连接
Linux/MacOS(Shell):
```bash
./tyut-net-connector -u <username> -p <password> -mode motionpro [OPTIONS]
```

Windows类似。


## 注意
- 若从公共网络使用MotionPro连接，请确保网络环境干净畅通，连接校园网后启动其他VPN或虚拟网卡或在其他VPN或虚拟网卡之上连接校园网会发生不可预料的问题。使用校内网络无此问题。
- 已经安装过MotionPro的系统，如在连接时发生问题，可选择在运行本程序前卸载原来的版本

-->