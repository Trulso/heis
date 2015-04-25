package queue

import (
	"fmt"
	io "../driver"
	."../struct"
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

var elevators = map[string]*Elevator{}
var myIP string ="myIP" 
var elev = Elevator{1,0,[]bool{false,false,false,false},[]bool{false,false,false,false},[]bool{false,false,false,false}}

func Init(upOrderChan chan int, downOrderChan chan int, commandOrderChan chan int, orderOnSameFloorChan chan int, orderInEmptyQueueChan chan int) {
	elevators[myIP]=&elev
	for {
		select {
		case floor := <-upOrderChan:
			newOrder := Order{UP, floor}
			i := addInternalOrder(newOrder)
			fmt.Println("Opp ordre")
			switch i {
			case "empty":
				orderInEmptyQueueChan <- floor
			case "sameFloor":
				orderOnSameFloorChan <- floor
			}

		case floor := <-downOrderChan:
			newOrder := Order{DOWN, floor}
			i:=addInternalOrder(newOrder)
			fmt.Println("Ned ordre")
			switch i {
			case "empty":
				orderInEmptyQueueChan <- floor
			case "sameFloor":
				orderOnSameFloorChan <- floor
			}

		case floor := <-commandOrderChan:	
			newOrder := Order{COMMAND, floor}
			i:=addInternalOrder(newOrder)
			fmt.Println("Command ordre")
			switch i {
			case "empty":
				orderInEmptyQueueChan <- floor
			case "sameFloor":
				orderOnSameFloorChan <- floor
			}
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
	fmt.Println("Vi skal finne neste retning")
	fmt.Printf("COM:%v\n", elevators[myIP].CommandOrders)
	fmt.Printf("DWN:%v\n", elevators[myIP].DownOrders)
	fmt.Printf("UP: %v\n", elevators[myIP].UpOrders)
	if elevators[myIP].Direction == UP || elevators[myIP].Direction == STOP {
		if ordersAbove(myIP) {
			elevators[myIP].Direction = UP
			return UP
		} else if ordersBellow(myIP) {
			elevators[myIP].Direction = DOWN
			return DOWN
		}
	} else if elevators[myIP].Direction == DOWN {
		if ordersBellow(myIP) {
			elevators[myIP].Direction = DOWN
			return DOWN
		} else if ordersAbove(myIP) {
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
		if elevators[IP].UpOrders[floor] || elevators[IP].CommandOrders[floor] || elevators[IP].DownOrders[floor] {
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
	for _, elevator := range elevators {
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


func findCheapestElevator(newOrder Order) string {
	cheapestElevator := myIP
	minCost := 9999
	for IP, Ele := range elevators {
		cost := costFunction(Ele, newOrder)
		if cost < minCost {
			cheapestElevator = IP
		}
		if cost == 0 {
			break
		}
	}
	return cheapestElevator
}//Ferdig

func costFunction(elevator *Elevator, newOrder Order) int {
	cost := 0
	difference := elevator.LastPassedFloor - newOrder.Floor
	if elevator.Direction == STOP {
		cost = cost + 1*difference
		return cost
	} else if elevator.Direction == UP {
		if difference > 0 {
		}
	}
	return 1
} //Ikke laget ennå.


func addInternalOrder(newOrder Order) string{
	if isIdenticalOrder(newOrder) {
		return ""
	}

	defer func() {
		fmt.Println("Her sender vi ordre oppdatring til alle")
		io.SetButtonLed(newOrder.Floor,newOrder.Type)
	}()

	cheapestElevator := findCheapestElevator(newOrder)
	FirstOrder:= isQueueEmpty(myIP)
	for IP, Ele := range elevators {
		if IP == cheapestElevator{
			if newOrder.Type == UP {
				Ele.UpOrders[newOrder.Floor]=true
			}else if newOrder.Type == DOWN {
				Ele.DownOrders[newOrder.Floor]=true
			}else{
				Ele.CommandOrders[newOrder.Floor]=true
			}
		}
	}
	if cheapestElevator == myIP{
		if newOrder.Floor == elevators[myIP].LastPassedFloor {
			return "sameFloor" 
		}else if FirstOrder{
			return "empty"
		}
	}
	fmt.Println("Kom helt til enden")
	return ""
}

func addExternalOrder(newOrder Order){

	if isIdenticalOrder(newOrder) == false {

		if newOrder.Type == UP {
			io.SetButtonLed(newOrder.Floor,newOrder.Type)
			elevators[myIP].UpOrders[newOrder.Floor]=true
		}else if newOrder.Type == DOWN {
			io.SetButtonLed(newOrder.Floor,newOrder.Type)
			elevators[myIP].DownOrders[newOrder.Floor]=true
		}	
	}
}

 //Bør skrives om for å kunne vite om det var en tom kø, eller om det var en bestilling i samme etg.
//Må også sende alt till nettet.


//Do this order
//Elevater X update
//I have done this Floor (remove exterior)
//Acknowlage
//I am a new elevator!

// queue funksjoner som er nyttig :
// func addInternalOrder(newOrder Order) string{
// func AddElevator(newElevator Elevator, IP string) {
//func OrderCompleted(floor int) {

/*

type Message struct {
	MessageType string //neworder,just arrived, status update, completed order,
	SenderIP    string
	Elevators   map[string]Elevator
	ThisFloor   Order  //Bruker denne til å dele kort informasjon. Gå dit, jeg har kommet hit OSV. 

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


*/


func StatusDecoder(toGet chan Message){

	RxMessage := <-toGet
	fmt.Println("EVER INNNNNNNNNNNNNNNNNNNNNNNNNNNNN")
	if RxMessage.MessageType == "newOrder" {
		addExternalOrder(RxMessage.ThisFloor)

	}

}