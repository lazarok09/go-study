package main

import (
	"fmt"
	"time"
)

// T1
func main() {
	operationNameStream := make(chan string)
	operationValueStream := make(chan int)

	// create a rotine

	go func() {
		for {

			time.Sleep((time.Millisecond * 500))
			operationNameStream <- "Item 1"
		}
	}()

	go func() {
		for {

			time.Sleep((time.Millisecond * 500))
			time.Sleep((time.Millisecond * 500))
			operationValueStream <- 20
		}
	}()
	for {
		select {
		case name := <-operationNameStream:
			fmt.Println(name)
		case value := <-operationValueStream:
			fmt.Println(value)
		}

	}

}
