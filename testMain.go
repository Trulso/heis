package main


import(
	"fmt"
	"time"
	//"./queue"
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
var elev1 = Elevator{1,1,[]bool{true,false,false,false},[]bool{false,false,false,false},[]bool{true,true,false,false}}
var elev2 = Elevator{1,2,[]bool{false,false,false,true},[]bool{true,false,true,false},[]bool{true,false,true,true}}

func main(){

 	elevators["IP1"] = &elev1
 	elevators["IP2"] = &elev2

 	fmt.Println("Oppordrer: ", elevators["IP1"].UpOrders)
 	fmt.Println("Nedordrer: ", elevators["IP1"].DownOrders)
 	if len(elevators) < 3 {
		for floor :=0; floor < N_FLOORS; floor++ {
 		elevators["IP1"].UpOrders[floor] = elevators["IP1"].UpOrders[floor] || elevators["IP2"].UpOrders[floor]
 		elevators["IP1"].DownOrders[floor] = elevators["IP1"].DownOrders[floor] || elevators["IP2"].DownOrders[floor]
 		}
	}
	fmt.Println("Oppordrer: ", elevators["IP1"].UpOrders)
	fmt.Println("Nedordrer: ", elevators["IP1"].DownOrders)

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


func ordersAbove() bool{
	for floor := elevators[myIP].LastPassedFloor + 1; floor < N_FLOORS; floor++ {
		if elevators[myIP].UpOrders[floor] || elevators[myIP].CommandOrders[floor] || elevators[myIP].DownOrders[floor]{
			return true
		}
	}
	return false
}