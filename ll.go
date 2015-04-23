package main
	

import (
	."fmt"
	."./network"
	//."time"
)

func main() {
	
	myc := make (chan []byte)


	go UDPRx(myc,30003)


	for{
		rx := make([]byte,1024)
		rx  = <- myc
		Printf(string(rx))
	}

}

