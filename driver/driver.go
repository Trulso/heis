package driver

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

	status := Io_init()
	if (status == 0) {
		fmt.Printf("Hw init failed!")
		return -1
	}
	

	for i := 0; i < 16; i++ {
		Io_clear_bit(0x300+i)
	}

	fmt.Printf("Hw init sucsess.")
	return 1
}


func GetFloorSensorSignal() int {
	if Io_read_bit(SENSOR_FLOOR1) == 1{
    	return 0
	}	
	if Io_read_bit(SENSOR_FLOOR2) == 1{
	    return 1
	}
	if Io_read_bit(SENSOR_FLOOR3) == 1{
	    return 2
	}
	if Io_read_bit(SENSOR_FLOOR4) == 1{
	    return 3
	}else{
	    return -1
	}
}

func GetObstructionSignal() int {
	return Io_read_bit(OBSTRUCTION)
}

func GetStopSignal() int {
	return Io_read_bit(STOP)
}

func SetDoorLamp(value int) {
	if (value) > 0 {
	    Io_set_bit(LIGHT_DOOR_OPEN)
	}else{
        Io_clear_bit(LIGHT_DOOR_OPEN)
    }	
}

func SetStopLamp(value int) {
	if (value) > 0{
        Io_set_bit(LIGHT_STOP)
	}else{
	    Io_clear_bit(LIGHT_STOP)
	}
}

func SetMotorDir(dir int) {
	if (dir == 0){
	    Io_write_analog(MOTOR, 0)
	}
	if (dir > 0) {
       	Io_clear_bit(MOTORDIR)
        Io_write_analog(MOTOR, 2800)
	}
    	if (dir < 0) {
    	Io_set_bit(MOTORDIR)
    	Io_write_analog(MOTOR, 2800)
    }
}

func getUpOrdersSignal(floor int) int {
	if floor == 0 {
		return Io_read_bit(BUTTON_UP1)
	}else if floor == 1 {
		return Io_read_bit(BUTTON_UP2)
	}else if floor == 2 {
		return Io_read_bit(BUTTON_UP3)
	}
	fmt.Printf("No Up orderbuttons exist over the 3rd floor.\n")
	return -1
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
	if floor == 1 {
		return Io_read_bit(BUTTON_DOWN2)
	}else if floor == 2 {
		return Io_read_bit(BUTTON_DOWN3)
	}else if floor == 3 {
		return Io_read_bit(BUTTON_DOWN4)
	}
	fmt.Printf("No down orderbuttons exist bellow the 1rd floor.\n")
	return -1
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
	if floor == 0 {
		return Io_read_bit(BUTTON_COMMAND1)
	}else if floor == 1 {
		return Io_read_bit(BUTTON_COMMAND2)
	}else if floor == 2 {
		return Io_read_bit(BUTTON_COMMAND3)
	}else if floor == 3 {
		return Io_read_bit(BUTTON_COMMAND4)
	}
	fmt.Printf("No Command order buttons exist outside floors 1-4\n")
	return -1
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
    if (floor & 0x02) != 0 {
        Io_set_bit(LIGHT_FLOOR_IND1)
    }else {
        Io_clear_bit(LIGHT_FLOOR_IND1)
    }

    if (floor & 0x01) != 0 {
        Io_set_bit(LIGHT_FLOOR_IND2)
    }else {
        Io_clear_bit(LIGHT_FLOOR_IND2)
    }
}

func SetButtonLed(floor int,button int){
	if(floor<=N_FLOORS){
		if button == Command {
			Io_set_bit(LIGHT_COMMAND1-floor)
		}
	}
	
	if button == Up {
		if floor == 0 {
			Io_set_bit(LIGHT_UP1)
		}
		if floor == 1 {
			Io_set_bit(LIGHT_UP2)
		}
		if floor == 2 {
			Io_set_bit(LIGHT_UP3)
		}
	}
	if button == Down {
		if floor == 1 {
			Io_set_bit(LIGHT_DOWN2)
		}
		if floor == 2 {
			Io_set_bit(LIGHT_DOWN3)
		}
		if floor == 3 {
			Io_set_bit(LIGHT_DOWN4)
		}
	}
}

func ClearButtonLed(floor int,button int){

	if button == Command {
		Io_clear_bit(LIGHT_COMMAND1-floor)
	}
	if button == Up {
		if floor == 0 {
			Io_clear_bit(LIGHT_UP1)
		}
		if floor == 1 {
			Io_clear_bit(LIGHT_UP2)
		}
		if floor == 2 {
			Io_clear_bit(LIGHT_UP3)
		}
	}
	if button == Down {
		if floor == 1 {
			Io_clear_bit(LIGHT_DOWN1)
		}
		if floor == 2 {
			Io_clear_bit(LIGHT_DOWN2)
		}
		if floor == 3 {
			Io_clear_bit(LIGHT_DOWN3)
		}
	}
}
