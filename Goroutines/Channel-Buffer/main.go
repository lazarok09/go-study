package main

import "fmt"

func main() {
	// make a channel
	messagesStream := make(chan string, 2)
	messagesStream <- "Hello"
	messagesStream <- "Wolrd"

	fmt.Println(<-messagesStream)
	fmt.Println(<-messagesStream)
}
