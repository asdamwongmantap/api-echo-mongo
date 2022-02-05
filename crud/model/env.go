package model

import (
	"github.com/asdamwongmantap/api-echo-mongo/lib/db"
	"github.com/asdamwongmantap/api-echo-mongo/lib/queue"
)

type (
	EnvConfig struct {
		Host     string
		Port     int
		Mongo    db.MongoConfig
		RabbitMQ queue.QueueConfig
	}
)
