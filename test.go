package main

import (
         "fmt"
         "./network"
         "time"
		."./struct"         
 )


func sendorder(toPass chan Message){
	for{
		send := Message{
		MessageType: "newOrder",
		SenderIP: network.GetIP(),
		ReceiverIP: "",
		Elevators: nil,
		ThisFloor: Order{
					Type:  1,
					Floor: 2,
					},
		}
		time.Sleep(10*time.Second)
		toPass <- send
		
	}
}

func main() {

	newEle := make (chan string)
	deadEle := make (chan string)

	go network.HeartbeatTransceiver(newEle,deadEle) 

	toPass := make(chan Message)
	toGet := make(chan Message)

	go sendorder(toPass)
	go network.StatusTransceiver(toPass,toGet)


	for{
		select {
			case ele:= <-newEle:	
				fmt.Println("Connected ", ele)
		 	case ele:= <-deadEle:
				fmt.Println("Dead ", ele)
			case temp := <-toGet:
				 fmt.Println(temp.MessageType,"from",temp.SenderIP)
			default: 

		 }
	}
}
