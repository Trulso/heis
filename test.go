package main

import (
	. "fmt"
	"runtime"
	"time"
	"./driver"
)

func main(){
	driver.Io_init()
	driver.Io_clear_bit(driver.LIGHT_DOOR_OPEN)
}
