package queue

import (
//"fmt"
//io "../driver"
//io "../fakeDriver"
)

const (
//Bør fjernes og bruke de fra driveren
	N_FLOORS = 4

	COMMAND = 2
	UP   = 1
	DOWN = -1	
// ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

	STOP = 0
)

type Message struct {
	MessageType string //neworder,just arrived, status update, completed order,
	SenderIP    string
	elevators   map[string]Elevator
}

type Order struct {
	Type  int
	Floor int
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

func Init(upOrderChan chan int,downOrderChan chan int,commandOrderChan chan int,receiveMsgChan chan int,heartbeatChan chan string,) {

	for {
		select {
		case floor := <-upOrderChan:
			newOrder = Order{UP, floor}

		case floor := <-downOrderChan:
			newOrder = Order{DOWN, floor}

		case floor := <-commandOrderChan:
			newOrder = Order{COMMAND, floor}

		}
	}
}
func isIdenticalOrder(newOrder Order) bool {
	for IP, elevator := range elevators {
		switch newOrder.Type {
		case UP:
			if elevator.UpOrders[newOrder.Floor] {
				return true
			}
		case DOWN:
			if elevator.DownOrders[newOrder.Floor] {
				return true
			}
		case COMMAND:
			if elevator.CommandOrders[newOrder.Floor] {
				return true
			}
		}
	}
	return false
}//Ferdig

func addOrder(newOrder Order) {
	if isIdenticalOrder(newOrder) {
		return
	}
defer func(){/* Send Oppdateringen til alle sammen*/}()
	if newOrder.Type == COMMAND {
		elevator[myIP].CommandOrders[newOrder.floor]
		return
	}
	cheapestElevator := myIP
	minCost := 
	for IP,Elevator := range elevators {
		 cost := costFunction(Elevator, newOrder)
		 if cost < minCost {
		 	cheapestElevator = IP
		 }
		 if cost = 0 {
		 	break
		 }
	}
	return


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
	return 1
}

func ShouldStop(floor int) bool {
	if elevators[myIP].CommandOrders[floor] || elevators[myIP].

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
}//Ferdig

func ordersBellow(IP string) bool{
	for floor := elevators[IP].LastPassedFloor - 1; floor > -1; floor-- {
		if elevators[IP].UpOrders[floor] || elevators[IP].CommandOrders[floor] || elevators[IP].DownOrders[floor] {
			return true
		}
	}
	return false
}//Ferdig

func isQueueEmpty(IP string) bool {
	if ordersAbove(IP) || ordersBellow(IP) {
		return false
	}
	floor := elevators[IP].LastPassedFloor
	if elevators[IP].UpOrders[floor] || elevators[IP].DownOrders[floor] ||  elevators[IP].CommandOrders[floor] {
		return false
	}
	return true
}//Ferdig
	