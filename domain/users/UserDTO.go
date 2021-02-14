package users

import (
	"strings"

	"github.com/abmomin/users_api/utils/errors"
)

const (
	StatusActive = "active"
)
//User model
type User struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	DateCreated string `json:"dateCreated"`
	Status 		string `json:"status"`
	Password 	string `json:"password"`
}

//Validate validates an user.
func (user *User) Validate() *errors.RestErr {
	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestError("Invalid email address")
	}

	user.Password = strings.TrimSpace(user.Password)

	if user.Password == "" {
		return errors.NewBadRequestError("invalid password")
	}
	return nil
}
