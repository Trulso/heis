package main

import (
         "fmt"
         "./network"
         ."time"
 )

func main() {

	newEle := make (chan string)
	deadEle := make (chan string)

	go network.HeartMonitor(newEle,deadEle) 

	for{
		fmt.Println("Connected ", <-newEle)

		fmt.Println("Dead ",<-deadEle)
		Sleep(3*Second)
	}


}
