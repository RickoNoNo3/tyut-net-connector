package network

import "net"

type INetworker interface {
	getCampusNetwork() (net.IP, net.HardwareAddr, error)
	CheckCampus() bool
	CheckInternet() bool
	Connect() error
}
