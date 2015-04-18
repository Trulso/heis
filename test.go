package main

import (
	"fmt"
	"./driver"
)

const (

	A = iota
	B
	C
	D
)


func main(){
	driver.HwInit()
	
	fmt.Printf("%d",A)
	fmt.Printf("%d",B)
	fmt.Printf("%d",C)
	fmt.Printf("%d",D)
	for {
		if driver.Io_read_bit(driver.SENSOR_FLOOR4) == 1{
			driver.Io_set_bit(driver.MOTORDIR)
			driver.Io_write_analog(driver.MOTOR, 2800)
		}
		if driver.Io_read_bit(driver.SENSOR_FLOOR1) == 1{
			driver.Io_clear_bit(driver.MOTORDIR)
			driver.Io_write_analog(driver.MOTOR, 2800)
		}		
	}
}
