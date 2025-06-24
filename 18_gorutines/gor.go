package main 

import (
	"fmt"
	"time"
)

func sendEmail(orderID int) {	
	fmt.Println("Email sent for order:", orderID)
}

func sendData( dt string  , chanl chan string )  {
	chanl <- dt 
}

func main(){
    fmt.Println("Placing Order...")
	orderID := 101
    
    go sendEmail(orderID)
	// it creates a thread if we not wait main func will complete and we cant see  send email 
	// try comemnting time sleep 
	// and with out comment sleep
	time.Sleep(3 * time.Second) 
	fmt.Println("Order placed successfully!")


	// --- about the channels 

	mychannel := make(chan string )

	go func(){
		mychannel <- "data"
	}()

	msg := <- mychannel

	fmt.Println(msg)

    // many thread sinfgle channel 

	ch := make( chan string )
    
	go sendData( "Message from goroutine 1" , ch)
	go sendData( "Message from goroutine 2" , ch )
	go sendData( "Message from goroutine 3" , ch )

	for i :=0 ; i<3 ; i++ {
      msg := <- ch 
	  fmt.Println(msg)
	} 

	time.Sleep(time.Second)

	//  if channel size is given it is bufered channel ie there is space revieved 

	// else it is unbuffere channe; 

	// unbuffered channel will allow the go routuine to fill data only when there is a reciever 
    
	// buffered channel will get  the data of its size and then  will allow the other data to fill 
	//  when  the existing data is read 

	ch3 := make(chan string , 3 )

    go sendData( "Message from goroutine ch3 4" , ch3)
	go sendData( "Message from goroutine ch3 5" , ch3 )
	go sendData( "Message from goroutine ch3 6" , ch3)	
	// go sendData( "Message from goroutine ch3 4" , ch3)	
	time.Sleep(time.Second)
	close(ch3)

	for res := range ch3{
		fmt.Println(res)
	}

	time.Sleep(time.Second)


}