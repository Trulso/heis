package network 

import (
	"fmt"
	"net"
	"encoding/json"
	"time"
	//"os"
)


const(

	elevatorDead = 1000000000
	HeartBeatPort = 30103


)


type Heartbeat struct {
		Id string
		Time time.Time
}


func GetIP() string {

	addrs, error := net.InterfaceAddrs()
	if error != nil {	
    	fmt.Println("error:",error)
    	}

   	for _, address := range addrs {
    	if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
    		if ipnet.IP.To4() != nil {
            	return ipnet.IP.String()
    		}
		}
	}
	return ""
}




func UDPDial(port int) *net.UDPConn {

	casting,error:= net.ResolveUDPAddr("udp",fmt.Sprintf("129.241.187.255:%d", port)) 
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
		time.Sleep(10*time.Millisecond)	

	}
}


func HeartMonitor(newElevator chan string,deadElevator chan string) {
	
	receive := make(chan []byte)
	send := make(chan []byte)
	go UDPRx(receive,HeartBeatPort)
	go UDPTx(send,HeartBeatPort)
	

	heartbeats := make(map[string]*time.Time)
	
	for{	
		myBeat := Heartbeat{GetIP(),time.Now()}
		myBeatBs,error := json.Marshal(myBeat)
		
		if error !=nil{
			fmt.Println("error:", error)
		}
	 	send <- myBeatBs
	 	otherBeatBs := <-receive
	 	
	 	otherBeat := Heartbeat{}
	 	error = json.Unmarshal(otherBeatBs,&otherBeat)
		if error !=nil{
			fmt.Println("error:", error)
		}
		_,ok := heartbeats[otherBeat.Id]
		
		if ok {
			heartbeats[otherBeat.Id] = &otherBeat.Time
		}else{
			newElevator <- otherBeat.Id
			heartbeats[otherBeat.Id] = &otherBeat.Time 
			
		}

		for i,t := range heartbeats {
			fmt.Println(i)
			fmt.Println("\n \n")
			dur := time.Since(t)
			fmt.Println(dur.Nanoseconds()) 
			fmt.Println("before dur")
			if dur.Nanoseconds() > 300000000000 {
				fmt.Println("why u go in here")
				deadElevator <- i
				delete(heartbeats,i)
			}
		}
		time.Sleep(10*time.Millisecond)
	}
}

func StatusMonitor(){



}

func AcknowledgeMonitor(){


}