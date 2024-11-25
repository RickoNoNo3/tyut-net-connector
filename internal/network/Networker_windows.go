//go:build windows

package network

import (
	"errors"
	"fmt"
	"net"
	"reflect"
	"syscall"

	"github.com/rickonono3/tyut-net-connector/internal/config"
	"github.com/rickonono3/tyut-net-connector/internal/network/rasapi"
)

func (n *Networker) ConnectByPPP(i *net.Interface) error {
	return rasapiConnect(i)
}

func rasapiConnect(i *net.Interface) (err error) {
	// recover any panic
	defer func() {
		if r := recover(); r != nil {
			if reflect.TypeOf(r) == reflect.TypeOf(err) {
				err = errors.New("PPP connection failed: " + r.(error).Error())
			}
			err = errors.New("PPP connection failed")
		}
	}()
	var ret uint32
	_, _, err = rasapi.RasSetEntryPropertiesW("tyut.ppp0", i.Name)
	// try to create tyut.ppp0 connection, skip if exists
	if ret = uint32((interface{})(err).(syscall.Errno)); ret == 0 {
		fmt.Println("[PPPoE] Connection created, connecting...")
	} else if ret == 183 {
		fmt.Println("[PPPoE] Connection exists, connecting...")
	} else {
		return err
	}
	// connect to tyut.ppp0 using configured username and password
	_, _, _, err = rasapi.RasDialW("", "tyut.ppp0", config.C["username"], config.C["password"])
	// clean syscall errors if success
	if ret = uint32((interface{})(err).(syscall.Errno)); ret == 0 {
		err = nil
	}
	return
}
