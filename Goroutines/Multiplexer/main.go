package main

import (
	"fmt"
	"math/rand"

	"time"
)

func sleepRandom() {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))

}

// generator
func sleepThenPrint(word string) <-chan string {
	messageStream := make(chan string)
	go func() {
		for {
			messageStream <- word
			sleepRandom()
		}
	}()
	return messageStream
}
func multiplexer(messageStream1, messageStream2 <-chan string) <-chan string {
	resultMessageStream := make(chan string)

	go func() {
		for {
			select {
			case message := <-messageStream1:
				resultMessageStream <- message

			case message := <-messageStream2:
				resultMessageStream <- message
			}
		}
	}()

	return resultMessageStream
}

func main() {
	message := multiplexer(sleepThenPrint("Hello"), sleepThenPrint("World"))

	for {
		fmt.Println(<-message)
	}

}
