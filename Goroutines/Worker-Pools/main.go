package main

import "fmt"

func fibonnaci(position int) int {
	if position <= 1 {
		return position
	}
	return fibonnaci(position-2) + fibonnaci(position-1)
}
func main() {

	taskStream := make(chan int, 40)
	resultStream := make(chan int, 40)
	// the go routine goes outside the below code and
	// stop ate range that is waiting somenthing to be in the taskStream
	// because there is a for there and he is iterating a channel stream
	// which has the behavior to pause the rest of execution
	// this is why we need to close the task stream channel
	// after all our code is done

	go worker(taskStream, resultStream)
	go worker(taskStream, resultStream)
	go worker(taskStream, resultStream)

	position := 40
	fmt.Println("All fibonnaci numbers in the position", position)

	// add the buffer quantity to the taskStream
	for i := 0; i <= 40; i++ {
		taskStream <- i
	}
	// after taskStream has been fullfiled we need to close
	close(taskStream)

	// after fullfill the taskStream, iterate over results and print the items
	for i := 0; i <= 40; i++ {
		result := <-resultStream
		fmt.Println(result)
	}

}

func worker(taskStream chan int, resultStream chan int) {
	for task := range taskStream {
		resultStream <- fibonnaci(task)
	}
}
