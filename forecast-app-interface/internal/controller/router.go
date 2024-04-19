package controller

import (
	"github.com/valyala/fasthttp"
	"log"
)

type UseCase interface {
	AuthUseCase
	ForecastUseCase
}

type AuthUseCase interface {
	Register(username, password string) error
	Login(username, password string) error
}

type ForecastUseCase interface {
	MakeForecast(name string, data []byte) ([]byte, error)
	GetForecast(name string) ([]byte, error)
}

type Router struct {
	useCase UseCase
}

func NewRouter(useCase UseCase) *Router {
	return &Router{
		useCase: useCase,
	}
}

func (r *Router) Handle(ctx *fasthttp.RequestCtx) {
	switch string(ctx.Path()) {
	// static
	case "/":
		log.Println(321)
		ctx.SetStatusCode(fasthttp.StatusOK)

	// auth
	case "/register":
		r.HandleRegister(ctx)
	case "/login":
		r.HandleLogin(ctx)

	// main
	case "/app":
		AuthMiddleware(r.HandleApp)(ctx)

	default:
		ctx.SetStatusCode(fasthttp.StatusNotFound)
	}
}

func (r *Router) HandleApp(ctx *fasthttp.RequestCtx) {
	n, err := ctx.WriteString("application response eqtqwrm,g;ernwmglhimer")
	if err != nil {
		log.Println(n, err)
	}
	ctx.SetStatusCode(fasthttp.StatusOK)
}
