package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.LoadEnvs()

	var loadedAt = fmt.Sprintf("Server loaded at %d", config.Port)
	fmt.Println(loadedAt)

	r := router.Generate()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
