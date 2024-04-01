package main

import "fmt"

func doARoutineThenBringTheChannel(number int) <-chan int {
	numberStream := make(chan int)

	go func() {
		numberStream <- number
	}()

	return numberStream
}

func main() {

	for i := 0; i < 5; i++ {
		resultStream := doARoutineThenBringTheChannel(i)
		fmt.Println(<-resultStream)
	}

}
