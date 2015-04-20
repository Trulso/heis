package main


import(
	"fmt"
	"time"


)

func main(){
	fmt.Print("Hei\n")

	doorTimer := time.NewTimer(3*time.Second)
	<-doorTimer.C
	fmt.Print("Why deadlock\n")
	doorTimer.Stop()
	for{
		<-doorTimer.C
		fmt.Print("Det har gaatt tre sekunder\n")
		doorTimer.Reset(3 * time.Second)
	}
}
