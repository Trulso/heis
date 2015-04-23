package main

import (
	."fmt"
	."./network"
	."time"
)




func main(){


	myc := make (chan []byte)


	send := make([]byte,1024)
	go UDPTx(myc,30003)
	
	for{
		send = []byte("Hallo i verden!")
		myc <- send
		Sleep(2*Second)
		Printf("Sendt")
	}

}