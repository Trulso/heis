package network 

import (
	"fmt"
	"net"
	//"encoding/json"
	"time"
)



type ElevatorInfo struct {
		Id int
		Tx string
}


func UDPDial(port int) *net.UDPConn {

	casting,error:= net.ResolveUDPAddr("udp",fmt.Sprintf("255.255.255.255:%d", port)) 
	if error !=nil{
		fmt.Println("error:", error)	
	}
	socket,error:=net.DialUDP("udp",nil,casting)
	if error !=nil{
		fmt.Println("error:", error)
	}
	return socket
}


func UDPListen(port int) *net.UDPConn {
	local,error:= net.ResolveUDPAddr("udp",fmt.Sprintf(":%d", port)) 
	if error !=nil{
		fmt.Println("error:", error)		
	}


	socket,error:=net.ListenUDP("udp",local)
	if error !=nil{
		fmt.Println("error:", error)	
	}
	return socket
}


func UDPRx(rx chan []byte ,port int){
	
	socket := UDPListen(port)

	for{
		socket.SetReadDeadline(time.Now().Add(10*time.Second)) //ingen aktivitet på net i løpet av 10s, noe er feil ?=
		buffer := make([]byte,1024)
		n,_,error := socket.ReadFromUDP(buffer) 
		
		if error !=nil{
			fmt.Println("error:", error)		
		}
		
		buffer = buffer[:n]
		rx <- buffer
	}

}

func UDPTx(tx chan []byte,port int)  {
	
	socket := UDPDial(port)

	for{
		socket.SetWriteDeadline(time.Now().Add(10*time.Second))
		_,error := socket.Write(<- tx)
		if error !=nil{
			fmt.Println("error:", error)
		}
	}
}

func HeartMonitor() {

	//un_buffer,_ := json.Marshal(buffer)


	//error = json.Unmarshal(buffer[:n], &temp_struct)
	//if error != nil {
	//	fmt.Println("error:", error)
	//}

}

func StatusMonitor(){



}

func AcknowledgeMonitor(){


}