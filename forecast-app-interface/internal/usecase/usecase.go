package usecase

import (
	"log"

	"forecast-app-interface/internal/controller/gen/go/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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

	u.authConn, err = grpc.Dial(authServiceAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	log.Println("Successfully connected to auth service")

	u.authClient = auth.NewAuthServiceClient(u.authConn)

	return err
}
