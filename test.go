package main

import (
	//"fmt"
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

	test_struckt:= network.Dummy{"hallo"}
	

	test <- test_struckt

	network.UDPSend(test,30000)



}
