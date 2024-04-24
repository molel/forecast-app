package usecase

import (
	"context"
	"fmt"
	"forecast-app-interface/internal/controller/gen/go/predict"
	"google.golang.org/grpc/status"
	"path/filepath"
	"strings"
	"sync"
)

const (
	getForecastErrorTemplate  = "cannot get forecast [name=%s]: %s"
	makeForecastErrorTemplate = "cannot get forecast [name=%s unit=%s period=%d]: %s"
)

var getPredictRequestPool = sync.Pool{
	New: func() any {
		return &predict.GetPredictRequest{}
	},
}

func (u *UseCase) GetForecast(username, name string) (string, int64, any, error) {
	request := getPredictRequestPool.Get().(*predict.GetPredictRequest)
	request.Username = username
	request.Name = name

	response, err := u.predictClient.GetPredict(context.Background(), request)
	if err != nil {
		err = fmt.Errorf(getForecastErrorTemplate, name, status.Convert(err).Message())
	}

	getPredictRequestPool.Put(request)

	if response == nil {
		return "", 0, nil, err
	}

	return response.Unit, response.Delimiter, response.Items, err
}

var makePredictRequestPool = sync.Pool{
	New: func() any {
		return &predict.MakePredictRequest{}
	},
}

var timeSeriesItemPool = sync.Pool{New: func() any {
	return &predict.TimeSeriesItem{}
}}

func (u *UseCase) MakeForecast(username, name, unit string, period, predictPeriods int32, tss []int64, values []float64) error {
	request := makePredictRequestPool.Get().(*predict.MakePredictRequest)

	request.Username = username
	request.Name = strings.TrimSuffix(name, filepath.Ext(name))
	request.Unit = unit
	request.Period = period
	request.PredictPeriods = predictPeriods
	request.Items = make([]*predict.TimeSeriesItem, 0, len(tss))

	for i := 0; i < len(tss); i++ {
		item := timeSeriesItemPool.Get().(*predict.TimeSeriesItem)

		item.Ts = tss[i]
		item.Value = values[i]

		request.Items = append(request.Items, item)
	}

	_, err := u.predictClient.MakePredict(context.Background(), request)
	if err != nil {
		err = fmt.Errorf(makeForecastErrorTemplate, name, unit, period, status.Convert(err).Message())
	}

	for i := range request.Items {
		timeSeriesItemPool.Put(request.Items[i])
	}
	getPredictRequestPool.Put(request)

	return err
}
