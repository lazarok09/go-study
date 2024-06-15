package models

import (
	"errors"
	"strings"
	"time"
)

type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"CreatedAt,omitempty"`
}

// Validate user for empty values and trim the nick name and email of the user
func (user *User) Prepare() error {
	if err := user.validate(); err != nil {
		return err
	}
	user.format()
	return nil
}

func (user *User) validate() error {
	var NoEmptyValues = "no empt values"

	if len(user.Name) == 0 {
		return errors.New(NoEmptyValues)
	}
	if len(user.Nick) == 0 {
		return errors.New(NoEmptyValues)
	}
	if len(user.Email) == 0 {
		return errors.New(NoEmptyValues)
	}
	if len(user.Password) == 0 {
		return errors.New(NoEmptyValues)
	}

	return nil
}

func (user *User) format() {
	user.Name = strings.TrimSpace(user.Name)
	user.Email = strings.TrimSpace(user.Email)
	user.Nick = strings.TrimSpace(user.Nick)
}
