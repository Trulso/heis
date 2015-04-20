package fakeDriver

import (
	"fmt"
	"time"
)

const (
	N_FLOORS=4
	Up   = 1
	Down = -1
	Stop = 0
	On   = 1
	Off  = 0
	Command = 2
)



func HwInit() int {

	fmt.Printf("initalizing hw\n")
	return 1
}


func GetFloorSensorSignal() int {
	return -1
}

func GetObstructionSignal() int {
	return 0
}

func GetStopSignal() int {
	return 0
}

func SetDoorLamp(value int) {
	if (value) > 0 {
	   fmt.Printf("Doeren er aapen.\n")
	}else{
        fmt.Printf("Doeren er lukket.\n")
    }	
}

func SetStopLamp(value int) {
	if (value) > 0{
        fmt.Printf("Lyset i stoppknappen er paa.\n")
	}else{
	    fmt.Printf("Lyset i stoppknappen er av.\n")
	}
}

func SetMotorDir(dir int) {
	if (dir == 0){
	    fmt.Printf("Motor staar stille.\n")
	}
	if (dir > 0) {
       	fmt.Printf("Motor gaar oppover\n")
	}
    	if (dir < 0) {
    	fmt.Printf("Motor gaar nedover\n")
    }
}

func getUpOrdersSignal(floor int) int {
	return 0
}

func UpOrdersPolling(upOrdersChannel chan int) {
    var ButtonPressed [N_FLOORS-1]int
	for {
		for floor := 0; floor<N_FLOORS-1; floor++ {
			if getUpOrdersSignal(floor) == 1 {
				if ButtonPressed[floor] == 0 {
					upOrdersChannel<-floor
					ButtonPressed[floor] = 1
				}
			}else {
				if ButtonPressed[floor] == 1 {
					ButtonPressed[floor] = 0
				}
			}
		}
		time.Sleep(1*time.Millisecond)
	}
}

func getDownOrdersSignal(floor int) int {
	return 0
}

func DownOrdersPolling(downOrdersChannel chan int) {
    var ButtonPressed [N_FLOORS-1]int
	for {
		for floor := 1; floor<N_FLOORS; floor++ {
			if getDownOrdersSignal(floor) == 1 {
				if ButtonPressed[floor-1] == 0 {
					downOrdersChannel<-floor
					ButtonPressed[floor-1] = 1
				}
			}else {
				if ButtonPressed[floor-1] == 1 {
					ButtonPressed[floor-1] = 0
				}
			}
		}
		time.Sleep(1*time.Millisecond)
	}
}

func getCommandOrdersSignal(floor int) int {
	return 0
}

func CommandOrdersPolling(commandOrdersChannel chan int) {
    var ButtonPressed [N_FLOORS]int
	for {
		for floor := 0; floor<N_FLOORS; floor++ {
			if getCommandOrdersSignal(floor) == 1 {
				if ButtonPressed[floor] == 0 {
					commandOrdersChannel<-floor
					ButtonPressed[floor] = 1
				}
			}else {
				if ButtonPressed[floor] == 1 {
					ButtonPressed[floor] = 0
				}
			}
		}
		time.Sleep(1*time.Millisecond)
	}
}

func SetFloorIndicator(floor int) {
    fmt.Printf("Heisen er i etasje %d\n", floor)
}

func SetButtonLed(floor int,button int){
	// if(floor<=N_FLOORS){
	// 	if button == Command {
	// 		Io_set_bit(LIGHT_COMMAND1-floor)
	// 	}
	// 	if button == Up {
	// 		if floor == 0 {
	// 			Io_set_bit(LIGHT_UP1)
	// 		}
	// 		if floor == 1 {
	// 			Io_set_bit(LIGHT_UP2)
	// 		}
	// 		if floor == 2 {
	// 			Io_set_bit(LIGHT_UP3)
	// 		}
	// 	}
	// 	if button == Down {
	// 		if floor == 1 {
	// 			Io_set_bit(LIGHT_DOWN2)
	// 		}
	// 		if floor == 2 {
	// 			Io_set_bit(LIGHT_DOWN3)
	// 		}
	// 		if floor == 3 {
	// 			Io_set_bit(LIGHT_DOWN4)
	// 		}
	// 	}
	// }
	fmt.Printf("Setter paa bestillingslys i etasje %d\n", floor)
}

func ClearButtonLed(floor int,button int){

	// if button == Command {
	// 	Io_clear_bit(LIGHT_COMMAND1-floor)
	// }
	// if button == Up {
	// 	if floor == 0 {
	// 		Io_clear_bit(LIGHT_UP1)
	// 	}
	// 	if floor == 1 {
	// 		Io_clear_bit(LIGHT_UP2)
	// 	}
	// 	if floor == 2 {
	// 		Io_clear_bit(LIGHT_UP3)
	// 	}
	// }
	// if button == Down {
	// 	if floor == 1 {
	// 		Io_clear_bit(LIGHT_DOWN1)
	// 	}
	// 	if floor == 2 {
	// 		Io_clear_bit(LIGHT_DOWN2)
	// 	}
	// 	if floor == 3 {
	// 		Io_clear_bit(LIGHT_DOWN3)
	// 	}
	// }
	fmt.Printf("Slaar av lys i etasje %d\n", floor)
}