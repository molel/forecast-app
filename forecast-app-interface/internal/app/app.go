package app

import (
	"fmt"
	"forecast-app-interface/config"
	"forecast-app-interface/internal/controller"
	"forecast-app-interface/internal/usecase"
	"github.com/valyala/fasthttp"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func Run(cfg *config.Config) {
	useCase := usecase.NewUseCase()
	if err := useCase.Init(cfg.AuthServiceAddress, cfg.PredictServiceAddress); err != nil {
		log.Fatalf("Cannot init use case: %s\n", err)
	}

	router := controller.NewRouter(useCase)

	go func() {
		log.Fatal(fasthttp.ListenAndServe(fmt.Sprintf(":%d", cfg.HTTP), router.Handle))
	}()

	terminationChan := make(chan os.Signal, 1)
	signal.Notify(terminationChan, syscall.SIGINT, syscall.SIGTERM)
	<-terminationChan
}
