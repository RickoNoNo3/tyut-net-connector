# [TYUT-Net-Connector](https://github.com/rickonono3/tyut-net-connector)
本程序用于：以校内网线连接到校园网但需要网页登录，**由本程序自动完成其中的网页登录步骤**

## 食用方法
### 获取
在[Release页](https://github.com/RickoNoNo3/tyut-net-connector/releases)下载最新发行版本。

*手工编译：Go 1.17 下可从源代码编译，执行根目录下的 `build-unix.sh` / `build.ps1`，将在 `build` 文件夹中生成编译完成的可执行程序及配套启动脚本*

### 连接
本程序为命令行程序，**已经为不同系统配备了启动脚本，可以编辑脚本文件，配置好个人连接账号和密码后，直接运行脚本**。

也可手工输入命令启动，手工启动方式如下（以Windows为例）：

```bash
./tyut-net-connector.exe -u <username> -p <password>
```

Windows下，需要静默启动（不显示命令行窗口）可以添加参数`silent`：

```bash
./tyut-net-connector.exe -u <username> -p <password> -silent
```

### 部署自启动
本程序设计为守护程序，轮询检查并自动维护网络状态，因此不要求系统开机自动联网，也可以在系统从睡眠、休眠等状态还原后自动恢复连接。

根据具体操作系统的不同，自启动的部署方式有所区别，这里给出一些建议：
- Linux系统：Debian/Ubuntu建议以systemctl服务形式部署，放置在network服务之后；其他linux系统类似
- Windows系统：建议将本程序的`tyut-net-connector-startup.cmd`脚本的**快捷方式**放置于`开始菜单-启动`目录下，注意提前编辑好脚本文件中的账号密码，并将silent置为1
- MacOS系统：建议在`设置-用户与群组-登录项`中配置启动脚本`tyut-net-connector-startup.sh`为登录项

### 详细参数
本程序为命令行程序，支持以`-key value`格式传递参数，其中`u`和`p`是必传参数：

| 参数   | 值     | 功能描述                         |
| ------ | ------ | -------------------------------- |
| u      | string | 你的账号, 无默认值               |
| p      | string | 你的密码, 无默认值               |
| silent | 无     | 静默启动（常驻后台），默认不启用 |

### TODO
- [ ] 单进程约束
- [ ] 支持从外部网连接（MotionPro）
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
