package main

import (
	"fmt"
	"time"
)

func main() {

	c := make(chan bool)

	go func() {
		// Do something
		time.Sleep(time.Second * 5)
		close(c)
	}()

	// Wait a message from channel or when it is close.
	<-c

	fmt.Println("Done")
}
