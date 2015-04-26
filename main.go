package main

import (
	"./driver"
	"./network"
	"./queue"
	"./stateMachine"
	. "./struct"
	"fmt"
)

func main() {
	fmt.Println("Da starter vi.")
	driver.Init()
	queue.Init()

	//Ordner med bestillinger mot hardware
	upOrdersChan := make(chan int)
	downOrdersChan := make(chan int)
	commandOrdersChan := make(chan int)
	orderOnSameFloorChan := make(chan int)
	orderInEmptyQueueChan := make(chan int)
	go queue.OrderButtonHandler(upOrdersChan, downOrdersChan, commandOrdersChan, orderOnSameFloorChan, orderInEmptyQueueChan)
	go driver.OrderButtonPolling(commandOrdersChan, upOrdersChan, downOrdersChan)

	//Ordner med etasjesensor
	floorReachedChan := make(chan int)
	go driver.FloorSensorPolling(floorReachedChan)

	//Starter tilstandsmaskin
	go stateMachine.Init(floorReachedChan, orderOnSameFloorChan, orderInEmptyQueueChan)

	//Ordner Heartbeat
	newElevatorChan := make(chan string)
	deadElevatorChan := make(chan string)
	go network.HeartbeatTransceiver(newElevatorChan, deadElevatorChan)
	go queue.HeartbeatReceiver(newElevatorChan, deadElevatorChan)

	//Ordner med beskjeer mellom heisene
	receiveChan := make(chan Message)
	go network.MessageTransceiver(receiveChan)
	go queue.MessageReceiver(receiveChan, orderOnSameFloorChan, orderInEmptyQueueChan)

	dontEndChan := make(chan int)
	<-dontEndChan
}
