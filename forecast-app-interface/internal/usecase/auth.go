package usecase

import (
	"context"
	"fmt"
	"google.golang.org/grpc/status"
	"sync"

	"forecast-app-interface/internal/controller/gen/go/auth"
)

const (
	registerErrorTemplate = "cannot register user: %s"
	loginErrorTemplate    = "cannot login user: %s"
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

	_, err := u.authClient.Register(context.Background(), request)
	if err != nil {
		err = fmt.Errorf(registerErrorTemplate, status.Convert(err).Message())
	}

	registerRequestsPool.Put(request)

	return err
}

func (u *UseCase) Login(username, password string) error {
	request := loginRequestsPool.Get().(*auth.LoginRequest)
	request.Username = username
	request.Password = password

	_, err := u.authClient.Login(context.Background(), request)
	if err != nil {
		err = fmt.Errorf(loginErrorTemplate, status.Convert(err).Message())
	}

	loginRequestsPool.Put(request)

	return err
}
