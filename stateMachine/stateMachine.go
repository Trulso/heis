package stateMachine

import (
	io "../driver"
	"../queue"
	"fmt"
	"time"
)

const (
	IDLE      = 0
	DOOR_OPEN = 1
	MOVING    = 2
)

func Init(floorReachedChan chan int, orderOnSameFloorChan chan int, orderInEmptyQueueChan chan int) {
	fmt.Printf("Initializing the state machine.\n")
	state := IDLE

	doorTimer := time.NewTimer(3 * time.Second)
	doorTimer.Stop()

	fmt.Printf("Initializing complete. Running state machine.\n")
	for {
		select {
		case floor := <-floorReachedChan:
			if floor == -1 {
				queue.LeftFloor("")
			}else{
				fmt.Printf("Vi ankom etasje %d\n", floor)
				io.SetFloorIndicator(floor)
				switch state {
				case MOVING:
					if queue.ShouldStop(floor) {
						io.SetMotorDir(0)
						doorTimer.Reset(3 * time.Second)
						io.SetDoorLamp(1)
						state = DOOR_OPEN
						time.Sleep(time.Millisecond)
						queue.OrderCompleted(floor, "self")
					}
				}
			}
		case floor := <-orderOnSameFloorChan:
			fmt.Println("Det er kommet en ordre i samme etg")
			switch state {
			case IDLE:
				doorTimer.Reset(3 * time.Second)
				io.SetDoorLamp(1)
				state = DOOR_OPEN
				queue.OrderCompleted(floor, "self")
			case DOOR_OPEN:
				doorTimer.Reset(3 * time.Second)
				queue.OrderCompleted(floor, "self")
			}

		case <-orderInEmptyQueueChan:
			fmt.Println("Det er kommet en ordre i en tom ko")
			if state == IDLE {
				io.SetDoorLamp(0)
				direction := queue.NextDirection()
				if direction == 0 {
					state = IDLE
				} else {
					io.SetMotorDir(direction)
					state = MOVING
				}
			}
		case <-doorTimer.C:
			switch state {
			case DOOR_OPEN:
				io.SetDoorLamp(0)
				direction := queue.NextDirection()
				if direction == 0 {
					state = IDLE
				} else {
					io.SetMotorDir(direction)
					state = MOVING
				}
			}
		}
	}
}
