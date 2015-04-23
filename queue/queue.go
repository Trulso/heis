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
	DOWN 				= -1
	STOP				= 0
)



type Elevator struct {
	IP int
	Direction int
	LastPassedFloor int
	UpOrders [N_FLOORS]bool
	DownOrders [N_FLOORS]bool
	CommandOrders [N_FLOORS]bool
}

var myIP = //SKaff lokalIPadresse
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

	){

	for {
		select {
		case floor := <-upOrderChan:

		case floor := <-downOrderChan:

		case floor := <-commandOrderChan:

		}
	}
}
func isIdenticalOrder(floor int, direction int){

}

func cheapestElevator() Elevator {
	//TODO: Bruker costFunction til å finne den billigste heisen
}

//b
func costFunction(elevator Elevator, floor int, direction int) int {
	cost := 0
	difference := elevator.LastPassedFloor - floor
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


}