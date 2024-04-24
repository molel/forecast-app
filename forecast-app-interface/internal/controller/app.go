package controller

import (
	"encoding/csv"
	"fmt"
	"forecast-app-interface/internal/controller/gen/go/predict"
	"forecast-app-interface/utils"
	"github.com/valyala/fasthttp"
	"html/template"
	"os"
	"strconv"
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
	Error     string
	Items     []*predict.TimeSeriesItem
}

func (r *Router) HandleRoot(ctx *fasthttp.RequestCtx) {
	if string(ctx.Method()) != fasthttp.MethodGet {
		ctx.SetStatusCode(fasthttp.StatusMethodNotAllowed)
		return
	}

	ctx.Response.Header.SetContentType("text/html")

	ctx.SendFile("./web/static/pages/root.html")
	ctx.SetStatusCode(fasthttp.StatusOK)
}

func (r *Router) HandleFavicon(ctx *fasthttp.RequestCtx) {
	if string(ctx.Method()) != fasthttp.MethodGet {
		ctx.SetStatusCode(fasthttp.StatusMethodNotAllowed)
		return
	}

	data, err := os.ReadFile("./web/static/favicon/favicon.ico")
	if err != nil {
		ctx.Error(fmt.Sprintf("cannot read favicon: %s", err), fasthttp.StatusInternalServerError)
		return
	}

	ctx.Response.Header.SetContentType("image/x-icon")

	ctx.SetBody(data)
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
		return
	}

	ctx.Response.Header.SetContentType("text/html")
	ctx.SetStatusCode(fasthttp.StatusOK)
}

func (r *Router) HandleGetPredict(ctx *fasthttp.RequestCtx) {
	if string(ctx.Method()) != fasthttp.MethodGet {
		ctx.SetStatusCode(fasthttp.StatusMethodNotAllowed)
		return
	}

	username := string(ctx.Request.Header.Cookie("username"))

	templateData := appTemplateData{
		Username: username,
	}

	name := string(ctx.FormValue("name"))
	if len(name) > 0 {
		unit, predictStart, items, err := r.useCase.GetForecast(username, name)
		if err != nil {
			templateData.Error = err.Error()
			ProcessTemplate(ctx, fasthttp.StatusInternalServerError, "app.gohtml", templateData)
			return
		} else {
			templateData.Name = name
			templateData.Unit = unit
			templateData.Delimiter = predictStart
			templateData.Items = items.([]*predict.TimeSeriesItem)
		}
	}

	ProcessTemplate(ctx, fasthttp.StatusOK, "app.gohtml", templateData)
	return
}

func (r *Router) HandleMakePredict(ctx *fasthttp.RequestCtx) {
	if string(ctx.Method()) != fasthttp.MethodPost {
		ctx.SetStatusCode(fasthttp.StatusMethodNotAllowed)
		return
	}

	var (
		err          error
		templateData appTemplateData
	)

	templateData.Username = string(ctx.Request.Header.Cookie("username"))

	fileHeader, err := ctx.FormFile("time_series_file")
	if err != nil {
		templateData.Error = err.Error()
		ProcessTemplate(ctx, fasthttp.StatusInternalServerError, "app.gohtml", templateData)
		return
	}

	file, err := fileHeader.Open()
	if err != nil {
		templateData.Error = err.Error()
		ProcessTemplate(ctx, fasthttp.StatusInternalServerError, "app.gohtml", templateData)
		return
	}

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		templateData.Error = err.Error()
		ProcessTemplate(ctx, fasthttp.StatusInternalServerError, "app.gohtml", templateData)
		return
	}

	username := string(ctx.Request.Header.Cookie("username"))
	period, _ := strconv.ParseInt(string(ctx.FormValue("period")), 10, 32)
	unit := string(ctx.FormValue("unit"))
	name := fileHeader.Filename
	tss := make([]int64, len(records))
	values := make([]float64, len(records))

	for i := range records {
		tss[i], err = strconv.ParseInt(records[i][0], 10, 64)
		if err != nil {
			templateData.Error = err.Error()
			ProcessTemplate(ctx, fasthttp.StatusInternalServerError, "app.gohtml", templateData)
			return
		}

		values[i], err = strconv.ParseFloat(records[i][1], 64)
		if err != nil {
			templateData.Error = err.Error()
			ProcessTemplate(ctx, fasthttp.StatusInternalServerError, "app.gohtml", templateData)
			return
		}
	}

	if err = r.useCase.MakeForecast(username, name, unit, int32(period), tss, values); err != nil {
		templateData.Error = err.Error()
		ProcessTemplate(ctx, fasthttp.StatusInternalServerError, "app.gohtml", templateData)
		return
	}

	ProcessTemplate(ctx, fasthttp.StatusOK, "app.gohtml", templateData)
	return
}

func ProcessTemplate(ctx *fasthttp.RequestCtx, statusCode int, templateName string, templateData any) {
	if err := templates.ExecuteTemplate(ctx, templateName, templateData); err != nil {
		ctx.Error("cannot process template: "+err.Error(), fasthttp.StatusInternalServerError)
		return
	}

	ctx.Response.Header.SetContentType("text/html")
	ctx.SetStatusCode(statusCode)
}
