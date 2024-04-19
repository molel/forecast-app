package usecase

import (
	"errors"
	"fmt"
)

var (
	userCheckErrorTemplate   = "cannot check user existence: %w"
	userCreateErrorTemplate  = "cannot create user: %w"
	getPasswordErrorTemplate = "cannot get password: %w"

	userAlreadyExistsError = errors.New("user already exists")
	userDoesNotExistError  = errors.New("user does not exist")
	wrongPasswordError     = errors.New("incorrect password")
)

func (u *UseCase) Register(username, password string) error {
	if exists, err := u.repo.CheckUser(username); err != nil {
		return fmt.Errorf(userCheckErrorTemplate, err)
	} else if exists {
		return userAlreadyExistsError
	}

	err := u.repo.CreateUser(username, password)
	if err != nil {
		return fmt.Errorf(userCreateErrorTemplate, err)
	}

	return nil
}

func (u *UseCase) Login(username, password string) error {
	if exists, err := u.repo.CheckUser(username); err != nil {
		return fmt.Errorf(userCheckErrorTemplate, err)
	} else if !exists {
		return userDoesNotExistError
	}

	storedPassword, err := u.repo.GetUserPassword(username)
	if err != nil {
		return fmt.Errorf(getPasswordErrorTemplate, err)
	}
	if password != storedPassword {
		return wrongPasswordError
	}

	return nil
}
