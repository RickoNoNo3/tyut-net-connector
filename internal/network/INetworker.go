package network

import "net"

type INetworker interface {
	getCampusNetwork() (net.IP, net.HardwareAddr, error)
	CheckCampus() error
	CheckInternet() error
	Connect() error
}
