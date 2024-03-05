package UserModel

import "fmt"

func InsertData(disconnectDB func()) {
	defer disconnectDB()
	fmt.Println("Inserted 200 rows at Table User")
}
func DeleteData(disconnectDB func()) {
	defer disconnectDB()
	fmt.Println("Deleted 150 data at able User")
}
