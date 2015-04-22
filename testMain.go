package main


import(
	"fmt"
	"time"


)

func main(){
	fmt.Print("Hei\n")

	doorTimer := time.NewTimer(1000*time.Millisecond)
	<-doorTimer.C
	fmt.Print("Why deadlock\n")
	doorTimer.Stop()
	for{
		doorTimer.Reset(1 * time.Second)
		<-doorTimer.C
		fmt.Print("Det har gaatt ett sekund\n")
		
	}
}
