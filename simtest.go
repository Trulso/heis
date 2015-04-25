package main

import (
	"fmt"
	"./driver"
	"time"
)



func main(){
	driver.Init()

	fmt.Print("hei\n")
	driver.SetMotorDir(0)
	for floor := 0; floor<4; floor++{
		for button:=-1; button<3;button++{
			if button == 0{
				continue
			}
			driver.SetButtonLed(floor,button)
			time.Sleep(500*time.Millisecond)
		}
	}
	for floor := 0; floor<4; floor++{
		for button:=-1; button<3;button++{
			if button == 0{
				continue
			}
			driver.ClearButtonLed(floor,button)
			time.Sleep(500*time.Millisecond)
		}
	}


	upOrderChan := make(chan int)
	downOrderChan := make(chan int)
	commandOrderChan := make(chan int)
	//floorChan := make(chan int)
	go driver.OrderButtonPolling(commandOrderChan,upOrderChan,downOrderChan)
	//go driver.FloorSensorPolling(floorChan)
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
			// case floor := <-floorChan:
			// 	fmt.Printf("Vi er na i etg %d\n", floor)
			// 	if  floor == 3 {
			// 		fmt.Println("Burde kjore nedover")
			// 	}else if floor == 0 {
			// 		fmt.Println("Burde kjore oppover")
			// 	}
			// 	fmt.Println("Kommer vi hit?")
		}
		fmt.Println("Kommer vi hit?")
	}
}