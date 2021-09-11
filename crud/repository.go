package crud

import (
	"context"
	"github.com/asdamwongmantap/api-echo-mongo/crud/model"
)

type CrudRepositoryI interface {
	GetAllData(ctx context.Context) (crudResp model.GetDataResponse, err error)
	InsertData(ctx context.Context, req model.DataProductRequest) error
	UpdateData(ctx context.Context, req model.DataProductRequest) error
	DeleteData(ctx context.Context, req model.DataProductRequest) error
}
