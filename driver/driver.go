package driver

import (
	"fmt"
	"time"
)

const (
	N_FLOORS=4

	COMMAND = 2
	UP   = 1
	DOWN = -1
)



func Init() int {

	status := Io_init()
	if (status == 0) {
		fmt.Println("Hardware init failed!")
		return -1
	}

	for i := 0; i < 16; i++ {
		Io_clear_bit(0x300+i)
	}

	SetMotorDir(-1)
Loop:
	for {
		for floor := SENSOR_FLOOR1; floor < SENSOR_FLOOR4+1; floor++ {
			if Io_read_bit(floor) == 1 {
				SetMotorDir(0)
				SetFloorIndicator(floor-SENSOR_FLOOR1)
				break Loop
			}
		time.Sleep(1*time.Millisecond)
		}
	}


	fmt.Println("Hardware init success.")
	return 1
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
}//Ferdig

func SetButtonLed(floor int,button int){
	if(floor<=N_FLOORS){
		if button == COMMAND {
			Io_set_bit(LIGHT_COMMAND1-floor)
		}
	}	
	if button == UP {
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
	if button == DOWN {
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
}//Ferdig

func ClearButtonLed(floor int,button int){
	if button == COMMAND {
		Io_clear_bit(LIGHT_COMMAND1-floor)
	}
	if button == UP {
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
	if button == DOWN {
		if floor == 1 {
			Io_clear_bit(LIGHT_DOWN2)
		}
		if floor == 2 {
			Io_clear_bit(LIGHT_DOWN3)
		}
		if floor == 3 {
			Io_clear_bit(LIGHT_DOWN4)
		}
	}
}//Ferdig

func SetDoorLamp(value int) {
	if (value) > 0 {
	    Io_set_bit(LIGHT_DOOR_OPEN)
	}else{
        Io_clear_bit(LIGHT_DOOR_OPEN)
    }	
}//Ferdig

func SetStopLamp(value int) {
	if (value) > 0{
        Io_set_bit(LIGHT_STOP)
	}else{
	    Io_clear_bit(LIGHT_STOP)
	}
}

func GetObstructionSignal() int {
	return Io_read_bit(OBSTRUCTION)
}

func GetStopSignal() int {
	return Io_read_bit(STOP)
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
}//Ferdig

func OrderButtonPolling(commandOrdersChan chan int,upOrdersChan chan int, downOrdersChan chan int){
	fmt.Println("Na starter pollinga av buttons")
	buttons := [3*N_FLOORS-2]int{BUTTON_COMMAND1,BUTTON_COMMAND2,BUTTON_COMMAND3,BUTTON_COMMAND4,BUTTON_UP1,BUTTON_UP2,BUTTON_UP3,BUTTON_DOWN2,BUTTON_DOWN3,BUTTON_DOWN4}
	buttonPressed := make([]bool, 3*N_FLOORS-2)
	for {
		for button := 0; button < len(buttons); button++ {
			if Io_read_bit(buttons[button]) == 1{
				if !buttonPressed[button] {
					buttonPressed[button] = true
					if button < N_FLOORS {
						commandOrdersChan <- button
					}else if button < 2*N_FLOORS-1 {
						upOrdersChan <- button - N_FLOORS
					}else {
						downOrdersChan <- button - (2*N_FLOORS-2)
					}
				}
			}else if buttonPressed[button]{
				buttonPressed[button] = false
			}
		}
		time.Sleep(10*time.Millisecond)		
	}
}//Ferdig

func FloorSensorPolling(floorSensorChan chan int){
	fmt.Println("Naa starter pollinga av etg")
	pushed := make([]bool, N_FLOORS)
	for {
		for floor := SENSOR_FLOOR1; floor < SENSOR_FLOOR4+1; floor++{
			if Io_read_bit(floor) == 1{
				if !pushed[floor-SENSOR_FLOOR1]{
					pushed[floor-SENSOR_FLOOR1] = true
					floorSensorChan <- floor-SENSOR_FLOOR1
				}
			}else if pushed[floor-SENSOR_FLOOR1] {
				pushed[floor-SENSOR_FLOOR1] = false
			}
		}
		time.Sleep(10*time.Millisecond)
	}
}//Ferdig

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