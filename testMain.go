package main


import(
	"fmt"
	"time"
	//"net"


)
const (
	N_FLOORS = 4
	UP=1
	DOWN=-1
	)

type Order struct {
	Type  int
	Floor int
}

type Elevator struct {
	Active bool
	Direction int
	LastPassedFloor int
	UpOrders []bool
	DownOrders []bool
	CommandOrders []bool
}
type Message struct {
	MessageType string //newOrder,just arrived, status update, completed order,
	SenderIP    string
	TargetIP	string //Which elevator that changes
	Elevator Elevator
	Order   Order
}

// var myIP = //SKaff lokalIPadresse
var elevators = map[string]*Elevator{}
var myIP = "IP1"
var elev1 = Elevator{true,1,1,[]bool{false,false,false,false},[]bool{false,true,false,false},[]bool{false,false,false,false}}
var elev2 = Elevator{false,1,2,[]bool{false,false,false,false},[]bool{false,false,false,false},[]bool{false,false,false,false}}
var elev3 = Elevator{true,1,2,[]bool{false,false,false,false},[]bool{false,false,false,false},[]bool{false,false,false,false}}
var elev4 = Elevator{true,1,2,[]bool{false,false,false,false},[]bool{false,false,false,false},[]bool{false,false,false,false}}


func main(){

 	elevators["129.241.187.141"] = &elev1
 	elevators["129.241.187.143"] = &elev2
 	AddElevator(&elev3, "IP3")
 	AddElevator(&elev4, "IP4")

 	printElevator("129.241.187.141")
 
 

 // 	fmt.Println("Oppordrer: ", elevators["IP1"].UpOrders)
 // 	fmt.Println("Nedordrer: ", elevators["IP1"].DownOrders)
 // 	if len(elevators) < 3 {
	// 	for floor :=0; floor < N_FLOORS; floor++ {
 // 		elevators["IP1"].UpOrders[floor] = elevators["IP1"].UpOrders[floor] || elevators["IP2"].UpOrders[floor]
 // 		elevators["IP1"].DownOrders[floor] = elevators["IP1"].DownOrders[floor] || elevators["IP2"].DownOrders[floor]
 // 		}
	// }
	// fmt.Println("Oppordrer: ", elevators["IP1"].UpOrders)
	// fmt.Println("Nedordrer: ", elevators["IP1"].DownOrders)

	/*fmt.Println(ordersAbove())
	elevators["IP1"].LastPassedFloor = 0
	fmt.Println(ordersAbove())
*/

	doorTimer := time.NewTimer(1000*time.Millisecond)
	<-doorTimer.C
	fmt.Print("Why deadlock\n")
	doorTimer.Stop()
	for{
		doorTimer.Reset(3 * time.Second)
		<-doorTimer.C
		printElevator("129.241.187.141")
		
	}
}


func ordersAbove(IP string) bool{
	for floor := elevators[IP].LastPassedFloor + 1; floor < N_FLOORS; floor++ {
		if elevators[IP].UpOrders[floor] || elevators[IP].CommandOrders[floor] || elevators[IP].DownOrders[floor]{
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

func turnUpQueueFull(IP string){
	
	defer func() {
		fmt.Println("deferprint", elevators[IP].UpOrders)
	}()
	return
	for i:=0;i<4;i++{
		elevators[IP].UpOrders[i]=true
	}

}

func ShouldStop(floor int) bool {
	elevators[myIP].LastPassedFloor=floor
	if elevators[myIP].CommandOrders[floor]{
		return true
	}
	if elevators[myIP].Direction == UP {
		if elevators[myIP].UpOrders[floor] || floor == N_FLOORS-1 {
			return true
		}else if ordersAbove(myIP){
			return false
		}else{
			return true
		}
	}else if elevators[myIP].Direction == DOWN {
		if elevators[myIP].DownOrders[floor] || floor == 0 {
			return true
		}else if ordersBellow(myIP){
			return false
		}else{
			return true
		}
	}
	return true
}

func AddElevator(newElevator *Elevator, IP string) {
	elevators[IP] = newElevator
} //Ferdig

func printElevator(elevatorIP string){
	fmt.Printf("\nHeis IP: %s\n", elevatorIP)
	fmt.Printf("Active: %t\n", elevators[elevatorIP].Active)
	if elevators[elevatorIP].Direction == 1 {
		fmt.Printf("Direction: UP\n")
	}else if elevators[elevatorIP].Direction == -1 {
		fmt.Printf("Direction: DOWN\n")
	}else {
		fmt.Printf("Direction: STOP\n")
	}
	fmt.Printf("Last Passed Floor: %d\n", elevators[elevatorIP].LastPassedFloor)

	fmt.Printf("|  UP\t| DOWN\t|COMMAND|\n")
	for floor := N_FLOORS-1; floor>-1; floor--{
		fmt.Printf("| %t\t| %t\t| %t\t|\n",elevators[elevatorIP].UpOrders[floor],elevators[elevatorIP].DownOrders[floor],elevators[elevatorIP].CommandOrders[floor])
	}
	fmt.Printf("\n")

}