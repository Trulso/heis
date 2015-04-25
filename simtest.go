package main

import (
	"fmt"
	"./driver"
	//"time"
)



func main(){
	driver.Init()

	fmt.Print("hei\n")

	upOrderChan := make(chan int)
	downOrderChan := make(chan int)
	commandOrderChan := make(chan int)
	floorChan := make(chan int)
	go driver.OrderButtonPolling(commandOrderChan,upOrderChan,downOrderChan)
	go driver.FloorSensorPolling(floorChan)
	// go driver.UpOrdersPolling(upOrderChannel)
	// go driver.DownOrdersPolling(downOrderChannel)
	// go driver.CommandOrdersPolling(commandOrderChannel)
	for{
		select {
			case floor := <-upOrderChan:
				fmt.Printf("Vi fikk en oppoverbestilling i etg %d\n", floor)
			case floor := <-downOrderChan:
				fmt.Printf("Vi fikk en nedoverbestilling i etg %d\n", floor)
			case floor := <- commandOrderChan:
				fmt.Printf("Vi fikk en commandbestilling i etg %d\n", floor)
			case floor := <-floorChan:
				fmt.Printf("Vi er nå i etg %d\n", floor)
		}
		//time.Sleep(100*time.Millisecond)
	}
}