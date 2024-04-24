package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/lazarok09/treinandosql/database"
)

type BookBody struct {
	Name string `json:"name"`
}
type BookResponse struct {
	Response interface{} `json:"response"`
}
type Book struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	connection, err := database.Connect()
	if err != nil {
		w.Write([]byte("An error ocurred when connecting the database"))

	}
	defer connection.Close()
	var book BookBody
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("An error ocurred when reading the body database"))

	}

	if bookMarshallError := json.Unmarshal(requestBody, &book); bookMarshallError != nil {
		w.Write([]byte("An error ocurred when parsing the message"))
	}

	statment, err := connection.Prepare("INSERT INTO Book (name) VALUES (?)")

	if err != nil {
		w.Write([]byte("An error ocurred when creating the SQL statment"))
	}
	defer statment.Close()

	resultOfInsertOperation, err := statment.Exec(book.Name)

	if err != nil {
		w.Write([]byte("An error ocurred when creating a book"))
	}

	bookIdResult, err := resultOfInsertOperation.LastInsertId()
	if err != nil {
		w.Write([]byte("An error ocurred when getting rows affected"))
	}

	if err != nil {
		w.Write([]byte("An error ocurred when getting the final message"))

	}
	response := BookResponse{Response: bookIdResult}

	finalResponse, err := json.Marshal(response)
	if err != nil {
		w.Write([]byte("An error ocurred when getting the final response"))

	}
	w.WriteHeader(http.StatusCreated)
	w.Write(finalResponse)
}
func GetBook(w http.ResponseWriter, r *http.Request) {
	connection, err := database.Connect()
	if err != nil {
		w.Write([]byte("An error ocurred when connecting the database"))

	}
	defer connection.Close()
	rows, err := connection.Query("SELECT * FROM Book")

	if err != nil {
		w.Write([]byte("An error ocurred when scaning book"))

	}
	defer rows.Close()

	var books []Book

	for rows.Next() {
		var book Book
		if err := rows.Scan(
			&book.ID,
			&book.Name,
		); err != nil {
			w.Write([]byte("An error ocurred when scaning books"))

		}
		books = append(books, book)
	}

	response := BookResponse{Response: books}

	finalResponse, err := json.Marshal(response)
	if err != nil {
		w.Write([]byte("An error ocurred when getting the final response"))

	}

	w.Write(finalResponse)
}
