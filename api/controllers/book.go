package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"io"
	"strconv"

	"net/http"

	"github.com/gorilla/mux"
	"github.com/lazarok09/treinandosql/database"
	"github.com/lazarok09/treinandosql/helpers"
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
		return
	}
	defer connection.Close()
	var book BookBody
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("An error ocurred when reading the body database"))
		return
	}

	if bookMarshallError := json.Unmarshal(requestBody, &book); bookMarshallError != nil {
		w.Write([]byte("An error ocurred when parsing the message"))
		return
	}

	statment, err := connection.Prepare("INSERT INTO Book (name) VALUES (?)")

	if err != nil {
		w.Write([]byte("An error ocurred when creating the SQL statment"))
		return
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

	params := mux.Vars(r)

	paramName := "id"
	bookId, err := strconv.ParseUint(params[paramName], 10, 32)

	if err != nil {
		helpers.ThrowParamMissing(w, paramName)
		return
	}

	connection, err := database.Connect()
	if err != nil {
		helpers.ThrowDBConnectionError(w, err)
		return
	}
	defer connection.Close()

	var book Book
	queryRowError := connection.QueryRow("SELECT * FROM Book WHERE id = ?", bookId).Scan(&book.ID, &book.Name)
	if queryRowError != nil {
		if queryRowError == sql.ErrNoRows {
			helpers.ThrowEntityNotFounded(queryRowError.Error(), w, bookId)
			return
		}
		status := http.StatusInternalServerError
		w.WriteHeader(status)
		response := helpers.ResponseErrorShape{Message: "An error occurred when scanning book value to the struct", Error: queryRowError.Error(), Status: status}
		json.NewEncoder(w).Encode(response)
		return
	}

	json.NewEncoder(w).Encode(book)
}
func UpdateBook(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	paramName := "id"
	bookId, err := strconv.ParseUint(params[paramName], 10, 32)

	if err != nil {
		helpers.ThrowParamMissing(w, paramName)
		return
	}

	connection, err := database.Connect()
	if err != nil {
		helpers.ThrowDBConnectionError(w, err)
		return
	}
	defer connection.Close()

	var originalEntity Book
	queryRowError := connection.QueryRow("SELECT * FROM Book WHERE id = ?", bookId).Scan(&originalEntity.ID, &originalEntity.Name)
	if queryRowError != nil {
		if queryRowError == sql.ErrNoRows {
			helpers.ThrowEntityNotFounded(queryRowError.Error(), w, bookId)
			return
		}

		status := http.StatusInternalServerError
		w.WriteHeader(status)
		response := helpers.ResponseErrorShape{Message: "An error occurred when scanning book value to the struct", Error: queryRowError.Error(), Status: status}
		json.NewEncoder(w).Encode(response)
		return
	}
	var requestBodyBook Book

	json.NewDecoder(r.Body).Decode(&requestBodyBook)
	// INSERT INTO Book (name) VALUES (?)
	statment, err := connection.Prepare("UPDATE Book SET name = ? WHERE id = ?")
	if err != nil {
		helpers.ThrowAStatmentIssue(w, err.Error())

	}
	e, err := statment.Exec(requestBodyBook.Name, bookId)
	defer statment.Close()

	if err != nil {
		fmt.Println(e)
		helpers.ThrowAStatmentIssue(w, err.Error())
	}

	var responseBook Book

	connection.QueryRow("SELECT * FROM Book WHERE id = ?", bookId).Scan(&responseBook.ID, &responseBook.Name)
	if queryRowError != nil {
		helpers.ThrowEntityNotFounded(queryRowError.Error(), w, bookId)
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(responseBook)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	paramName := "id"
	bookId, err := strconv.ParseUint(params[paramName], 10, 32)

	if err != nil {
		helpers.ThrowParamMissing(w, paramName)
		return
	}

	connection, err := database.Connect()
	if err != nil {
		helpers.ThrowDBConnectionError(w, err)
		return
	}
	defer connection.Close()
	statment, err := connection.Prepare("DELETE FROM Book WHERE id = ?")
	if err != nil {
		helpers.ThrowAStatmentIssue(w, err.Error())
		return
	}
	defer statment.Close()

	if _, err := statment.Exec(bookId); err != nil {
		helpers.ThrowAStatmentIssue(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)

}
