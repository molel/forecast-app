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
	MakeForecast(username, name, unit string, period, predictPeriods int32, tss []int64, values []float64) error
	GetForecast(username, name string) (string, int64, any, error)
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
		r.HandleRoot(ctx)
	case "/favicon.ico":
		r.HandleFavicon(ctx)

	// auth
	case "/register":
		r.HandleRegister(ctx)
	case "/login":
		r.HandleLogin(ctx)

	// main
	case "/app":
		AuthMiddleware(r.HandleApp)(ctx)
	case "/app/get":
		AuthMiddleware(r.HandleGetPredict)(ctx)
	case "/app/predict":
		AuthMiddleware(r.HandleMakePredict)(ctx)

	default:
		ctx.Redirect("/app", fasthttp.StatusSeeOther)
	}
}
