package helper

import (
	"ScheduleApiGo/logger"
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/streadway/amqp"
)

type Rabbit struct {
	User     string
	Password string
	Host     string
	Port     uint32
}

func (r *Rabbit) HasEmptyParams() bool {
	return r.User == "" || r.Password == "" || r.Host == "" || r.Port == 0
}

func (r *Rabbit) Connection() (*amqp.Connection, error) {
	connectionString := r.GetStringConnection()
	con, err := amqp.Dial(connectionString)
	if err != nil {
		log.Printf("Failed to connect to RabbitMQ: %s", err)
		return nil, err
	}
	log.Println("Successfully connected to RabbitMQ")
	return con, nil
}

func (r *Rabbit) TestConnection() (bool, error) {
	_, err := amqp.Dial(r.GetStringConnection())
	if err != nil {
		return false, fmt.Errorf("erro na conexão: %v", err)
	}
	return true, nil
}

func (r *Rabbit) SendMessage(message interface{}, queue string, con *amqp.Connection) error {
	jobJSON, err := json.Marshal(message)
	if err != nil {
		logger.Log.Errorf("Error al serializar json: %v", err)
		return err
	}
	defer con.Close()

	ch, err := con.Channel()
	if err != nil {
		logger.Log.Error(err.Error())
		return err
	}
	defer ch.Close()

	logger.Log.Info("Declare queue: " + queue)
	_, err = ch.QueueDeclare(
		queue,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		logger.Log.Error(err.Error())
		return err
	}

	err = ch.Publish(
		"",
		queue,
		true,
		false,
		amqp.Publishing{
			ContentType:  "text/plain",
			Body:         jobJSON,
			DeliveryMode: amqp.Persistent,
		},
	)
	if err != nil {
		logger.Log.Error(err.Error())
		return err
	}
	logger.Log.Info("Send message with successfull!")
	logger.Log.Info("Message:")
	logger.Log.Info(message)

	return nil
}

func (r *Rabbit) GetStringConnection() string {
	return "amqp://" + r.User + ":" + r.Password + "@" + r.Host + ":" + strconv.Itoa(int(r.Port)) + "/"
}
