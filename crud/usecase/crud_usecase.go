package usecase

import (
	"context"
	"fmt"
	"log"

	"github.com/asdamwongmantap/api-echo-mongo/crud"
	"github.com/asdamwongmantap/api-echo-mongo/crud/model"
	"github.com/asdamwongmantap/api-echo-mongo/lib/logging"
	"github.com/pkg/errors"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
)

type CrudUseCase struct {
	config   *model.EnvConfig
	crudRepo crud.CrudRepositoryI
	rabbitMq *amqp.Connection
}

func NewCrudUseCase(config *model.EnvConfig, crudRepo crud.CrudRepositoryI, rabbitMq *amqp.Connection) crud.CrudUseCaseI {
	return &CrudUseCase{
		config:   config,
		crudRepo: crudRepo,
		rabbitMq: rabbitMq,
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

	//publish rabbit mq
	ch, err := cuc.rabbitMq.Channel()
	if err != nil {
		fmt.Println(err)
	}
	defer ch.Close()

	// with this channel open, we can then start to interact
	// with the instance and declare Queues that we can publish and
	// subscribe to
	q, err := ch.QueueDeclare(
		"MsgRabbitGo",
		false,
		false,
		false,
		false,
		nil,
	)
	// We can print out the status of our Queue here
	// this will information like the amount of messages on
	// the queue
	fmt.Println(q)
	// Handle any errors if we were unable to create the queue
	if err != nil {
		fmt.Println(err)
	}

	// attempt to publish a message to the queue!
	err = ch.Publish(
		"",
		"MsgRabbitGo",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("Hello World"),
		},
	)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Published Message to Queue")

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
