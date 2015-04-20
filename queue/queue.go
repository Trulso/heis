package queue

import(
	"fmt"
	"../driver"
)

//struct order
//struct elevator
//orders= []order
//elevators= []elevator
const (	
	N_FLOORS			= 4

	IDLE				= 0
	DOOR_OPEN			= 1
	MOVING				= 2
	
	UP 					= 1
	DOWN 				= -1
	STOP				= 0
)



type Elevator struct {
	IP int
	state int
	direction int
	lastPassedFloor int
	orders = [N_FLOORS]Order
}

type Order struct {
    Type int
}

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
	elevators := []Elevator




	for {
		select {
		case floor := <-upOrderChan:

			cheapestElevator := costFunction()

		case floor := <-downOrderChan:

		case floor := <-commandOrderChan:

		}
	}
}

func costFunction(floor int, direction int) Elevator {
	//TODO: Sette inn regnestykket.
}

func ShouldStop(floor int) bool {
	//TODO: sjekke egen bestillingliste
}

func RemoveElevator(Elevator){
	//TODO: fjerne heisen fra listen med heiser
}

func AddElevator(newElevator Elevator){}

func reDistributeOrdersFrom(removedElevator Elevator){

}

