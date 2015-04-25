package struckt 

import (
		"time"
)



type Heartbeat struct {
		Id string
		Time time.Time
}

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

type Elevator struct {
	Direction       int
	LastPassedFloor int
	UpOrders        []bool
	DownOrders      []bool
	CommandOrders   []bool
}


