package controller

import (
	"encoding/hex"
	"forecast-app-interface/utils"
	"github.com/valyala/fasthttp"
)

func (r *Router) HandleRegister(ctx *fasthttp.RequestCtx) {
	if string(ctx.Method()) != fasthttp.MethodPost {
		ctx.SetStatusCode(fasthttp.StatusMethodNotAllowed)
		return
	}

	username := ctx.QueryArgs().Peek("username")
	if len(username) == 0 {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	password := ctx.QueryArgs().Peek("password")
	if len(username) == 0 {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	hashedPassword := utils.HashValue(password)

	if err := r.useCase.Register(string(username), hex.EncodeToString(hashedPassword)); err != nil {
		// redirect
	}
}

func (r *Router) HandleLogin(ctx *fasthttp.RequestCtx) {
	if string(ctx.Method()) != fasthttp.MethodPost {
		ctx.SetStatusCode(fasthttp.StatusMethodNotAllowed)
		return
	}

	username := ctx.QueryArgs().Peek("username")
	if len(username) == 0 {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	password := ctx.QueryArgs().Peek("password")
	if len(username) == 0 {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	hashedPassword := utils.HashValue(password)

	if err := r.useCase.Login(string(username), hex.EncodeToString(hashedPassword)); err != nil {
		// redirect
	}
}
