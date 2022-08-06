package main

import (
	"fmt"
	"log"
	"net/http"

	httpDelivery "github.com/asdamwongmantap/api-echo-mongo/crud/delivery/http"
	"github.com/asdamwongmantap/api-echo-mongo/crud/model"
	"github.com/asdamwongmantap/api-echo-mongo/crud/repository"
	"github.com/asdamwongmantap/api-echo-mongo/crud/usecase"
	"github.com/asdamwongmantap/api-echo-mongo/lib/config"
	"github.com/asdamwongmantap/api-echo-mongo/lib/db"
	"github.com/asdamwongmantap/api-echo-mongo/lib/logging"
	"github.com/asdamwongmantap/api-echo-mongo/lib/queue"
	"github.com/labstack/echo/v4/middleware"

	"github.com/labstack/echo/v4"
)

//to initialize viper config
func init() {
	config.SetConfigFile("config", "lib/config", "json")
}

func main() {
	envConfig := getConfig()

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))

	e.Use(logging.MiddlewareLogging)

	// Mongo
	mongo, err := db.Connect(envConfig.Mongo)
	if err != nil {
		log.Println(err)
		return
	}

	// Queue
	rabbitMq, err := queue.Connect(envConfig.RabbitMQ)
	if err != nil {
		log.Println(err)
		return
	}

	crudRepo := repository.NewCrudRepository(mongo)
	crudUseCase := usecase.NewCrudUseCase(&envConfig, crudRepo, rabbitMq)
	// Router
	httpDelivery.NewRouter(e, crudUseCase)

	e.Logger.Fatal(e.Start(fmt.Sprintf("%s%s%v", envConfig.Host, ":", envConfig.Port)))
}

func getConfig() model.EnvConfig {

	return model.EnvConfig{
		Host: config.GetString("host.address"),
		Port: config.GetInt("host.port"),
		Mongo: db.MongoConfig{
			Timeout:  config.GetInt("database.mongodb.timeout"),
			DBname:   config.GetString("database.mongodb.dbname"),
			Username: config.GetString("database.mongodb.user"),
			Password: config.GetString("database.mongodb.password"),
			Host:     config.GetString("database.mongodb.host"),
			Port:     config.GetString("database.mongodb.port"),
		},
		RabbitMQ: queue.QueueConfig{
			Timeout:  config.GetInt("queue.rabbit.timeout"),
			Username: config.GetString("queue.rabbit.user"),
			Password: config.GetString("queue.rabbit.password"),
			Host:     config.GetString("queue.rabbit.host"),
			Port:     config.GetString("queue.rabbit.port"),
		},
	}
}
