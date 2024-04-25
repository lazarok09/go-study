package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lazarok09/treinandosql/controllers"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/book", controllers.CreateBook).Methods(http.MethodPost)
	router.HandleFunc("/book", controllers.GetBooks).Methods(http.MethodGet)
	router.HandleFunc("/book/{id}", controllers.GetBook).Methods(http.MethodGet)

	fmt.Println("Server listening at port 6000")

	log.Fatal(http.ListenAndServe(":6000", router))

}
