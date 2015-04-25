package main

import (
         "fmt"
         "./network"
         //"time"
		."./struct"         
 )


func sendorder(toPass chan Message){

	send := Message{
	MessageType: "newOrder",
	SenderIP: network.GetIP(),
	Elevators: nil,
	ThisFloor: Order{
				Type: -1,
				Floor: 3,
				},
	}

	toPass <- send
	//time.Sleep(5*time.Second)

}

func main() {

	newEle := make (chan string)
	deadEle := make (chan string)

	go network.HeartbeatTransceiver(newEle,deadEle) 

	toPass := make(chan Message)
	toGet := make(chan Message)


	go network.StatusTransceiver(toPass,toGet)


	for{
		fmt.Println("hallo?1")
		select {
			case ele:= <-newEle:	
				fmt.Println("Connected ", ele)
		 	case ele:= <-deadEle:
				fmt.Println("Dead ", ele)
			default: 
				
		 }
		 fmt.Println("hallo?2")
		 temp := <-toGet
		 fmt.Println("hallo?3")
		 fmt.Println(temp.MessageType, "from:",temp.SenderIP)
	}
}
