package main



import (
	"fmt"
	io "./driver"
	"./queue"
	"./stateMachine"
	"./network"
	."./struct"
)


func main() {
	fmt.Println("Da starter vi.")
	io.Init()
	floorReachedChan := make(chan int)
	upOrdersChan := make(chan int)
	downOrdersChan := make(chan int)
	commandOrdersChan := make(chan int)
	orderOnSameFloorChan := make(chan int)
	orderInEmptyQueueChan := make(chan int)
	newElevator	 := make(chan string)
	deadElevator := make(chan string)
	toPass := make(chan Message)
	toGet := make(chan Message)
	fmt.Println("Har opprettet alle channels")


	go network.HeartbeatTransceiver(newElevator,deadElevator)
	go network.StatusTransceiver(toPass,toGet)



	go queue.Init(upOrdersChan,downOrdersChan,commandOrdersChan,orderOnSameFloorChan,orderInEmptyQueueChan)
	go stateMachine.Init(floorReachedChan,orderOnSameFloorChan,orderInEmptyQueueChan)
	go queue.StatusDecoder(upOrdersChan,downOrdersChan,toGet,toPass)
	fmt.Println("Laget en goroutine med queue og FSM")


	go io.OrderButtonPolling(commandOrdersChan,upOrdersChan, downOrdersChan)
	go io.FloorSensorPolling(floorReachedChan)
	fmt.Println("Laget en goroutine med bolling")

	dontEndChan := make(chan int)
	<-dontEndChan
	fmt.Println("Dette bor ikke skrives ut")
}
