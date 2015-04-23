package queue

import (
//"fmt"
//io "../driver"
//io "../fakeDriver"
)

//struct order
//struct elevator
//orders= []order
//elevators= []elevator
const (
	N_FLOORS = 4

	UP   = 1
	STOP = 0
	DOWN = -1
)

type Message struct {
	MessageType string //neworder,just arrived, status update, completed order,
	SenderIP    string
	elevators   map[string]Elevator
}

type Order struct {
	Direction int
	Floor     int
}

type Elevator struct {
	Direction       int
	LastPassedFloor int
	UpOrders        []bool
	DownOrders      []bool
	CommandOrders   []bool
}



var myIP = string
var elevators = make(map[string]*Elevator)

//Channels
/*
Oppbestillingsknapper
Nedbestillingsknapper
Commandbestillingsknapper
Nybestilling
Statusoppdatering
*/
func Init(
	upOrderChan chan int,
	downOrderChan chan int,
	commandOrderChan chan int,
	receiveMsgChan chan int,
) {

	for {
		select {
		case floor := <-upOrderChan:
			newOrder = Order{UP, floor}

		case floor := <-downOrderChan:

		case floor := <-commandOrderChan:

		}
	}
}
func isIdenticalOrder(newOrder Order) {
	for IP, elevator := range elevators{
		if 
	}
}

func cheapestElevator() string {
	return myIP
	//TODO: Bruker costFunction til å finne den billigste heisen
}

func costFunction(elevator Elevator, newOrder Order) int {
	cost := 0
	difference := elevator.LastPassedFloor - newOrder.floor
	if elevator.Direction == STOP {
		cost = cost + 1*difference
		return cost
	} else if elevator.Direction == UP {
		if difference > 0 {
		}
	}
	return cost
}

func ShouldStop(floor int) bool {
	return true
	//TODO: sjekke egen bestillingliste
}

func AddElevator(newElevator Elevator, IP string) {
	elevators[IP] = &newElevator
}

func OrderCompleted(floor int) {
	for IP, elevator := range elevators {
		elevator.UpOrders[floor] = false
		elevator.DownOrders[floor] = false
		if myIP == IP {
			elevator.CommandOrders[floor] = false
		}
	}
	//TODO: send en beskjed til netverket om at en ordre har blitt fjernet.
}

func NextDirection() int {
	if elevators[myIP].Direction == UP {
		if ordersAbove() {
			return UP
		}else if ordersBellow() {
			return DOWN
		}
	} else if elevators[myIP].Direction == DOWN {
		if ordersBellow() {
			return DOWN
		}else if ordersAbove() {
			return UP
		}
	}
}

func ordersAbove(IP string) bool{
	for floor := elevators[IP].LastPassedFloor + 1; floor < N_FLOORS; floor++ {
		if elevators[IP].UpOrders[floor] || elevators[IP].CommandOrders || elevators[IP].DownOrders{
			return true
		}
	}
	return false
}

func ordersBellow(IP string) bool{
	for floor := elevators[IP].LastPassedFloor - 1; floor > -1; floor-- {
		if elevators[IP].UpOrders[floor] || elevators[IP].CommandOrders[floor] || elevators[IP].DownOrders[floor] {
			return true
		}
	}
	return false
}

func isQueueEmpty(IP string) bool {
	if ordersAbove(IP) || ordersBellow(IP) {
		return false
	}
	floor := elevators[IP].LastPassedFloor
	if elevators[IP].UpOrders[floor] || elevators[IP].DownOrders[floor] ||  elevators[IP].CommandOrders[floor] {
		return false
	}
	return true
}
	