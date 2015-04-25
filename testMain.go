package main


import(
	"fmt"
	"time"
	//"net"


)
const N_FLOORS = 4
type Elevator struct {
	Direction int
	LastPassedFloor int
	UpOrders []bool
	DownOrders []bool
	CommandOrders []bool
}

// var myIP = //SKaff lokalIPadresse
var elevators = map[string]*Elevator{}
var myIP = "IP1"
var elev1 = Elevator{1,1,[]bool{false,false,false,false},[]bool{false,false,false,false},[]bool{false,false,false,false}}
var elev2 = Elevator{1,2,[]bool{false,false,true,false},[]bool{false,false,false,false},[]bool{false,false,false,false}}

func main(){
	ButtonPressed := make([]int, 10)
	ButtonPressed[2]=10
	fmt.Println(ButtonPressed)

 	elevators["IP1"] = &elev1
 	elevators["IP2"] = &elev2

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
	fmt.Println("preprint", elevators["IP1"].UpOrders)
	turnUpQueueFull("IP1")
	fmt.Println("postprint", elevators["IP1"].UpOrders)

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