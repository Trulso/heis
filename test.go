package main

import (
         "fmt"
         "./network"
         //"time"
 )

func main() {

	newEle := make (chan string)
	deadEle := make (chan string)

	go network.HeartMonitor(newEle,deadEle) 


	for{
		select {
		case ele:= <-newEle:
			
			fmt.Println("Connected ", ele)
		case ele:= <-deadEle:
			fmt.Println("Dead ", ele)
		}
		fmt.Println("Inne i for-løkken")
	}

}
