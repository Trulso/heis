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
	iP int
	direction int
	lastPassedFloor int
	upOrders [N_FLOORS]bool
	downOrders [N_FLOORS]bool
	commandOrders [N_FLOORS]bool
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

			cheapestElevator := cheapestElevator()

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
	difference := elevator.lastPassedFloor - floor
	if elevator.direction == STOP {
		cost = cost + 1*difference
		return cost
	} else if elevator.direction == UP {
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
		elevator.upOrders[floor] = false
		elevator.downOrders[floor] = false
		if myIP == IP{
			elevator.commandOrders[floor] = false
		}
	
	}
	//TODO: send en beskjed til netverket om at en ordre har blitt fjernet.
}

func reDistributeOrdersFrom(string IP){
	//Regner vekt til alle ordre som removedElevator hadde, og tilegner dem en ny heis.
}

func NextOrder() int {


}