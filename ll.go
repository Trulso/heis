package main
	

import (
	"fmt"
	"./network"
	"time"
	."./struct"
)

func main() {
	
	toPass := make(chan Message)
	toGet := make(chan Message)


	go network.StatusTransceiver(toPass,toGet)


	for{
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
			time.Sleep(20*time.Second)

		}
}

/*
type Message struct {
	MessageType string //neworder,just arrived, status update, completed order,
	SenderIP    string
	Elevators   map[string]Elevator
	ThisFloor   Order
}

type Order struct {
	Type  int
	Floor int
}
*/