package main

import (
	"fmt"
	httpDelivery "github.com/asdamwongmantap/api-echo-mongo/crud/delivery/http"
	"github.com/asdamwongmantap/api-echo-mongo/crud/model"
	"github.com/asdamwongmantap/api-echo-mongo/crud/repository"
	"github.com/asdamwongmantap/api-echo-mongo/crud/usecase"
	"github.com/asdamwongmantap/api-echo-mongo/lib/db"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	envConfig := getConfig()

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))

	// Mongo
	mongo, err := db.Connect(envConfig.Mongo)
	if err != nil {
		log.Println(err)
		return
	}

	crudRepo := repository.NewCrudRepository(mongo)
	crudUseCase := usecase.NewCrudUseCase(&envConfig, crudRepo)
	// Router
	httpDelivery.NewRouter(e, crudUseCase)

	e.Logger.Fatal(e.Start(fmt.Sprintf("%s%s%v",envConfig.Host,":",envConfig.Port)))
}

func getConfig() model.EnvConfig {

	return model.EnvConfig{
			Host: "0.0.0.0",
			Port: 9595,
			Mongo: db.MongoConfig{
			Timeout:  5000,
			DBname:   "crud_learn",
			Username: "",
			Password: "",
			Host:     "0.0.0.0",
			Port:     "27017",
		},
	}
}
