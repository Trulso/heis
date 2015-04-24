package main

import (
	//"fmt"
	"./driver"
	//"./network"
	"os"
	"net"
	"time"
)


 
func main() {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		os.Stderr.WriteString("Oops: " + err.Error() + "\n")
		os.Exit(1)
	}
 
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				os.Stdout.WriteString(ipnet.IP.String() + "\n")
			}
		}
	}
	
	driver.HwInit()
	
	time.Sleep(100*time.Millisecond)
	
	driver.SetDoorLamp(1)
	
	select{}
}
