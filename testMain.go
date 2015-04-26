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
var elev1 = Elevator{1,1,[]bool{false,false,false,false},[]bool{false,true,false,false},[]bool{false,false,false,false}}
var elev2 = Elevator{1,2,[]bool{false,false,false,false},[]bool{false,false,false,false},[]bool{false,false,false,false}}
var elev3 = Elevator{1,2,[]bool{false,false,false,false},[]bool{false,false,false,false},[]bool{false,false,false,false}}
var elev4 = Elevator{1,2,[]bool{false,false,false,false},[]bool{false,false,false,false},[]bool{false,false,false,false}}


func main(){

 	elevators["IP1"] = &elev1
 	elevators["IP2"] = &elev2
 	AddElevator(&elev3, "IP3")
 	AddElevator(&elev4, "IP4")

 	for IP, elev := range elevators {
 		fmt.Println(IP)
 		fmt.Println(elev)
 	}



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
		doorTimer.Reset(5 * time.Second)
		<-doorTimer.C
		fmt.Print("Det har gaatt fem sekund\n")
		
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
