package queue

import (
	"fmt"
	"../driver"
	."../struct"
	"../network"
	//"math/rand"
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
var myIP string =network.GetIP() 

func Init() {
	CurrentFloor:= driver.GetFloorSensorSignal()
	elev := Elevator{true, 1,CurrentFloor,[]bool{false,false,false,false},[]bool{false,false,false,false},[]bool{false,false,false,false}}

	elevators[myIP]=&elev
}

func OrderButtonHandler(upOrderChan chan int, downOrderChan chan int, commandOrderChan chan int, orderOnSameFloorChan chan int, orderInEmptyQueueChan chan int){
	for {
		select {
		case floor := <-upOrderChan:
			newOrder := Order{UP, floor}
			i := addInternalOrder(newOrder,UP)
			fmt.Println("Opp ordre")
			switch i {
			case "empty":
				orderInEmptyQueueChan <- floor
			case "sameFloor":
				orderOnSameFloorChan <- floor
			}

		case floor := <-downOrderChan:
			newOrder := Order{DOWN, floor}
			i:=addInternalOrder(newOrder,DOWN)
			fmt.Println("Ned ordre")
			switch i {
			case "empty":
				orderInEmptyQueueChan <- floor
			case "sameFloor":
				orderOnSameFloorChan <- floor
			}

		case floor := <-commandOrderChan:	
			newOrder := Order{COMMAND, floor}
			i:=addInternalOrder(newOrder,COMMAND)
			fmt.Printf("Command ordre: %s\n", i)
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
	defer messageTransmitter("newFloor", myIP, Order{-1,floor}) 
	//defer fmt.Println("Sender newFloor nå")
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
}

func AddElevator(newElevator *Elevator, IP string) {
	elevators[IP] = newElevator
} //Ferdig kanskje...

func OrderCompleted(floor int, byElevator string) {
	for IP, elevator := range elevators {
		elevator.UpOrders[floor] = false
		elevator.DownOrders[floor] = false
		if byElevator == IP  || (byElevator == "self" && myIP == IP){
			elevator.CommandOrders[floor] = false
		}
	}
	//fmt.Println("KOmmer vi hit?")
	driver.ClearButtonLed(floor,UP)
	driver.ClearButtonLed(floor,DOWN)
	if byElevator == "self" {
		driver.ClearButtonLed(floor,COMMAND)
		messageTransmitter("completedOrder",  myIP , Order{-1,floor})
		fmt.Println("Sender fullført ordre nå")
	}
}

func NextDirection() int {
	fmt.Println("Vi skal finne neste retning")
	fmt.Printf("COM:%v\n", elevators[myIP].CommandOrders)
	fmt.Printf("DWN:%v\n", elevators[myIP].DownOrders)
	fmt.Printf("UP: %v\n", elevators[myIP].UpOrders)
	defer func() {
		fmt.Println("Sender ny retning nå")
		messageTransmitter("newDirection", myIP, Order{elevators[myIP].Direction, -1})
	}()
	lastDir := elevators[myIP].Direction
	if lastDir == UP || lastDir == STOP {
		if ordersAbove(myIP) {
			elevators[myIP].Direction = UP
			return UP
		} else if ordersBellow(myIP) {
			elevators[myIP].Direction = DOWN
			return DOWN
		}
	} else if lastDir == DOWN {
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
} //Utvid til å sjekke andre sine bestillingskøer, sende direction etterpå

func MessageReceiver(incommingMsgChan chan Message, orderOnSameFloorChan chan int, orderInEmptyQueueChan chan int){
	for{
		message := <-incommingMsgChan
		switch message.MessageType{
		case "newOrder":
			fmt.Println("Her får vi newOrder")
			i := addExternalOrder(message.TargetIP, message.Order)		
			switch i {
			case "empty":
				orderInEmptyQueueChan <- message.Order.Floor
			case "sameFloor":
				orderOnSameFloorChan <- message.Order.Floor
			}
		case "newDirection":
			fmt.Println("Her får vi newDirection")
			elevators[message.TargetIP].Direction = message.Order.Type
		case "newFloor":
			fmt.Println("Her får vi newFloor")
			elevators[message.TargetIP].LastPassedFloor = message.Order.Floor
		case "completedOrder":
			fmt.Println("Her får vi completedOrder")
			OrderCompleted(message.Order.Floor, message.TargetIP)
		case "statusUpdate":
			fmt.Println("Her får vi statusUpdate")
			if message.TargetIP == myIP {
				for floor:= 0; floor<N_FLOORS;floor++{
					elevators[myIP].UpOrders[floor]      = elevators[myIP].UpOrders[floor] || message.Elevator.UpOrders[floor]
					elevators[myIP].DownOrders[floor]    = elevators[myIP].DownOrders[floor] || message.Elevator.DownOrders[floor]
					elevators[myIP].CommandOrders[floor] = elevators[myIP].CommandOrders[floor] || message.Elevator.CommandOrders[floor]
				}	
			}else {
				elevators[message.TargetIP] = &message.Elevator
			}
			fmt.Println("Her har vi oppdatert status")
		}
	}
}

func HeartbeatReceiver(newElevatorChan chan string, deadElevatorChan chan string){
	for{
		select{
		case IP := <-newElevatorChan:
			fmt.Printf("Det er dukket opp en ny heis me IP: %s\n", IP)
			_, exist := elevators[IP]
			if exist{
				elevators[IP].Active = true
			}else{
				newElev := Elevator{true,1,0,[]bool{false,false,false,false},[]bool{false,false,false,false},[]bool{false,false,false,false}}
				elevators[IP]=&newElev
			}
			for _,elev := range elevators{
				fmt.Println(elev)
			}
			messageTransmitter("statusUpdate", myIP,Order{-1,-1})
			messageTransmitter("statusUpdate", IP,Order{-1,-1})
		case IP := <-deadElevatorChan:
			elevators[IP].Active = false
			fmt.Printf("Det er fjernet en heis me IP: %s\n", IP)
		}
	}
}

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
	for IP, elevator := range elevators {
		fmt.Println(IP)
		cost := costFunction(elevator, newOrder,IP)
		fmt.Println ("READ!!!!! :::::::::::::::",cost)
		if cost < minCost {
			minCost = cost
			cheapestElevator = IP
		}
		if cost == 0 {
			break
		}
	}
	fmt.Println(cheapestElevator)
	return cheapestElevator
}//Ferdig

func costFunction(elevator *Elevator, newOrder Order,ip string) int {
	
	cost := 0

	var topOrder int

	var lowestFloor int 

	if isQueueEmpty(ip){
		topOrder = newOrder.Floor
	}
	if isQueueEmpty(ip){
		lowestFloor = newOrder.Floor
	}

	//for i := 0; i < 5; i++
	fmt.Println("Elevator Direction: ",elevator.Direction)
	fmt.Println("Order Type",newOrder.Type)
	if elevator.LastPassedFloor < newOrder.Floor{


		if elevator.Direction == UP || elevator.Direction == 0 {
			fmt.Println("GÅR INN I ELEVATOR UP")
			if newOrder.Type == UP {
				fmt.Println("Elevator: UP .. Order: UP")
				cost += newOrder.Floor - elevator.LastPassedFloor
				for i := elevator.LastPassedFloor; i < newOrder.Floor; i++ {
					if elevator.UpOrders[i] == true || elevator.CommandOrders[i] == true{
						cost += 1
					}
				return cost	
				}
			}
			if newOrder.Type == DOWN {
				fmt.Println("Elevator: UP .. Order: Down")
				if elevator.CommandOrders[newOrder.Floor] == true || elevator.UpOrders[newOrder.Floor] == true{
					cost += newOrder.Floor - elevator.LastPassedFloor
					for i := elevator.LastPassedFloor; i < newOrder.Floor; i++ {
						if elevator.UpOrders[i] == true || elevator.CommandOrders[i] == true{
							cost += 1
						}
					}	
					return cost
				}else{
					
					for i := elevator.LastPassedFloor; i < N_FLOORS; i++{
						if elevator.UpOrders[i] == true || elevator.CommandOrders[i] == true  {
							cost += 1
							topOrder=i
						}else if elevator.DownOrders[i] == true{
							topOrder=i

						}				
					}

					for i := topOrder; i >= newOrder.Floor; i-- {
						if elevator.DownOrders[i] == true {
							cost += 1
						}

					}
					cost += topOrder - elevator.LastPassedFloor
					cost += topOrder - newOrder.Floor
					return cost
				} 
			}
		}
		if elevator.Direction == DOWN  || elevator.Direction == 0 { {
	
			fmt.Println("GÅR INN I ELEVATOR DOWN")
			for i := elevator.LastPassedFloor; i >= 0; i-- {
				if elevator.DownOrders[i] == true || elevator.CommandOrders[i] == true{
					cost += 1
					lowestFloor = i
				}				
			}
			cost += elevator.LastPassedFloor - lowestFloor
			}

			if newOrder.Type == UP {
				fmt.Println("Elevator: DOWN .. Order: UP")
				for i := lowestFloor; i < newOrder.Floor; i++{
					if elevator.UpOrders[i] == true {
						cost += 1
					}
				}
				cost += newOrder.Floor - lowestFloor
				
				return cost
			}
		
			if newOrder.Type == DOWN {
				fmt.Println("Elevator: DOWN .. Order: DOWN")
				if elevator.CommandOrders[newOrder.Floor] == true{

					for i := lowestFloor; i < newOrder.Floor; i++{
						if elevator.UpOrders[i] == true {
							cost += 1
						}
					cost += newOrder.Floor - lowestFloor
					}
					return cost

				}
				for i := 0; i < N_FLOORS; i++ {
					if elevator.UpOrders[i] == true {
						cost += 1
						topOrder = i
					}else if elevator.DownOrders[i] == true {
						topOrder = i
					}else{
						topOrder = newOrder.Floor
					}
				}
				fmt.Println("TOP ORDER :>-----",topOrder)
				fmt.Println("LOWESST FLOOR :>----",lowestFloor)
				cost += topOrder - lowestFloor
				cost += topOrder - newOrder.Floor
				return cost
				// burda EGETNLIG sjekka down orders på veien ned igjen meeeeeeeeeeeeeeeeen.
						
			}
			
		}
	}

	/*if elevator.LastPassedFloor > newOrder.Floor{  // ONE COST TO BIND THEM ALL

		fmt.Println("THE FLOOR IS UNDER")

		if elevator.Direction == UP || elevator.Direction == 0 {
			fmt.Println("GÅR INN I ELEVATOR UP")
			if newOrder.Type == UP {
				fmt.Println("hallo")		
			}
			if newOrder.Type == DOWN {
				fmt.Println("Elevator: UP .. Order: Down")
			}
		}
		if elevator.Direction == DOWN  || elevator.Direction == 0 { 
	
			fmt.Println("GÅR INN I ELEVATOR DOWN")


			if newOrder.Type == UP {
				fmt.Println("Elevator: DOWN .. Order: UP")
		
						
			}
		
			if newOrder.Type == DOWN {
				fmt.Println("Elevator: DOWN .. Order: DOWN")
						
			}
			
		}	
	}
fmt.Println("finner ikke kost!!!!!!!!!!!!!!!!!!!!!!!!!!")*/
return 1
} //Ikke laget ennå.


func addInternalOrder(newOrder Order,b_type int) string{
	var cheapestElevator string

	if isIdenticalOrder(newOrder) {
		return ""
	}
	defer func() {
		driver.SetButtonLed(newOrder.Floor,newOrder.Type)
	}()

	if b_type != COMMAND{
		cheapestElevator = findCheapestElevator(newOrder)
	}else{
		cheapestElevator = myIP
	}
	
	firstOrder:= isQueueEmpty(myIP)
	for IP, Ele := range elevators {
		if IP == cheapestElevator {
			if newOrder.Type == UP {
				Ele.UpOrders[newOrder.Floor]=true
				fmt.Println("Legger til en oppoverordre")
			}else if newOrder.Type == DOWN {
				Ele.DownOrders[newOrder.Floor]=true
				fmt.Println("Legger til en nedoverordre")
			}else{
				Ele.CommandOrders[newOrder.Floor]=true
				fmt.Println("Legger til en commandordre")
			}
		}
	}
	messageTransmitter("newOrder", cheapestElevator,newOrder)
	//fmt.Println("Sender newOrder nå")
	if cheapestElevator == myIP{
		if newOrder.Floor == elevators[myIP].LastPassedFloor {
			return "sameFloor" 
		}else if firstOrder{
			return "empty"
		}
	}
	return ""
}

func addExternalOrder(taskedElevator string, newOrder Order) string {
	firstOrder:= isQueueEmpty(myIP)
	if newOrder.Type == UP {
		elevators[taskedElevator].UpOrders[newOrder.Floor]=true
		driver.SetButtonLed(newOrder.Floor,newOrder.Type)
	}else if newOrder.Type == DOWN {
		elevators[taskedElevator].DownOrders[newOrder.Floor]=true
		driver.SetButtonLed(newOrder.Floor,newOrder.Type)
	}else{
		elevators[taskedElevator].CommandOrders[newOrder.Floor]=true
	}
	if taskedElevator == myIP {
		if newOrder.Floor == elevators[myIP].LastPassedFloor {
			return "sameFloor" 
		}else if firstOrder{
			return "empty"
		}
	}
	return ""
}

func messageTransmitter(msgType string, targetIP string, order Order){ //newOrder, floorUpdate, completedOrder, directionUpdate, 
	//fmt.Printf("Nå lager vi en %s type message\n", msgType)
	newMessage := Message{
		msgType,
		myIP,
		targetIP,
		*(elevators[targetIP]),
		order,
	}
	network.BroadcastMessage(newMessage)
}

