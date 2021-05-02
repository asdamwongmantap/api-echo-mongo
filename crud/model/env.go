package model

import "github.com/asdamwongmantap/api-echo-mongo/lib/db"

type (
	EnvConfig struct {
		Host string
		Port int
		Mongo db.MongoConfig
	}
)
