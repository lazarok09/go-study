package repository

import (
	"api/src/models"
	"database/sql"
)

// Represents the repository of users
type users struct {
	db *sql.DB
}

// Create a new user repository with the database.
func NewUsersRepository(db *sql.DB) *users {
	return &users{db}
}

// Create a user in the database and return his id if created
func (repository users) Create(user models.User) (uint64, error) {
	statment, err := repository.db.Prepare("INSERT INTO users (name, nick, email, password) values(?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer statment.Close()

	result, err := statment.Exec(user.Name, user.Nick, user.Email, user.Password)

	if err != nil {
		return 0, err
	}

	lastInsertedID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return uint64(lastInsertedID), nil
}
