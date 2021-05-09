package usecase

import (
	"context"
	"github.com/asdamwongmantap/api-echo-mongo/crud"
	"github.com/asdamwongmantap/api-echo-mongo/crud/model"
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
		return resp, err
	}

	return list, err
}

func (cuc *CrudUseCase) InsertDataUC(ctx context.Context, req model.DataProductRequest) (resp bool, err error) {
	//check if context is nil
	if ctx == nil {
		ctx = context.Background()
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

	return true, nil
}
