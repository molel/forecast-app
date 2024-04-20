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

func (u *UseCase) MakeForecast(username, name string, data []byte) ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UseCase) GetForecast(username, name string) ([]byte, error) {
	request := GetPredictRequestPool.Get().(*predict.GetPredictRequest)
	request.Username = username
	request.Name = name

	response, err := u.predictClient.GetPredict(context.Background(), request)
	if err != nil {
		err = fmt.Errorf(registerErrorTemplate, err)
	}

	GetPredictRequestPool.Put(request)

	return response.Data, err
}
