package main

import (
	"lazarok09/estudandohttp/controllers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/loja", controllers.Store)

	log.Fatal(http.ListenAndServe(":5001", nil))
}
