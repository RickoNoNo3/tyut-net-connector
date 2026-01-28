//go:build windows
// +build windows

package silentstart

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
	"github.com/gonutz/w32/v2"
	"github.com/rickonono3/tyut-net-connector/internal/config"
)

var (
	kernel32 = syscall.NewLazyDLL("kernel32.dll")
	procFreeConsole = kernel32.NewProc("FreeConsole")
)

func checkPowerShell() bool {
	cmd := exec.Command("powershell", "-command", "exit 0")
	err := cmd.Run()
	return err == nil
}

func SilentStart() {
	if checkPowerShell() {
		// 使用 PowerShell 启动隐藏的新进程
		exe, err := os.Executable()
		if err != nil {
			exe = os.Args[0] // 回退到 os.Args[0]
		}
		cmd := exec.Command("powershell", "-windowstyle", "hidden", "-command",
			fmt.Sprintf("Start-Process '%s' -ArgumentList '-u %s -p %s -silent -alreadysilent' -WindowStyle Hidden", exe, config.C["username"], config.C["password"]),
		)
		cmd.Start()
		os.Exit(0)
	} else {
		// 回退到原来的隐藏逻辑
		console := w32.GetConsoleWindow()
		if console != 0 {
			_, consoleProcID := w32.GetWindowThreadProcessId(console)
			currentProcID := w32.GetCurrentProcessId()
			fmt.Printf("[DEBUG] Console ProcID: %d, Current ProcID: %d\n", consoleProcID, currentProcID)
			if currentProcID == consoleProcID {
				_ = w32.ShowWindowAsync(console, w32.SW_HIDE)
			} else {
	            procFreeConsole.Call()
			}
		} else {
			procFreeConsole.Call()
		}
	}
}
