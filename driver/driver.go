package driver

import (
	"fmt"
)

const (
	Up   = 1
	Down = -1
	Stop = 0
	On   = 1
	Off  = 0
)



func HwInit() int{

	status := Io_init()
	if (status == 0) {
		fmt.Printf("Hw init failed!")
		return -1
	}
	

	for i := 0; i < 16; i++{
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









