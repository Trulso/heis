package queue

import(
	//"fmt"
	//io "../driver"
	//io "../fakeDriver"
)

//struct order
//struct elevator
//orders= []order
//elevators= []elevator
const (	
	N_FLOORS			= 4
	
	UP 					= 1
	STOP				= 0
	DOWN 				= -1
	
)


type Message struct {
	MessageType string //neworder,just arrived, status update, completed order,
	SenderIP string
	elevators map[string]Elevator

}

type Order struct{
	Direction int
	Floor int
}


type Elevator struct {
	Direction int
	LastPassedFloor int
	UpOrders []bool
	DownOrders []bool
	CommandOrders []bool
}

var myIP = string //SKaff lokalIPadresse
var elevators = map[string]Elevator

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


	){

	for {
		select {
		case floor := <-upOrderChan:
			newOrder=Order{UP,floor}

		case floor := <-downOrderChan:

		case floor := <-commandOrderChan:

		}
	}
}
func isIdenticalOrder(newOrder Order){

}

func cheapestElevator() Elevator {
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
			if
 
		} 

	}
	return cost
}

func ShouldStop(floor int) bool {

	//TODO: sjekke egen bestillingliste
}

func RemoveElevator(string IP){
	reDistributeOrdersFrom(IP)
	delete(elevators, IP)
}

func AddElevator(newElevator Elevator, IP string){
	elevators[IP] = newElevator
}

func OrderCompleted(floor int){
	for IP, elevator := range elevators {
		elevator.UpOrders[floor] = false
		elevator.DownOrders[floor] = false
		if myIP == IP{
			elevator.CommandOrders[floor] = false
		}	
	}
	//TODO: send en beskjed til netverket om at en ordre har blitt fjernet.
}

func reDistributeOrdersFrom(string IP){
	if len(elevators) < 3 { //Tar alle bestillingene selv om man er eneste heis igjen.
		for floor :=0; floor < N_FLOORS; floor++ {
 		elevators[myIP].UpOrders[floor] = elevators[myIP].UpOrders[floor] || elevators[IP].UpOrders[floor]
 		elevators[myIP].DownOrders[floor] = elevators[myIP].DownOrders[floor] || elevators[IP].DownOrders[floor]
 		}
	} else {//Regner vekt til alle ordre som removedElevator hadde, og tilegner dem en ny heis.
		for floor := 0; floor < N_FLOORS; floor++ {
			if 
		} 
	

	}

	
	//Sier fra til netverket
}

func NextOrder() int {
	if elevators[myIP].Direction == UP {
		for floor := elevators[myIP].LastPassedFloor + 1; floor < N_FLOORS; floor++ {
			if elevators[myIP].UpOrders[floor] || elevators[myIP].CommandOrders {
				return UP
			}
		}
		for floor := elevators[myIP].LastPassedFloor -1 ; floor > -1; floor-- {
			if elevators[myIP].UpOrders[floor] || elevators[myIP].CommandOrders {
				return DOWN
			}
		}
	}


}