package network 

import (
	"fmt"
	"net"
	"encoding/json"
)


type Dummy struct {
       s int
       b int
}



func UDPListen(data chan Dummy ,port int) {

	local,error:= net.ResolveUDPAddr("udp",fmt.Sprintf(":%d", port)) 
	if error !=nil{
		fmt.Printf("Error Resolving UDP")


		
	}


	socket,error:=net.ListenUDP("udp",local)
	if error !=nil{
		fmt.Printf("Error Listening to local adress")
		
	}
	for{
		dummy_st  := Dummy{}
		buffer := make([]byte,1024)
		_,_,error := socket.ReadFromUDP(buffer) 
		if error !=nil{
			fmt.Printf("Error Reading from UDP")
			
		}
		json.Unmarshal(buffer, &dummy_st)
		data <- dummy_st 
	}

}

func UDPSend(data chan Dummy ,port int)  {

	casting,error:= net.ResolveUDPAddr("udp",fmt.Sprintf("255.255.255.255:%d", port)) 
	if error !=nil{
		fmt.Printf("Error Resolving UDP")	
	}
	socket,error:=net.DialUDP("udp",nil,casting)
	if error !=nil{
		fmt.Printf("Error Dialing to UDP adress")
		
	}

	buffer := <- data
	un_buffer,_ := json.Marshal(buffer)
	_,error = socket.Write(un_buffer)
		if error !=nil{
		fmt.Printf("Error writing to UDP adress")
	}
}