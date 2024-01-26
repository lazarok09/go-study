package auxiliar

import (
	"fmt"
)

func SortNumbers(numbers [10]int) [10]int {

	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}

	return numbers
}
