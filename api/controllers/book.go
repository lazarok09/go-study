package controllers

import (
	"encoding/json"
	"fmt"
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
type ResponseErrorShape struct {
	Message string
	Error   string
	Status  int
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

	response := BookResponse{Response: bookIdResult}

	finalResponse, err := json.Marshal(response)
	if err != nil {
		w.Write([]byte("An error ocurred when getting the final response"))

	}
	w.WriteHeader(http.StatusCreated)
	w.Write(finalResponse)
}
func GetBooks(w http.ResponseWriter, r *http.Request) {
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
		err := rows.Scan(
			&book.ID,
			&book.Name,
		)

		if err != nil {
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
func GetBook(w http.ResponseWriter, r *http.Request) {
	connection, err := database.Connect()
	if err != nil {
		ThrowDBConnectionError(w, err)
		return
	}
	defer connection.Close()

	paramName := "id"
	bookId := "12"

	if len(bookId) < 1 {
		ThrowParamMissing(w, paramName)
		return
	}

	var book Book
	queryRowError := connection.QueryRow("SELECT * FROM Book WHERE id = ?", bookId).Scan(&book.ID, &book.Name)
	if queryRowError != nil {
		status := http.StatusInternalServerError
		w.WriteHeader(status)
		response := ResponseErrorShape{Message: "An error occurred when scanning book value to the struct", Error: queryRowError.Error(), Status: status}
		json.NewEncoder(w).Encode(response)
		return
	}

	json.NewEncoder(w).Encode(book)
}

func ThrowParamMissing(w http.ResponseWriter, param string) {
	status := http.StatusUnprocessableEntity
	message := fmt.Sprintf("Missing the %s param", param)

	responseShape := ResponseErrorShape{Message: message, Error: "", Status: status}

	responseJSON, _ := json.Marshal(responseShape)

	w.WriteHeader(status)
	w.Write(responseJSON)
	panic(message)
}
func ThrowDBConnectionError(w http.ResponseWriter, err error) {
	status := http.StatusInternalServerError

	responseShape := ResponseErrorShape{Message: "Seems like we're facing issues connecting the database.", Error: err.Error(), Status: status}

	responseJSON, _ := json.Marshal(responseShape)
	if err != nil {
		fmt.Println(err.Error())
	}

	w.WriteHeader(status)
	w.Write(responseJSON)
	panic(err)
}

// a function that receives name as being what you wanted to database prepare
func ThrowAStatmentIssue(name string) {
	message := fmt.Sprintf("An issue ocurred when you tried to: %s ", name)
	panic(message)
}
