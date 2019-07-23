package pkg

import (
	"database/sql"
	"github.com/streadway/amqp"
	"log"
	"strconv"
	"strings"
)

//Wrapper class for RabbitMQ
type RabbitMQ struct {
	Connection   *amqp.Connection
	Channel      *amqp.Channel
	Queue        *amqp.Queue
	ExchangeName string
	Body         string
}

func (rabbit *RabbitMQ) Dial(url string) error {
	var err error
	rabbit.Connection, err = amqp.Dial(url)
	return err
}

func (rabbit *RabbitMQ) OpenChannel() error {
	var err error
	rabbit.Channel, err = rabbit.Connection.Channel()
	return err
}

func (rabbit *RabbitMQ) DeclareExchange(exchangeName string, kind string) error {
	err := rabbit.Channel.ExchangeDeclare(
		exchangeName,
		kind,
		true,
		false,
		false,
		false,
		nil,
	)
	rabbit.ExchangeName = exchangeName
	return err
}

func (rabbit *RabbitMQ) DeclareQueue(queueName string) error {

	temp, err := rabbit.Channel.QueueDeclare(
		queueName,
		false,
		false,
		true,
		false,
		nil,
	)

	rabbit.Queue = &temp
	return err
}

func (rabbit *RabbitMQ) QueueBind(routingKey []string) error {

	var err error
	for _, key := range routingKey {
		err = rabbit.Channel.QueueBind(
			rabbit.Queue.Name,  // queue Name
			key, // routing key
			rabbit.ExchangeName,
			false,
			nil,
		)
	}
	return err
}

func (rabbit *RabbitMQ) Register(db *sql.DB, callback func(msg []byte) string, consumerName string, routingKeys []string) error {
	var err error

	if err = rabbit.QueueBind(routingKeys); err != nil {
		return err
	}

	var messages <- chan amqp.Delivery
	messages, err = rabbit.Channel.Consume(
		rabbit.Queue.Name,
		consumerName,
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return err
	}

	//Parsing debugging message with routing key name
	for index := range routingKeys {
		routingKeys[index] = strconv.Quote(routingKeys[index])
	}
	routingKeysString := strings.Join(routingKeys, ", ")

	log.Printf("Consumer %s successfully registered with routing key %v.", strconv.Quote(consumerName), routingKeysString)
	rabbit.listen(db, messages, callback)

	return err
}

func (rabbit *RabbitMQ) listen (db *sql.DB, messages <- chan amqp.Delivery, callback func(msg []byte) string) {

	go func() {
		for msg := range messages {
			toBeStoredMessage := callback(msg.Body)
			tx, err := db.Begin()
			checkError(err)

			_, err = tx.Exec("INSERT INTO amqp(message) VALUES(?)", toBeStoredMessage)
			if err != nil {
				tx.Rollback()
				log.Fatal(err)
			}
			checkError(tx.Commit())

		}
	}()
}