package controller

import (
	"encoding/hex"
	"forecast-app-interface/utils"
	"github.com/valyala/fasthttp"
	"google.golang.org/grpc/status"
)

func (r *Router) HandleRegister(ctx *fasthttp.RequestCtx) {
	if string(ctx.Method()) != fasthttp.MethodPost {
		ctx.SetStatusCode(fasthttp.StatusMethodNotAllowed)
		return
	}

	username := string(ctx.FormValue("username"))
	hashedPassword := utils.HashValue(ctx.FormValue("password"))

	if err := r.useCase.Register(username, hex.EncodeToString(hashedPassword)); err != nil {
		ctx.Error(status.Convert(err).Message(), fasthttp.StatusInternalServerError)
		return
	}

	token, err := createToken(username)
	if err != nil {
		ctx.Error("cannot create token", fasthttp.StatusInternalServerError)
		return
	}

	cookie := &fasthttp.Cookie{}
	cookie.SetKey("token")
	cookie.SetValue(token)
	cookie.SetKey("username")
	cookie.SetValue(username)
	ctx.Response.Header.SetCookie(cookie)

	ctx.Redirect("/app", fasthttp.StatusSeeOther)
}

func (r *Router) HandleLogin(ctx *fasthttp.RequestCtx) {
	if string(ctx.Method()) != fasthttp.MethodPost {
		ctx.SetStatusCode(fasthttp.StatusMethodNotAllowed)
		return
	}

	username := string(ctx.FormValue("username"))
	hashedPassword := utils.HashValue(ctx.FormValue("password"))

	if err := r.useCase.Login(username, hex.EncodeToString(hashedPassword)); err != nil {
		ctx.Error(status.Convert(err).Message(), fasthttp.StatusUnauthorized)
		return
	}

	token, err := createToken(username)
	if err != nil {
		ctx.Error("cannot create token", fasthttp.StatusInternalServerError)
		return
	}

	cookie := &fasthttp.Cookie{}
	cookie.SetKey("token")
	cookie.SetValue(token)
	cookie.SetKey("username")
	cookie.SetValue(username)
	ctx.Response.Header.SetCookie(cookie)

	ctx.Redirect("/app", fasthttp.StatusSeeOther)
}
