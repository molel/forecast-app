package controller

import (
	"github.com/valyala/fasthttp"
	"html/template"
	"math"
	"time"
)

var (
	templates = template.Must(template.ParseGlob("./web/templates/*.html"))
)

type app struct {
	Username string
	Items    []item
}

type item struct {
	Time  time.Time
	Value float64
}

func (r *Router) HandleRoot(ctx *fasthttp.RequestCtx) {
	if string(ctx.Method()) != fasthttp.MethodGet {
		ctx.SetStatusCode(fasthttp.StatusMethodNotAllowed)
		return
	}

	ctx.SendFile("./web/static/root.html")
	ctx.SetStatusCode(fasthttp.StatusOK)
}

func (r *Router) HandleApp(ctx *fasthttp.RequestCtx) {
	if string(ctx.Method()) != fasthttp.MethodGet {
		ctx.SetStatusCode(fasthttp.StatusMethodNotAllowed)
		return
	}

	data := app{Username: string(ctx.Request.Header.Cookie("username"))}
	data.Items = make([]item, 365)
	now := time.Now().UnixNano()
	for i := range data.Items {
		data.Items[i].Time = time.Unix(0, now+int64(time.Duration(i)*time.Hour*24))
		data.Items[i].Value = math.Sin(float64(i))
	}

	if err := templates.ExecuteTemplate(ctx, "app.html", data); err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
	}

	ctx.Response.Header.SetContentType("text/html")
	ctx.SetStatusCode(fasthttp.StatusOK)
}
