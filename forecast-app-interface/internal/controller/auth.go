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

	token, expiration, err := createToken(username)
	if err != nil {
		ctx.Error("cannot create token", fasthttp.StatusInternalServerError)
		return
	}

	cookie := fasthttp.AcquireCookie()
	defer fasthttp.ReleaseCookie(cookie)

	cookie.Reset()

	cookie.SetKey("token")
	cookie.SetValue(token)
	cookie.SetKey("username")
	cookie.SetValue(username)
	cookie.SetExpire(expiration)

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

	token, expiration, err := createToken(username)
	if err != nil {
		ctx.Error("cannot create token", fasthttp.StatusInternalServerError)
		return
	}

	cookie := fasthttp.AcquireCookie()
	defer fasthttp.ReleaseCookie(cookie)

	cookie.Reset()
	cookie.SetKey("token")
	cookie.SetValue(token)
	cookie.SetExpire(expiration)
	ctx.Response.Header.SetCookie(cookie)

	cookie.Reset()
	cookie.SetKey("username")
	cookie.SetValue(username)
	cookie.SetExpire(expiration)
	ctx.Response.Header.SetCookie(cookie)

	ctx.Redirect("/app", fasthttp.StatusSeeOther)
}
