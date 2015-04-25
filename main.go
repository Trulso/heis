package main



import (
	"fmt"
	io "./driver"
	"./queue"
	"./stateMachine"
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
	fmt.Println("Har opprettet alle channels")
	go queue.Init(upOrdersChan,downOrdersChan,commandOrdersChan,orderOnSameFloorChan,orderInEmptyQueueChan)
	go stateMachine.Init(floorReachedChan,orderOnSameFloorChan,orderInEmptyQueueChan)
	fmt.Println("Laget en routine med queue og FSM")
	go io.OrderButtonPolling(commandOrdersChan,upOrdersChan, downOrdersChan)
	go io.FloorSensorPolling(floorReachedChan)
	fmt.Println("Laget en routine med bolling")

	dontEndChan := make(chan int)
	<-dontEndChan
	fmt.Println("Dette bor ikke skrives ut")
}

https://github.com/Trulso/heis.git