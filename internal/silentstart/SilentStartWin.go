//go:build windows
// +build windows

package silentstart

import (
	"fmt"
	"syscall"
	"github.com/gonutz/w32/v2"
)

var (
	kernel32 = syscall.NewLazyDLL("kernel32.dll")
	procFreeConsole = kernel32.NewProc("FreeConsole")
)

func freeConsole() error {
	ret, _, err := procFreeConsole.Call()
	if ret == 0 {
		return err
	}
	return nil
}

func SilentStart() {
	console := w32.GetConsoleWindow()
	fmt.Printf("[DEBUG] Console handle: %d\n", console)
	if console != 0 {
		_, consoleProcID := w32.GetWindowThreadProcessId(console)
		currentProcID := w32.GetCurrentProcessId()
		fmt.Printf("[DEBUG] Console ProcID: %d, Current ProcID: %d\n", consoleProcID, currentProcID)
		if currentProcID == consoleProcID {
			result := w32.ShowWindowAsync(console, w32.SW_HIDE)
			fmt.Printf("[DEBUG] ShowWindowAsync result: %v\n", result)
		} else {
			fmt.Printf("[DEBUG] Console not owned by current process, trying FreeConsole\n")
			if err := freeConsole(); err != nil {
				fmt.Printf("[DEBUG] FreeConsole failed: %v\n", err)
			} else {
				fmt.Printf("[DEBUG] FreeConsole succeeded\n")
			}
		}
	} else {
		fmt.Printf("[DEBUG] No console window found, trying FreeConsole\n")
		if err := freeConsole(); err != nil {
			fmt.Printf("[DEBUG] FreeConsole failed: %v\n", err)
		} else {
			fmt.Printf("[DEBUG] FreeConsole succeeded\n")
		}
	}
}
