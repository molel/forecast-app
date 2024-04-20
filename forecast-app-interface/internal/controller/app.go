package controller

import (
	"encoding/json"
	"forecast-app-interface/utils"
	"html/template"
	"slices"

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
	Username string
	Name     string
	Unit     string `json:"unit"`
	Items    []item `json:"items"`
}

type item struct {
	Ts    int64   `json:"ts"`
	Value float64 `json:"value"`
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

	name := string(ctx.FormValue("name"))
	username := string(ctx.Request.Header.Cookie("username"))

	templateData := appTemplateData{
		Username: username,
		Name:     name,
	}

	if len(name) > 0 {
		data, err := r.useCase.GetForecast(username, name)
		if err != nil {
			templateData.Name = "Не удалось получить данные"
			ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		} else if err = json.Unmarshal(data, &templateData); err != nil {
			templateData.Name = "Не удалось получить данные"
			ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		}

		slices.SortStableFunc(templateData.Items, func(a, b item) int {
			return int(a.Ts - b.Ts)
		})
	}

	if err := templates.ExecuteTemplate(ctx, "app.gohtml", templateData); err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
	}

	ctx.Response.Header.SetContentType("text/html")
	ctx.SetStatusCode(fasthttp.StatusOK)
}
