package main

import (
	"fmt"
	UserModel "lazarok09/funcoes-defer/User"
)

func connectDB() {
	fmt.Println("Connect to DB")
}

func disconnectDB() {
	fmt.Println("Desconected DB")
}
func main() {

	connectDB()

	UserModel.InsertData(disconnectDB)

	UserModel.DeleteData(disconnectDB)

}
