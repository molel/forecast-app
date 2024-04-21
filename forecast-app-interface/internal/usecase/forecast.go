package usecase

import (
	"context"
	"fmt"
	"forecast-app-interface/internal/controller/gen/go/predict"
	"sync"
)

var GetPredictRequestPool = sync.Pool{
	New: func() any {
		return &predict.GetPredictRequest{}
	},
}

var MakePredictRequestPool = sync.Pool{
	New: func() any {
		return &predict.MakePredictRequest{}
	},
}

func (u *UseCase) MakeForecast(username, name, unit string, period int32, series []byte) (any, error) {
	request := MakePredictRequestPool.Get().(*predict.MakePredictRequest)
	request.Username = username
	request.Name = name
	request.Unit = unit
	request.Period = period
	request.Series = series

	response, err := u.predictClient.MakePredict(context.Background(), request)
	if err != nil {
		err = fmt.Errorf(registerErrorTemplate, err)
	}

	GetPredictRequestPool.Put(request)

	return response.Items, err
}

func (u *UseCase) GetForecast(username, name string) (string, int64, any, error) {
	request := GetPredictRequestPool.Get().(*predict.GetPredictRequest)
	request.Username = username
	request.Name = name

	response, err := u.predictClient.GetPredict(context.Background(), request)
	if err != nil {
		err = fmt.Errorf(registerErrorTemplate, err)
	}

	GetPredictRequestPool.Put(request)

	if response == nil {
		return "", 0, nil, err
	}

	return response.Unit, response.Delimiter, response.Items, err
}
