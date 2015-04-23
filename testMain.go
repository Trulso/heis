package main


import(
	"fmt"
	"time"
	//"./queue"
	"net"


)
const N_FLOORS = 4
type Elevator struct {
	direction int
	lastPassedFloor int
	upOrders []bool
	downOrders []bool
	commandOrders []bool
}

// var myIP = //SKaff lokalIPadresse
var elevators = map[string]Elevator{}

var elev1 = Elevator{1,1,[]bool{true,false,true,false},[]bool{false,false,false,false},[]bool{true,true,true,true}}
var elev2 = Elevator{1,2,[]bool{false,false,false,true},[]bool{true,false,true,false},[]bool{true,false,true,true}}

func main(){
	fmt.Println(net.IPv4bcast)
	Addr, _ := net.InterfaceAddrs()
	fmt.Println(Addr)
	fmt.Print("Hei\n")
 	fmt.Println(elev1.upOrders[1])

 	elevators["IP1"] = elev1
 	elevators["IP2"] = elev2

 	fmt.Println(elevators["IP1"].upOrders)
 	fmt.Println(elevators["IP1"].downOrders)
 	if len(elevators) < 3 {
		for floor :=0; floor < N_FLOORS; floor++ {
 		elevators["IP1"].upOrders[floor] = elevators["IP1"].upOrders[floor] || elevators["IP2"].upOrders[floor]
 		elevators["IP1"].downOrders[floor] = elevators["IP1"].downOrders[floor] || elevators["IP2"].downOrders[floor]
 	}
	}

	fmt.Println(elevators["IP1"].upOrders)
	fmt.Println(elevators["IP1"].downOrders)

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
