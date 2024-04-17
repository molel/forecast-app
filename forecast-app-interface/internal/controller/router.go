package controller

import (
	"github.com/valyala/fasthttp"
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
	// auth
	case "/register":
		r.HandleRegister(ctx)
	case "/login":
		r.HandleLogin(ctx)

		// forecast

	default:
		ctx.SetStatusCode(fasthttp.StatusNotFound)
	}
}
