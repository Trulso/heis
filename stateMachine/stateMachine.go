package stateMachine

import (
	io "../fakeDriver"
	"../queue"
	"fmt"
	"time"
	//io "../driver"
)

const (
	IDLE      = 0
	DOOR_OPEN = 1
	MOVING    = 2

	UP   = 1
	STOP = 0
	DOWN = -1
	
)

//state
//

/*Channels
Floor
NewOrder


*/

func Init(FloorReached chan int, NewOrder chan int) {
	fmt.Printf("Initializing the state machine.\n")
	state := IDLE

	doorTimer := time.NewTimer(3 * time.Second)
	doorTimer.Stop()

	fmt.Printf("Initializing complete. Running state machine.\n")
	for {
		select {

		case floor := <-FloorReached:
			io.SetFloorIndicator(floor)
			switch state {
			case MOVING:
				if queue.ShouldStop(floor) {
					io.SetMotorDir(0)
					doorTimer.Reset(3 * time.Second)
					io.SetDoorLamp(1)
					state = DOOR_OPEN
					queue.OrderCompleted(floor)
				}
			}
		case direction := <-NewOrder:
			switch state {
			case IDLE:
			}
		case <-doorTimer.C:
			switch state {
			case DOOR_OPEN:
				io.SetDoorLamp(0)
				order := queue.NextDirection()
				if order == 0 {
					
				}

			}
		}
	}
}
