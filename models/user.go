package models

import (
	"crud-api-golang/internal/constants"
	"errors"
	"strings"
	"time"
)

// User is a user entity
type User struct {
	ID       uint64    `json:"id,omitempty"`
	Name     string    `json:"name,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Password string    `json:"password,omitempty"`
	CreateAt time.Time `json:"create_at,omitempty"`
}

// PrepareUser prepare the user model to insert into the database
func (u *User) PrepareUser() error {

	err := u.checkFields()
	if err != nil {
		return err
	}

	u.formatFields()

	return nil
}

// checkFields checks fields that cannot be null
func (u *User) checkFields() error {
	if u.Name == "" {
		return errors.New(constants.ERROR_EMPTY_NAME)
	}
	if u.Nick == "" {
		return errors.New(constants.ERROR_EMPTY_NICK)
	}
	if u.Email == "" {
		return errors.New(constants.ERROR_EMPTY_EMAIL)
	}
	if u.Password == "" {
		return errors.New(constants.ERROR_EMPTY_PASSWORD)
	}

	return nil
}

// formatFields formats filled fields by removing spaces from the ends
func (u *User) formatFields() {
	u.Name = strings.TrimSpace(u.Name)
	u.Nick = strings.TrimSpace(u.Nick)
	u.Email = strings.TrimSpace(u.Email)
}
