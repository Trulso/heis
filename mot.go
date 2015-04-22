package main

import (
	"fmt"
	//"./driver"
	"./network"
)

const (

	A = iota
	B
	C
	D
)


func main(){

	
	test := make(chan network.Dummy)

	test_struckt:= network.Dummy{}


	go network.UDPListen(test,30000)

	for{
		test_struckt <- test
		fmt.Printf("test_struckt.s")
	}

}
