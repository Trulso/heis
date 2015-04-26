package cost



	var topOrder int

	var lowestFloor int 

	if isQueueEmpty(ip){
		topOrder = newOrder.Floor
	}
	if isQueueEmpty(ip){
		lowestFloor = newOrder.Floor
	}

	//for i := 0; i < 5; i++
	fmt.Println("Elevator Direction: ",elevator.Direction)
	fmt.Println("Order Type",newOrder.Type)
	if elevator.LastPassedFloor < newOrder.Floor{


		if elevator.Direction == UP || elevator.Direction == 0 {
			fmt.Println("GÅR INN I ELEVATOR UP")
			if newOrder.Type == UP {
				fmt.Println("Elevator: UP .. Order: UP")
				cost += newOrder.Floor - elevator.LastPassedFloor
				for i := elevator.LastPassedFloor; i < newOrder.Floor; i++ {
					if elevator.UpOrders[i] == true || elevator.CommandOrders[i] == true{
						cost += 1
					}
				return cost	
				}
			}
			if newOrder.Type == DOWN {
				fmt.Println("Elevator: UP .. Order: Down")
				if elevator.CommandOrders[newOrder.Floor] == true || elevator.UpOrders[newOrder.Floor] == true{
					cost += newOrder.Floor - elevator.LastPassedFloor
					for i := elevator.LastPassedFloor; i < newOrder.Floor; i++ {
						if elevator.UpOrders[i] == true || elevator.CommandOrders[i] == true{
							cost += 1
						}
					}	
					return cost
				}else{
					
					for i := elevator.LastPassedFloor; i < N_FLOORS; i++{
						if elevator.UpOrders[i] == true || elevator.CommandOrders[i] == true  {
							cost += 1
							topOrder=i
						}else if elevator.DownOrders[i] == true{
							topOrder=i

						}				
					}

					for i := topOrder; i >= newOrder.Floor; i-- {
						if elevator.DownOrders[i] == true {
							cost += 1
						}

					}
					cost += topOrder - elevator.LastPassedFloor
					cost += topOrder - newOrder.Floor
					return cost
				} 
			}
		}
		if elevator.Direction == DOWN  || elevator.Direction == 0 { {
	
			fmt.Println("GÅR INN I ELEVATOR DOWN")
			for i := elevator.LastPassedFloor; i >= 0; i-- {
				if elevator.DownOrders[i] == true || elevator.CommandOrders[i] == true{
					cost += 1
					lowestFloor = i
				}				
			}
			cost += elevator.LastPassedFloor - lowestFloor
			}

			if newOrder.Type == UP {
				fmt.Println("Elevator: DOWN .. Order: UP")
				for i := lowestFloor; i < newOrder.Floor; i++{
					if elevator.UpOrders[i] == true {
						cost += 1
					}
				}
				cost += newOrder.Floor - lowestFloor
				
				return cost
			}
		
			if newOrder.Type == DOWN {
				fmt.Println("Elevator: DOWN .. Order: DOWN")
				if elevator.CommandOrders[newOrder.Floor] == true{

					for i := lowestFloor; i < newOrder.Floor; i++{
						if elevator.UpOrders[i] == true {
							cost += 1
						}
					cost += newOrder.Floor - lowestFloor
					}
					return cost

				}
				for i := 0; i < N_FLOORS; i++ {
					if elevator.UpOrders[i] == true {
						cost += 1
						topOrder = i
					}else if elevator.DownOrders[i] == true {
						topOrder = i
					}else{
						topOrder = newOrder.Floor
					}
				}
				fmt.Println("TOP ORDER :>-----",topOrder)
				fmt.Println("LOWESST FLOOR :>----",lowestFloor)
				cost += topOrder - lowestFloor
				cost += topOrder - newOrder.Floor
				return cost
				// burda EGETNLIG sjekka down orders på veien ned igjen meeeeeeeeeeeeeeeeen.
						
			}
			
		}
	}

	if elevator.LastPassedFloor > newOrder.Floor{  // ONE COST TO BIND THEM ALL

		fmt.Println("THE FLOOR IS UNDER")

		if elevator.Direction == UP || elevator.Direction == 0 {
			fmt.Println("GÅR INN I ELEVATOR UP")
			if newOrder.Type == UP {
				fmt.Println("hallo")		
			}
			if newOrder.Type == DOWN {
				fmt.Println("Elevator: UP .. Order: Down")
			}
		}
		if elevator.Direction == DOWN  || elevator.Direction == 0 { 
	
			fmt.Println("GÅR INN I ELEVATOR DOWN")


			if newOrder.Type == UP {
				fmt.Println("Elevator: DOWN .. Order: UP")
				if elevator.CommandOrders[newOrder.Floor] == true || elevator.DownOrders[newOrder.Floor] == true{
					cost += elevator.LastPassedFloor - newOrder.Floor
					for i := elevator.LastPassedFloor; i > newOrder.Floor; i-- {
						if elevator.DownOrders[i] == true || elevator.CommandOrders[i] == true{
							cost += 1
						}
					}	
					return cost
				}else{ 
					
					for i := elevator.LastPassedFloor; i < N_FLOORS; i++{
						if elevator.UpOrders[i] == true || elevator.CommandOrders[i] == true  {
							cost += 1
							topOrder=i
						}else if elevator.DownOrders[i] == true{
							topOrder=i

						}				
					}

					for i := topOrder; i >= newOrder.Floor; i-- {
						if elevator.DownOrders[i] == true {
							cost += 1
						}

					}
					cost += topOrder - elevator.LastPassedFloor
					cost += topOrder - newOrder.Floor
					return cost
				} 
		
						
			}
		
			if newOrder.Type == DOWN {
				fmt.Println("Elevator: DOWN .. Order: DOWN")
				cost += elevator.LastPassedFloor - newOrder.Floor 
				for i := elevator.LastPassedFloor; i >= newOrder.Floor; i-- {
					if elevator.UpOrders[i] == true || elevator.CommandOrders[i] == true{
						cost += 1
					}
				return cost	
				}		
			}
			
		}	
	}
fmt.Println("finner ikke kost!!!!!!!!!!!!!!!!!!!!!!!!!!")
return 1
} //Ikke laget ennå.