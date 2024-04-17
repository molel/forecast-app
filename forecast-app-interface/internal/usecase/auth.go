package usecase

import (
	"context"
	"fmt"
	"forecast-app-interface/internal/controller/gen/go/auth"
	"sync"
)

const (
	registerErrorTemplate = "cannot register user [username=%s]: %s"
	loginErrorTemplate    = "cannot login user [username=%s]: %s"
)

var registerRequestsPool = sync.Pool{
	New: func() any {
		return &auth.RegisterRequest{}
	},
}

var loginRequestsPool = sync.Pool{
	New: func() any {
		return &auth.LoginRequest{}
	},
}

func (u *UseCase) Register(username, password string) error {
	request := registerRequestsPool.Get().(*auth.RegisterRequest)
	request.Username = username
	request.Password = password

	response, err := u.authClient.Register(context.Background(), request)
	if err != nil {
		return fmt.Errorf(registerErrorTemplate, username, err)
	}

	if !response.Success {
		return fmt.Errorf(registerErrorTemplate, username, response.Text)
	}

	return nil
}

func (u *UseCase) Login(username, password string) error {
	request := loginRequestsPool.Get().(*auth.LoginRequest)
	request.Username = username
	request.Password = password

	response, err := u.authClient.Login(context.Background(), request)
	if err != nil {
		return fmt.Errorf(loginErrorTemplate, username, err)
	}

	if !response.Success {
		return fmt.Errorf(loginErrorTemplate, username, response.Text)
	}

	return nil
}
