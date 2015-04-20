package main

import (
	//"fmt"
	"./driver"
)

const (

	A = iota
	B
	C
	D
)


func main(){
	driver.HwInit()
	
	BCmd := make(chan int)	
	BUp  := make(chan int)
	BDown:= make(chan int)

	go driver.CommandOrdersPolling(BCmd)
	go driver.DownOrdersPolling(BDown)
	go driver.UpOrdersPolling(BUp)

	for{
		driver.SetButtonLed(<-BCmd,driver.Command)
		//driver.SetButtonLed(<- BUp,driver.Up)
		//fmt.Printf("%d",<- BUp)
		//driver.SetButtonLed(<- BDown,driver.Down)

	}







}
