// go run Oppg2.go
package main

import (
	. "fmt"
	"runtime"
	"time"
)


var i int = 0

func Inkrementer(sem chan int) {
	for x := 0; x<1000000; x++{
		i= <- sem
		i= i+2
		sem <- i
	}	
}

func Dekrementer(sem chan int) {
	for x := 0; x<1000000; x++{
		i= <- sem
		i--
		sem <- i
	}
}

func main() {
	sem := make(chan int, 1)
	sem <- i
	runtime.GOMAXPROCS(runtime.NumCPU()) 
	go Inkrementer(sem)
	go Dekrementer(sem)
	time.Sleep(900*time.Millisecond)
	//i= <-sem
	Println(i)
}

