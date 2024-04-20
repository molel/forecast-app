package controller

import (
	"github.com/valyala/fasthttp"
	"html/template"
	"math"
	"time"
)

var (
	templates = template.Must(template.ParseGlob("./web/templates/*.gohtml"))
)

type app struct {
	Username   string
	TimeSeries []timeSeries
}

type timeSeries struct {
	Name  string
	Unit  string
	Items []timeSeriesItem
}

type timeSeriesItem struct {
	Time  time.Time
	Value float64
}

func (r *Router) HandleRoot(ctx *fasthttp.RequestCtx) {
	if string(ctx.Method()) != fasthttp.MethodGet {
		ctx.SetStatusCode(fasthttp.StatusMethodNotAllowed)
		return
	}

	ctx.SendFile("./web/static/pages/root.html")
	ctx.SetStatusCode(fasthttp.StatusOK)
}

func (r *Router) HandleApp(ctx *fasthttp.RequestCtx) {
	if string(ctx.Method()) != fasthttp.MethodGet {
		ctx.SetStatusCode(fasthttp.StatusMethodNotAllowed)
		return
	}

	data := app{Username: string(ctx.Request.Header.Cookie("username"))}
	data.TimeSeries = []timeSeries{
		{
			Name:  "Time series name",
			Unit:  "smth",
			Items: make([]timeSeriesItem, 365),
		},
	}
	now := time.Now().UnixNano()
	for i := range data.TimeSeries[0].Items {
		data.TimeSeries[0].Items[i].Time = time.Unix(0, now+int64(time.Duration(i)*time.Hour*24))
		data.TimeSeries[0].Items[i].Value = math.Sin(float64(i))
	}

	if err := templates.ExecuteTemplate(ctx, "app.gohtml", data); err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
	}

	ctx.Response.Header.SetContentType("text/html")
	ctx.SetStatusCode(fasthttp.StatusOK)
}
