package repository

import (
	"api/src/models"
	"database/sql"
	"fmt"
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

// Search all users based in the name or nick
func (repository users) Search(nameOrNick string) ([]models.User, error) {
	// %% = %
	// string
	// %% = %
	// final string = %nameOrNick%

	var searchStr = fmt.Sprintf("%%%s%%", nameOrNick)
	rows, err := repository.db.Query("SELECT id, name, nick, email, createdAt FROM users WHERE name LIKE ? or nick LIKE ?", searchStr, searchStr)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User

		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, err
}
