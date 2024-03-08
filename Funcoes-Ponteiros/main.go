package main

import "fmt"

// * pointer of string
// pointer receives new value.
func updatePointer(name *string, value string) {
	*name = value
}

func main() {
	userName := "Orlando"
	// & sends the pointer reference to the function
	updatePointer(&userName, "Patricia")

	fmt.Println(userName)

}
