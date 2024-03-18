package channels

import (
	"fmt"
	"sync"
	"time"
)

// The algorithm of a channel right from the class study
func Channel() {
	// the first thing to be aware is what a channel is
	// a channel is a way to use sync data of go routines
	// the way he expect to work is to have two basic operations in the channel
	// they are write and read
	// a function writes some type of thing in the channel
	// the rest of the program expect to the channel receives some data
	// by the end of the comunication, we suppost to close the channel.
	// the write system is channel <- item
	// the read system is <-channel
	// the result of the attribution of a channel is the item and the state (open or closed)
	// we can do a for loop in a channel
	// we can do a range in a channel
	// the range seems more easy to write
	// but the for is more clean code accurate cause checks for open/closed

	// first of all we start the algorith with a go routine
	lettersStream := make(chan string)

	fmt.Println("The begining of the program")
	go writeLetters(lettersStream)

	// now we need to wait for the channel receive a message
	// just wait for the channel will not give time enough to do all operations
	// so we need to loop over the items too.

	for {
		message, isOpen := <-lettersStream
		if !isOpen {
			break
		}
		fmt.Println(message)
	}
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(1)
	waitGroup.Done()
	waitGroup.Wait()
	// but if the code above just do a for in a channel algorithm then there will be a time
	// where the channel is not anymore waiting for nothing and we still in the for.
	// so we need to close the channel by the end of the go routine.

	fmt.Println("The end of the program")
	// ! IT IS IMPORTANT TO NOTE THAT DEADLOCK'S ARE ONLY AVAILABLE TO DEBUG IN RUNTIME
	// ! THAT MEANS CODE IN PRODUCTION WITH DEAD LOCK CERTAINLY WILL BROKE THE APPLICATION

}

func writeLetters(channel chan string) {
	lettersSlice := []string{"A", "B", "C"}

	for _, letter := range lettersSlice {
		time.Sleep(time.Second * 1)
		channel <- string(letter)
	}

	// ! HERE WE NEED TO CLOSE TO AVOID A FOR LOOP THAT CAN END IN A DEADLOCK
	close(channel)

}
