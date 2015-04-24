package main

import (
	"fmt"
	"./driver"
	"time"
)



func main(){
	driver.HwInit()
	driver.SetDoorLamp(1)
	fmt.Print("hei\n")
	for{
		if driver.GetStopSignal() == 1 {
			driver.SetStopLamp(1)
			time.Sleep(2*time.Second)
		} else {
			driver.SetStopLamp(0)
		}
		time.Sleep(100*time.Millisecond)
	}
}