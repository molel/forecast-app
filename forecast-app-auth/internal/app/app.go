package app

import (
	"fmt"
	"forecast-app-auth/config"
	"forecast-app-auth/internal/controller/gen/go/auth"
	"forecast-app-auth/internal/repo"
	"forecast-app-auth/internal/usecase"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func Run(cfg *config.Config) {
	repository := repo.NewRepository()
	if err := repository.Init(cfg.DatabaseUser, cfg.DatabasePassword, cfg.DatabaseAddress, cfg.DatabaseName); err != nil {
		log.Fatalf("Cannot init repository: %s\n", err)
	}

	useCase := usecase.NewUseCase(repository)

	server := auth.NewServer(useCase)
	grpcServer := grpc.NewServer()
	auth.RegisterAuthServiceServer(grpcServer, server)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.HTTP))
	if err != nil {
		log.Fatalf("Cannot init tcp listener: %s\n", err)
	}

	if err = grpcServer.Serve(listener); err != nil {
		log.Fatalf("Cannot serve gRPC server: %s\n", err)
	}

	terminationChan := make(chan os.Signal, 1)
	signal.Notify(terminationChan, syscall.SIGINT, syscall.SIGTERM)
	<-terminationChan
}
