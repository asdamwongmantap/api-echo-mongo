package crud

import (
	"context"
	"github.com/asdamwongmantap/api-echo-mongo/crud/model"
)

type CrudUseCaseI interface {
	GetDataUC(ctx context.Context) (resp model.GetDataResponse, err error)
	InsertDataUC(ctx context.Context, req model.DataProductRequest) (resp bool, err error)
	UpdateDataUC(ctx context.Context, req model.DataProductRequest) (resp bool, err error)
	DeleteDataUC(ctx context.Context, req model.DataProductRequest) (resp bool, err error)
}
