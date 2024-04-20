package usecase

import (
	"forecast-app-interface/internal/controller/gen/go/predict"
	"log"

	"forecast-app-interface/internal/controller/gen/go/auth"
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

	u.authConn, err = grpc.Dial(authServiceAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	log.Println("Successfully connected to auth service")

	u.authClient = auth.NewAuthServiceClient(u.authConn)

	u.predictConn, err = grpc.Dial(predictServiceAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	log.Println("Successfully connected to predict service")

	u.predictClient = predict.NewPredictServiceClient(u.predictConn)

	return err
}
