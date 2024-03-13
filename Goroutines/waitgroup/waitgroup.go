package waitgroup

import (
	"fmt"
	"io"
	"log"

	"net/http"
	"sync"
	"time"
)

const URL = "https://pokeapi.co/api/v2/pokemon/ditto"

func doMath(waitGroup *sync.WaitGroup) int {

	// to run the defer right before the return statement of result
	defer waitGroup.Done()

	// do some random math

	result := 2 + 2
	// sleep to goroutine check
	time.Sleep(time.Second * 3)

	fmt.Println(result)
	return result
}

func fetchPokemons(waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()

	res, err := http.Get(URL)

	if err != nil {
		log.Fatal("Error in fetch api at this url" + URL)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		log.Fatal("Error in fetch api at the reading of body")
	}
	var bodyText string
	for _, text := range body {

		bodyText += string(text)

	}
	fmt.Println(bodyText)

}
func WaitGroup() {
	var waitGroup sync.WaitGroup

	// create a go routine that call fetch

	waitGroup.Add(5)

	go doMath(&waitGroup)
	go doMath(&waitGroup)
	go doMath(&waitGroup)
	go doMath(&waitGroup)
	go fetchPokemons(&waitGroup)

	waitGroup.Wait()

}
