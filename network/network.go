package network 

import (
	"fmt"
	"net"
	"encoding/json"
)


type ElevatorInfo struct { // not done. 
       s int
       b int 
}



func UDPListen(rx chan ElevatorInfo ,port int){
	var temp_struct info
	local,error:= net.ResolveUDPAddr("udp",fmt.Sprintf(":%d", port)) 
	if error !=nil{
		fmt.Println("error:", error)		
	}


	socket,error:=net.ListenUDP("udp",local)
	if error !=nil{
		ffmt.Println("error:", error)
		
	}
	socket.SetReadDeadline(time.Now().Add(10*time.Second)) //ingen aktivitet på net i løpet av 10s, noe er feil ?=

	for{
		buffer := make([]byte,1024)
		n,_,error := socket.ReadFromUDP(buffer) 
		if error !=nil{
			fmt.Println("error:", error)
			//Mulig status chan
		}

		error = json.Unmarshal(buffer[:n], &temp_struct)
		if error != nil {
			fmt.Println("error:", error)
		}
		rx <- temp_struct
		//Mulig status chan
	}

}

func UDPSend(tx chan ElevatorInfo ,port int)  {

	casting,error:= net.ResolveUDPAddr("udp",fmt.Sprintf("255.255.255.255:%d", port)) 
	if error !=nil{
		fmt.Println("error:", error)	
	}
	socket,error:=net.DialUDP("udp",nil,casting)
	if error !=nil{
		fmt.Println("error:", error)	
	}
	socket.SetWriteDeadline(time.Now().Add(10*time.Second))
	for{
		buffer := <- tx
		un_buffer,_ := json.Marshal(buffer)
		_,error = socket.Write(un_buffer)
		if error !=nil{
			fmt.Println("error:", error)
		}
	}
}