package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	dataSourceName := "lazarok09:lazarok09@/devbook"
	db, error := sql.Open("mysql", dataSourceName)

	if error != nil {
		fmt.Println("Erro na abertura")

		log.Fatal(error)
	}
	defer db.Close()

	if error = db.Ping(); error != nil {
		fmt.Println("Erro no ping")

		log.Fatal(error)
	}
	fmt.Println("open conection")
}
