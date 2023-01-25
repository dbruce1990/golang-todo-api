package models

import (
	"fmt"
)

type User struct {
	ID                   string
	Email                string `json:"email"`
	Password             string `json:"pwd"`
	PasswordConfirmation string `json:"pwd_confirm"`
}

func (u *User) ConfirmPasswordMatch() error {
	if u.Password != u.PasswordConfirmation {
		return fmt.Errorf("passwords do not match")
	}
	return nil
}
