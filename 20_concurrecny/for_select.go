package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)


	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "Message from ch1"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "Message from ch2"
	}()
    
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("Received:", msg1)
		case msg2 := <-ch2:
			fmt.Println("Received:", msg2)
		}
	}

  	ch := make(chan string, )
    go func() {
		time.Sleep(3 * time.Second)
		ch <- "data from channel"
	}()

	select {
	case msg := <- ch :
		fmt.Println(msg)
	case <- time.After(2 * time.Second ):
		fmt.Println("time out ")
	}

	

}