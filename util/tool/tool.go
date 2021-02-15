package tool

import (
	"net"
	"time"
)

var IP string
var TimeLocal *time.Location

func init() {
	ipAcquire()
	TimeLocal, _ = time.LoadLocation("Asia/Shanghai")
}

func ipAcquire() {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		panic(err)
	}
	for _, value := range addrs {
		if ipnet, ok := value.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				IP = ipnet.IP.String()
			}
		}
	}
}
