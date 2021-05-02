package controller

import (
	"context"
	"github.com/asdamwongmantap/api-echo-mongo/crud"
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
