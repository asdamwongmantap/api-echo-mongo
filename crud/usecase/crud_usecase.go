package usecase

import (
	"context"
	"github.com/asdamwongmantap/api-echo-mongo/crud"
	"github.com/asdamwongmantap/api-echo-mongo/crud/model"
	"github.com/asdamwongmantap/api-echo-mongo/lib/logging"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"log"
)

type CrudUseCase struct {
	config   *model.EnvConfig
	crudRepo crud.CrudRepositoryI
}

func NewCrudUseCase(config *model.EnvConfig, crudRepo crud.CrudRepositoryI) crud.CrudUseCaseI {
	return &CrudUseCase{
		config:   config,
		crudRepo: crudRepo,
	}
}

func (cuc *CrudUseCase) GetDataUC(ctx context.Context) (resp model.GetDataResponse, err error) {
	if ctx == nil {
		ctx = context.Background()
	}
	list, err := cuc.crudRepo.GetAllData(ctx)
	if err != nil {
		log.Println("failed to show data product with default log")
		return list, err
	}

	return list, err
}

func (cuc *CrudUseCase) InsertDataUC(ctx context.Context, req model.DataProductRequest) (resp bool, err error) {
	//check if context is nil
	if ctx == nil {
		ctx = context.Background()
	}

	if req.ProductName == "" {
		err = errors.New("failed to add data product ")
		logging.Info(err)
		return false, err
	}

	//insert data
	err = cuc.crudRepo.InsertData(ctx, req)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (cuc *CrudUseCase) UpdateDataUC(ctx context.Context, req model.DataProductRequest) (resp bool, err error) {
	//check if context is nil
	if ctx == nil {
		ctx = context.Background()
	}

	//update data
	err = cuc.crudRepo.UpdateData(ctx, req)
	if err != nil {
		return false, err
	}
	zaplogger, _ := zap.NewProduction()
	defer zaplogger.Sync()
	zaplogger.Info("success to update product",
		zap.String("product ID", req.ProductID),
	)

	return true, nil
}

func (cuc *CrudUseCase) DeleteDataUC(ctx context.Context, req model.DataProductRequest) (resp bool, err error) {
	//check if context is nil
	if ctx == nil {
		ctx = context.Background()
	}

	//update data
	err = cuc.crudRepo.DeleteData(ctx, req)
	if err != nil {
		return false, err
	}

	return true, nil
}
