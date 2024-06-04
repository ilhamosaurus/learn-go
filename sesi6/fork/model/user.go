package model

import "errors"

var Users = []*User{}

type User struct {
	ID       uint64 `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	Username string `json:"username" gorm:"column:username;"`
	Email    string `json:"email" gorm:"column:email;"`
}

func (u *User) Validate() error {
	if u.Username == "" {
		return errors.New("username cannot be empty")
	}

	if u.Email == "" {
		return errors.New("email cannot be empty")
	}

	return nil
}