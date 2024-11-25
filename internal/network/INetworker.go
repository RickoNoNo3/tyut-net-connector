package network

import "net"

type INetworker interface {
	getCampusNetwork() (net.IP, net.HardwareAddr, *net.Interface, error)
	CheckCampus() error
	CheckInternet() error
	ConnectByUrl(ip net.IP, mac net.HardwareAddr) error
	ConnectByPPP(i *net.Interface) error
	Connect() error
}
