package models

import (
	"errors"
	"strings"
	"time"
)

// Model User represents a Users using the social network
type User struct {
	ID uint64 `json:"id,omitempty"`
	Nome string `json:"nome,omitempty"`
	Nick string `json:"nick,omitempty"`
	Pass string `json:"pass,omitempty"`
	Email string `json:"email,omitempty"`
	CreatedAt time.Time `json:"CreatedAt,omitempty"`
}

//prepare call the methods and formate the fields to insert in db
func (user *User) Prepare(step string) error {
	if erro := user.validate(step); erro != nil {
		return erro
	}
 
	user.formate()
	return nil
}

func (user *User) validate(step string) error {
	if user.Nome == "" {
		return errors.New("The name field is required and cannot be left blank")
	}
	if user.Nick == "" {
		return errors.New("The nick field is required and cannot be left blank")
	}
	if user.Email == "" {
		return errors.New("The email field is required and cannot be left blank")
	}

	if step == "register" && user.Pass == "" {
		return errors.New("The pass field is required and cannot be left blank")
	}
	return nil 
}


func ( user *User) formate() {
	user.Nome = strings.TrimSpace(user.Nome)
	user.Email = strings.TrimSpace(user.Email)
	user.Nick = strings.TrimSpace(user.Nick)
}