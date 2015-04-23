package main

import (
	."fmt"
	."./network"
	//."time"
)

type myTest struct {
		Id int
		Tx string
}

func dummy(rx chan myTest){

	for{
		
		To_chan := myTest{50,"hello"}
		rx <- To_chan

	}

}

func main(){


	myc := make (chan myTest)


	go UDPSend(myc,30000)
	
	for{
		rx := <- myc
		Printf(rx.Tx)
	}
}