package services

import (
	"github.com/pradyumna30/bookstore_users-api/domain/users"
	"github.com/pradyumna30/bookstore_users-api/utils/errors"
)

//Entire business logic will be here

func GetUser(id int64) (*users.User, *errors.RestErr) {
	if id <= 0 {
		return nil, errors.NewBadRequestError("invalid user id")
	}
	var user users.User
	user.Id = id
	if err := user.Get(); err != nil {
		return nil, err
	}
	return &user, nil
}

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

func FindUser() {

}
