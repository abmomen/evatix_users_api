package services

import (
	"github.com/abmomin/users_api/domain/users"
	"github.com/abmomin/users_api/utils/date"
	"github.com/abmomin/users_api/utils/errors"
)

//CreateUser a service to create a new user in database.
func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	user.Status = users.StatusActive
	user.DateCreated = date.GetNowDBFormat()
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

//GetUser returns a user with provided user id.
func GetUser(userID int64) (*users.User, *errors.RestErr) {
	result := &users.User{ID: userID}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestErr) {
	current, err := GetUser(user.ID)
	if  err != nil {
		return nil, err
	}

	if isPartial {
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}
		if user.LastName != "" {
			current.LastName = user.LastName
		}
		if user.Email != "" {
			current.Email = user.Email
		}
	} else {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
	}

	if err := current.Update(); err != nil {
		return nil, err
	}
	return current, nil
}

func DeleteUser(userId int64) *errors.RestErr  {
	user := &users.User{ID: userId}
	return user.Delete()
}

func Search(status string) ([]users.User, *errors.RestErr) {
	dao := users.User{}
	return dao.FindByStatus(status)
}
