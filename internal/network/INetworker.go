package network

import "net"

type INetworker interface {
	getCampusNetwork() (net.IP, net.HardwareAddr, error)
	Check() bool
	Connect()
}
