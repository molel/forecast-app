package usecase

import (
	"fmt"
	"log"

	"forecast-app-interface/internal/controller/gen/go/auth"
	"forecast-app-interface/internal/controller/gen/go/predict"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type UseCase struct {
	authConn   *grpc.ClientConn
	authClient auth.AuthServiceClient

	predictConn   *grpc.ClientConn
	predictClient predict.PredictServiceClient
}

func NewUseCase() *UseCase {
	return &UseCase{}
}

func (u *UseCase) Init(authServiceAddress, predictServiceAddress string) error {
	var err error

	u.authConn, err = grpc.Dial(fmt.Sprintf("dns:///%s", authServiceAddress), grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		return err
	}
	log.Println("Successfully connected to auth service")

	u.authClient = auth.NewAuthServiceClient(u.authConn)

	u.predictConn, err = grpc.Dial(fmt.Sprintf("dns:///%s", predictServiceAddress), grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		return err
	}
	log.Println("Successfully connected to predict service")

	u.predictClient = predict.NewPredictServiceClient(u.predictConn)

	return err
}
