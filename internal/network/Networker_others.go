//go:build !windows

package network

import (
	"errors"
	"net"
)

func (n *Networker) ConnectByPPP(i *net.Interface) error {
	return errors.New("PPP connection is Not implemented on this OS")
}
