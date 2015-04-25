package queue

import (
	"fmt"
	io "../driver"
)

const (
	//Bør fjernes og bruke de fra driveren
	N_FLOORS = 4

	COMMAND = 2
	UP      = 1
	DOWN    = -1
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


func Init(upOrderChan chan int, downOrderChan chan int, commandOrderChan chan int, orderOnSameFloorChan chan int, orderInEmptyQueueChan chan int) {

	for {
		select {
		case floor := <-upOrderChan:
			newOrder = Order{UP, floor}
			addInternalOrder(newOrder)

		case floor := <-downOrderChan:
			newOrder = Order{DOWN, floor}
			addInternalOrder(newOrder)

		case floor := <-commandOrderChan:
			newOrder = Order{COMMAND, floor}
			addInternalOrder(newOrder)

		}
	}
}

func ShouldStop(floor int) bool {
	elevators[myIP].LastPassedFloor = floor
	if elevators[myIP].CommandOrders[floor] {
		return true
	}
	if elevators[myIP].Direction == UP {
		if elevators[myIP].UpOrders[floor] || floor == N_FLOORS-1 {
			return true
		} else if ordersAbove(myIP) {
			return false
		} else {
			return true
		}
	} else if elevators[myIP].Direction == DOWN {
		if elevators[myIP].DownOrders[floor] || floor == 0 {
			return true
		} else if ordersBellow(myIP) {
			return false
		} else {
			return true
		}
	}
	return true
}//Ferdig

func AddElevator(newElevator Elevator, IP string) {
	elevators[IP] = &newElevator
} //Ferdig

func OrderCompleted(floor int) {
	for IP, elevator := range elevators {
		elevator.UpOrders[floor] = false
		elevator.DownOrders[floor] = false
		if myIP == IP {
			elevator.CommandOrders[floor] = false
		}
	}
	io.ClearButtonLed(floor,UP)
	io.ClearButtonLed(floor,DOWN)
	io.ClearButtonLed(floor,COMMAND)
} //TODO: send en beskjed til nettverket om at en ordre har blitt fjernet.

func NextDirection() int {
	if elevators[myIP].Direction == UP {
		if ordersAbove() {
			elevators[myIP].Direction = UP
			return UP
		} else if ordersBellow() {
			elevators[myIP].Direction = DOWN
			return DOWN
		}
	} else if elevators[myIP].Direction == DOWN {
		if ordersBellow() {
			elevators[myIP].Direction = DOWN
			return DOWN
		} else if ordersAbove() {
			elevators[myIP].Direction = UP
			return UP
		}
	}
	elevators[myIP].Direction = STOP
	return STOP
} //Utvid til å sjekke andre sine bestillingskøer


/*****************************************************************************************************************
Private
*/

func ordersAbove(IP string) bool {
	for floor := elevators[IP].LastPassedFloor + 1; floor < N_FLOORS; floor++ {
		if elevators[IP].UpOrders[floor] || elevators[IP].CommandOrders || elevators[IP].DownOrders {
			return true
		}
	}
	return false
} //Ferdig

func ordersBellow(IP string) bool {
	for floor := elevators[IP].LastPassedFloor - 1; floor > -1; floor-- {
		if elevators[IP].UpOrders[floor] || elevators[IP].CommandOrders[floor] || elevators[IP].DownOrders[floor] {
			return true
		}
	}
	return false
} //Ferdig

func isQueueEmpty(IP string) bool {
	if ordersAbove(IP) || ordersBellow(IP) {
		return false
	}
	floor := elevators[IP].LastPassedFloor
	if elevators[IP].UpOrders[floor] || elevators[IP].DownOrders[floor] || elevators[IP].CommandOrders[floor] {
		return false
	}
	return true
} //Ferdig

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
} //Ferdig

func addInternalOrder(newOrder Order) {
	if isIdenticalOrder(newOrder) {
		return
	}
	defer func() {
		fmt.Println("Her sender vi ordre oppdatring til alle")
		io.SetButtonLed(newOrder.Floor,newOrder.Type)
	}()

	if newOrder.Type == COMMAND {
		elevators[myIP].CommandOrders[newOrder.floor]
		if elevators[myIP].LastPassedFloor == newOrder.Floor {
			orderOnSameFloorChan <- newOrder.Floor
		}
		return
	}
	cheapestElevator := findCheapestElevator(newOrder)
	if cheapestElevator == myIP{
		if newOrder.Floor == elevators[myIP].LastPassedFloor {
			orderOnSameFloorChan <- newOrder.Floor
		}else if isQueueEmpty(myIP){
			orderInEmptyQueueChan <- newOrder.Floor
		}
	} 
	return
} //Lag Defer, skal ellers være ferdig. Sett på lys.

func findCheapestElevator(newOrder Order) string {
	cheapestElevator := myIP
	minCost := 9999
	for IP, Elevator := range elevators {
		cost := costFunction(Elevator, newOrder)
		if cost < minCost {
			cheapestElevator = IP
		}
		if cost == 0 {
			break
		}
	}
	return cheapestElevator
}//Ferdig

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
} //Ikke laget ennå.