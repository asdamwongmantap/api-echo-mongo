package controller

import (
	"context"
	"encoding/json"
	"github.com/asdamwongmantap/api-echo-mongo/crud"
	"github.com/asdamwongmantap/api-echo-mongo/crud/model"
	"github.com/labstack/echo/v4"
)

type CrudController struct {
	e       *echo.Echo
	usecase crud.CrudUseCaseI
}

func NewCrudController(e *echo.Echo, usecase crud.CrudUseCaseI) *CrudController {
	return &CrudController{
		e:       e,
		usecase: usecase,
	}
}

func (cc *CrudController) GetData(ec echo.Context) error {

	data, err := cc.usecase.GetDataUC(context.Background())
	if err != nil {
		return err
	}

	return ec.JSON(200, data)
}

func (cc *CrudController) InsertData(ec echo.Context) error {

	var req model.DataProduct
	err := json.NewDecoder(ec.Request().Body).Decode(&req)
	if err != nil {
		return err
	}
	data, err := cc.usecase.InsertDataUC(context.Background(), req)
	if err != nil {
		return err
	}

	return ec.JSON(200, data)
}
