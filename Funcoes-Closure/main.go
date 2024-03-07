package main

import "fmt"

// closure functions in the body may have a variable that will be storad
// this can be helpfull when we want to count how many times we sended a ticket
// or a error message in a execution.
func errorCounter() func() int {
	errors := 0

	return func() int {
		errors += 1
		return errors
	}
}

func main() {
	fmt.Println("hlelo")
	addError := errorCounter()
	// addError() will utilize the closure returned function
	// as we can see, that function had a errors initiated in the beginning (line 19) as zero
	// but through the time, we used to add more and more.
	// this will change the value of errors.

	fmt.Println(addError()) // 1
	fmt.Println(addError()) // 2
	fmt.Println(addError()) // 3
	fmt.Println(addError()) // 4

}
