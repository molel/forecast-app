package controller

import (
	"html/template"

	"forecast-app-interface/internal/controller/gen/go/predict"
	"forecast-app-interface/utils"
	"github.com/valyala/fasthttp"
)

var (
	templates = template.
		Must(
			template.
				New("").
				Funcs(template.FuncMap{"formatTs": utils.FormatTs}).
				ParseGlob("./web/templates/*.gohtml"),
		)
)

type appTemplateData struct {
	Username  string
	Name      string
	Unit      string
	Delimiter int64
	Items     []*predict.TimeSeriesItem
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

	username := string(ctx.Request.Header.Cookie("username"))

	templateData := appTemplateData{
		Username: username,
	}

	if err := templates.ExecuteTemplate(ctx, "app.gohtml", templateData); err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
	}

	ctx.Response.Header.SetContentType("text/html")
	ctx.SetStatusCode(fasthttp.StatusOK)
}

func (r *Router) HandleGetPredict(ctx *fasthttp.RequestCtx) {
	if string(ctx.Method()) != fasthttp.MethodGet {
		ctx.SetStatusCode(fasthttp.StatusMethodNotAllowed)
		return
	}

	name := string(ctx.FormValue("name"))
	username := string(ctx.Request.Header.Cookie("username"))

	templateData := appTemplateData{
		Username: username,
		Name:     name,
	}

	if len(name) > 0 {
		unit, predictStart, items, err := r.useCase.GetForecast(username, name)
		if err != nil {
			templateData.Name = "Не удалось получить данные"
			ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		} else {
			templateData.Unit = unit
			templateData.Delimiter = predictStart
			templateData.Items = items.([]*predict.TimeSeriesItem)
		}
	}

	if err := templates.ExecuteTemplate(ctx, "app.gohtml", templateData); err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
	}

	ctx.Response.Header.SetContentType("text/html")
	ctx.SetStatusCode(fasthttp.StatusOK)
}
