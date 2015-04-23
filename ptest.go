package main

import(
	"fmt"
	"time"
	"os/exec"
	"net"
	"encoding/json"
)

type info struct{
	Nr int
}



func UDPListen() (info,int){
	var rx info
	local,error:= net.ResolveUDPAddr("udp",fmt.Sprintf(":%d", 30003)) 
	if error !=nil{
		fmt.Println("error:", error)		
	}


	socket,error:=net.ListenUDP("udp",local)
	socket.SetReadDeadline(time.Now().Add(2*time.Second))
	if error !=nil{
		fmt.Println("error:", error)
		socket.Close()
		return rx,0
	}
	for{
		buffer := make([]byte,1024)
		n,_,error := socket.ReadFromUDP(buffer) 
		if error !=nil{
			fmt.Println("error:", error)
			socket.Close()
			return rx,0
		}
		if error !=nil{
			fmt.Println("error:", error)	
		}
		error = json.Unmarshal(buffer[:n], &rx)
		if error != nil {
			fmt.Println("error:", error)
		}
		socket.Close()
		return rx,1
	}

}

func UDPSend(data info)  {

	casting,error:= net.ResolveUDPAddr("udp",fmt.Sprintf("129.241.187.255:%d", 30003)) 
	if error !=nil{
		fmt.Printf("Error Resolving UDP")	
	}
	socket,error:=net.DialUDP("udp",nil,casting)
	if error !=nil{
		fmt.Printf("Error Dialing to UDP adress")
		
	}

	un_buffer,_ := json.Marshal(data)
	_,error = socket.Write(un_buffer)
		if error !=nil{
		fmt.Printf("Error writing to UDP adress")
	}
}

func Spawn(){
	
	cmd := exec.Command("gnome-terminal", "-x", "go", "run", "ptest.go")
	cmd.Start()
	

}


func slave(masterAlive int)int{
	var s_count int
	for(masterAlive ==1 ){
		send := info{}
		send,masterAlive = UDPListen()
		

		if masterAlive == 0 {
			return s_count
		} else {
			s_count=send.Nr
		}
	}
	return s_count
}

func main(){
	var tx info
	masterAlive := 1
	counter := 0


	counter=slave(masterAlive)

	Spawn()

	for{



		UDPSend(tx)
		counter++
		tx.Nr=counter		
		fmt.Printf("Master Counter: %d \n",tx.Nr)
		time.Sleep(time.Second)

	}
}