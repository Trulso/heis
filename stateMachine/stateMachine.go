package statemachine


import(
	"fmt"
	"../driver"
	"../queue"
	"time"
)

const (	
	IDLE				= 0
	DOOR_OPEN			= 1
	MOVING				= 2
	
	UP 					= 1
	DOWN 				= -1
	STOP				= 0
)
//state
//


/*Channels
Floor
NewOrder


*/

func Init(FloorReached chan int, NewOrder chan int){
	fmt.Printf("Initializing the state machine.\n")
	state := IDLE

	doorTimer := time.New_Timer(3*time.Seconds)
	doorTimer.Stop()




	fmt.Printf("Initializing complete. Running state machine.\n")
	for {
		select {
		case floor := <-FloorReached:
			driver.SerFloorIndicator(floor)
			switch state {
			case MOVING:
				if queue.ShouldStop(floor) {
					driver.SetMotorDir(0)
					state = DOOR_OPEN
				}
			}
		case  direction := <- NewOrder:
			switch state {
			case IDLE:
			}
		case <- doorTimer.C:
			switch state {
			case DOOR_OPEN:
				
			}
		}
	}
}