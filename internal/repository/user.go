package repository

import (
	"crud-api-golang/models"
	"database/sql"
	"fmt"
)

// user represents a user repository
type user struct {
	db *sql.DB
}

// NewUserRepository create a user repository
func NewUserRepository(db *sql.DB) *user {
	return &user{db}
}

// Create insert a user into the database
func (u user) Create(user *models.User) (uint64, error) {

	statement, err := u.db.Prepare(
		"INSERT INTO users (name, nick, email, password) VALUES (?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(user.Name, user.Nick, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	userCreatedID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(userCreatedID), nil
}

func (u user) GetList(keyword string) ([]models.User, error) {

	nameOrNick := fmt.Sprintf("%%%s%%", keyword)

	usersReturned, err := u.db.Query(
		"SELECT id, name, nick, email, createAt FROM users WHERE name LIKE ? OR nick LIKE ?", nameOrNick, nameOrNick)
	if err != nil {
		return nil, err
	}

	var users []models.User

	for usersReturned.Next() {
		var user models.User

		if err = usersReturned.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreateAt,
		); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
