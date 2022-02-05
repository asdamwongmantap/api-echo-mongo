package queue

import (
	"fmt"

	"github.com/streadway/amqp"
)

type QueueConfig struct {
	Timeout  int
	DBname   string
	Username string
	Password string
	Host     string
	Port     string
}

func Connect(c QueueConfig) (*amqp.Connection, error) {
	connPattern := "amqp://%v:%v@%v:%v"
	if c.Username == "" {
		connPattern = "amqp://%s%s%v:%v"
	}

	clientUrl := fmt.Sprintf(connPattern,
		c.Username,
		c.Password,
		c.Host,
		c.Port,
	)
	conn, err := amqp.Dial(clientUrl)
	if err != nil {
		fmt.Println("Failed Initializing Broker Connection")
		panic(err)
	}

	return conn, err
}
