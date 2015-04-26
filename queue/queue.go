package queue

import (
	"fmt"
	"../driver"
	."../struct"
	"../network"
	"math"
	"time"
	"os"
	"os/exec"
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
var lightUpdateChan = make(chan int)

func Init() {
	CurrentFloor:= driver.GetFloorSensorSignal()
	elev := Elevator{true,true, 1,CurrentFloor,[]bool{false,false,false,false},[]bool{false,false,false,false},[]bool{false,false,false,false}}

	elevators[myIP]=&elev

	go lightUpdater()
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
	elevators[myIP].InFloor = true
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
	//driver.ClearButtonLed(floor,UP)
	//driver.ClearButtonLed(floor,DOWN)
	if byElevator == "self" {
		//driver.ClearButtonLed(floor,COMMAND)
		messageTransmitter("completedOrder",  myIP , Order{-1,floor})
		//fmt.Println("Sender fullført ordre nå")
	}
	lightUpdateChan <-1
}//LIGHTUPDATE

func NextDirection() int {
	//fmt.Println("Vi skal finne neste retning")
	defer func() {
		//fmt.Println("Sender ny retning nå")
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
			//fmt.Println("Her får vi newOrder")
			i := addExternalOrder(message.TargetIP, message.Order)		
			switch i {
			case "empty":
				orderInEmptyQueueChan <- message.Order.Floor
			case "sameFloor":
				orderOnSameFloorChan <- message.Order.Floor
			}
		case "newDirection":
			//fmt.Println("Her får vi newDirection")
			elevators[message.TargetIP].Direction = message.Order.Type
		case "newFloor":
			//fmt.Println("Her får vi newFloor")
			elevators[message.TargetIP].LastPassedFloor = message.Order.Floor
			elevators[message.TargetIP].InFloor = true
		case "completedOrder":
			//fmt.Println("Her får vi completedOrder")
			OrderCompleted(message.Order.Floor, message.TargetIP)
		case "statusUpdate":
			fmt.Println("Her får vi statusUpdate")
			if message.SenderIP != myIP {
				_, exist := elevators[message.TargetIP]
				if !exist{
					newElev := Elevator{true,true,1,0,[]bool{false,false,false,false},[]bool{false,false,false,false},[]bool{false,false,false,false}}
					elevators[message.TargetIP]=&newElev
				}
				elevators[message.TargetIP].InFloor = message.Elevator.InFloor
				elevators[message.TargetIP].LastPassedFloor = message.Elevator.LastPassedFloor
				elevators[message.TargetIP].Direction = message.Elevator.Direction
			
				for floor:=0; floor <N_FLOORS;floor++ {
					elevators[message.TargetIP].UpOrders[floor] = elevators[message.TargetIP].UpOrders[floor] || message.Elevator.UpOrders[floor]
	 				elevators[message.TargetIP].DownOrders[floor] = elevators[message.TargetIP].DownOrders[floor] || message.Elevator.DownOrders[floor]
	 				elevators[message.TargetIP].CommandOrders[floor] = elevators[message.TargetIP].CommandOrders[floor] || message.Elevator.CommandOrders[floor]
				}
				orderInEmptyQueueChan<-1
				lightUpdateChan <-1
			}

			fmt.Println("HER ER STATUSEN VI FÅR TILSENDT")
			fmt.Println(message.TargetIP)
			fmt.Println(message.Elevator)


		case "leftFloor":
			fmt.Printf("Heis %s har forlatt etasjen:\n", message.TargetIP)
			LeftFloor(message.TargetIP)
		}
	}
}//LIGHTCHAN

func HeartbeatReceiver(newElevatorChan chan string, deadElevatorChan chan string){
	for{
		select{
		case IP := <-newElevatorChan:
			if IP != myIP {
				fmt.Printf("Det er dukket opp en ny heis me IP: %s\n", IP)
				_, exist := elevators[IP]
				if exist{
					elevators[IP].Active = true

					fmt.Println("\nSENDER DENNE INFO OM MEG SELV")
					printElevator(myIP)
					messageTransmitter("statusUpdate", myIP,Order{-1,-1})
					time.Sleep(1*time.Millisecond)

					fmt.Println("\nSENDER DENNE INFO OM DEN NYE HEISEN")
					printElevator(IP)
					messageTransmitter("statusUpdate", IP,Order{-1,-1})
				}else {
					newElev := Elevator{true,true,1,0,[]bool{false,false,false,false},[]bool{false,false,false,false},[]bool{false,false,false,false}}
					elevators[IP]=&newElev
					fmt.Println("\nSENDER DENNE INFO OM MEG SELV")
					printElevator(myIP)
					messageTransmitter("statusUpdate", myIP,Order{-1,-1})
				}
			}
		case IP := <-deadElevatorChan:
			elevators[IP].Active = false
			fmt.Printf("Det er fjernet en heis med IP: %s\n", IP)
		}
	}
}

func LeftFloor(IP string){
	if IP != ""{
		elevators[IP].InFloor = false	
	}else{
		messageTransmitter("leftFloor", myIP, Order{-1,-1})
		elevators[myIP].InFloor = false
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
		if elevators[IP].Active == true {
			fmt.Println(IP)
			cost := costFunction(elevator, newOrder,IP)
			fmt.Println ("The cost is ---------------------------------------->",cost)
			if cost < minCost {
				minCost = cost
				cheapestElevator = IP
			}
			if cost == 0 {
				break
			}
		}
	}
	fmt.Println(cheapestElevator)
	return cheapestElevator
}//Ferdig

func costFunction(elevator *Elevator, newOrder Order,ip string) int {
	
	cost := int(math.Abs(float64(newOrder.Floor-elevator.LastPassedFloor)))
	if elevator.Direction == UP && newOrder.Floor<elevator.LastPassedFloor{
		cost += 5
	}else if elevator.Direction == DOWN && newOrder.Floor>elevator.LastPassedFloor{
		cost += 5
	}else if elevator.LastPassedFloor == newOrder.Floor{
		cost +=4
	}
	return cost	
} //Ikke laget ennå.

func addInternalOrder(newOrder Order,b_type int) string{
	var cheapestElevator string

	if isIdenticalOrder(newOrder) {
		return ""
	}
	/*defer func() {
		driver.SetButtonLed(newOrder.Floor,newOrder.Type)
	}()*/

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
	lightUpdateChan <-1
	if cheapestElevator == myIP{
		if newOrder.Floor == elevators[myIP].LastPassedFloor {
			return "sameFloor" 
		}else if firstOrder{
			return "empty"
		}
	}
	return ""
}//LIGHTUPDATE

func addExternalOrder(taskedElevator string, newOrder Order) string {
	firstOrder:= isQueueEmpty(myIP)
	if newOrder.Type == UP {
		elevators[taskedElevator].UpOrders[newOrder.Floor]=true
		//driver.SetButtonLed(newOrder.Floor,newOrder.Type)
	}else if newOrder.Type == DOWN {
		elevators[taskedElevator].DownOrders[newOrder.Floor]=true
		//driver.SetButtonLed(newOrder.Floor,newOrder.Type)
	}else{
		elevators[taskedElevator].CommandOrders[newOrder.Floor]=true
	}
	lightUpdateChan <-1
	if taskedElevator == myIP {
		if newOrder.Floor == elevators[myIP].LastPassedFloor {
			return "sameFloor" 
		}else if firstOrder{
			return "empty"
		}
	}
	return ""
}//LIGHTUPDATE

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

func printElevator(elevatorIP string){
	fmt.Printf("\nHeis IP: %s\n", elevatorIP)
	//fmt.Printf("Active: %t\n", elevators[elevatorIP].Active)
	if elevators[elevatorIP].Direction == 1 {
		fmt.Printf("Direction: UP\n")
	}else if elevators[elevatorIP].Direction == -1 {
		fmt.Printf("Direction: DOWN\n")
	}else {
		fmt.Printf("Direction: STOP\n")
	}
	fmt.Printf("In floor: %t\n", elevators[elevatorIP].InFloor)
	fmt.Printf("Last Passed Floor: %d\n", elevators[elevatorIP].LastPassedFloor)

	fmt.Printf("|  UP\t| DOWN\t|COMMAND|\n")
	for floor := N_FLOORS-1; floor>-1; floor--{
		fmt.Printf("| %t\t| %t\t| %t\t|\n",elevators[elevatorIP].UpOrders[floor],elevators[elevatorIP].DownOrders[floor],elevators[elevatorIP].CommandOrders[floor])
	}
}

func StatusPrint(){
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
	statusTimer:= time.NewTimer(1 * time.Second)
	statusTimer.Stop()
	for{
		statusTimer.Reset(10 * time.Second)
		<-statusTimer.C
		fmt.Println("\n\t\tELEVATOR STATUS")
		for IP, _ := range elevators{
			printElevator(IP)
		}
	}
}

func lightUpdater(){
	commandLights := make([]bool, N_FLOORS)
	upLights := make([]bool,N_FLOORS)
	downLights := make([]bool,N_FLOORS)
	for{
		<-lightUpdateChan
		for floor:=0; floor <N_FLOORS; floor++{
			for IP,_ := range elevators {
				if elevators[myIP].CommandOrders[floor]{
					commandLights[floor] = true
				}
				if elevators[IP].UpOrders[floor]{
					upLights[floor] = true
				}
				if elevators[IP].DownOrders[floor]{
					downLights[floor] = true		
				}
			}
		}
		for floor:=0;floor<N_FLOORS;floor++{
			if floor > 0 && downLights[floor] {
				driver.SetButtonLed(floor,DOWN)
			}else{
				driver.ClearButtonLed(floor,DOWN)
			}
			if floor < N_FLOORS-1 && upLights[floor] {
				driver.SetButtonLed(floor,UP)
			}else{
				driver.ClearButtonLed(floor,UP)
			}
			if commandLights[floor] {
				driver.SetButtonLed(floor,COMMAND)
			}else{
				driver.ClearButtonLed(floor,COMMAND)
			}
			downLights[floor]=false
			commandLights[floor]=false
			upLights[floor]=false
		}

	}
}