package crud

import (
	"context"
	"github.com/asdamwongmantap/api-echo-mongo/crud/model"
)

type CrudUseCaseI interface {
	GetDataUC(ctx context.Context) (resp model.GetDataResponse, err error)
}
