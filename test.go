package main

import (
         "fmt"
         "./network"
         "time"
		."./struct"         
 )


func main() {

	newEle := make (chan string)
	deadEle := make (chan string)

	go network.HeartbeatTransceiver(newEle,deadEle) 

	toPass := make(chan Message)
	toGet := make(chan Message)


	go network.StatusTransceiver(toPass,toGet)


	for{
		select {
			case ele:= <-newEle:	
				fmt.Println("Connected ", ele)
		 	case ele:= <-deadEle:
				fmt.Println("Dead ", ele)
			default: 

		 }
		send := Message{
		MessageType: "newOrder",
		SenderIP: network.GetIP(),
		Elevators: nil,
		ThisFloor: Order{
					Type: -1,
					Floor: 4,
					},
		}
		toPass <- send
		fmt.Println("Sendt!")
		time.Sleep(1*time.Second)

	}
}
