package usecase

import (
	"forecast-app-interface/internal/controller/gen/go/auth"
	"google.golang.org/grpc"
)

type UseCase struct {
	authConn   *grpc.ClientConn
	authClient auth.AuthServiceClient
}

func NewUseCase() *UseCase {
	return &UseCase{}
}

func (u *UseCase) Init(authServiceAddress string) error {
	var err error

	u.authConn, err = grpc.Dial(authServiceAddress)
	if err != nil {
		return err
	}

	u.authClient = auth.NewAuthServiceClient(u.authConn)

	return err
}
