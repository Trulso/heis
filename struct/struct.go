package structs

import (
	"time"
)

type Heartbeat struct {
	Id   string
	Time time.Time
}

type Message struct {
	MessageType string
	SenderIP    string
	TargetIP    string //Which elevator that changes status
	Elevator    Elevator
	Order       Order
}

type Order struct {
	Type  int
	Floor int
}

type Elevator struct {
	Active          bool
	InFloor         bool
	Direction       int
	LastPassedFloor int

	UpOrders      []bool
	DownOrders    []bool
	CommandOrders []bool
}
