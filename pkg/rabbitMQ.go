package pkg

import (
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

func (rabbit *RabbitMQ) QueueBind(exchangeName string, routingKey []string) error {

	var err error
	for _, key := range routingKey {
		err = rabbit.Channel.QueueBind(
			rabbit.Queue.Name,  // queue Name
			key, // routing key
			exchangeName,
			false,
			nil,
		)
	}
	return err
}

func (rabbit *RabbitMQ) Register(exchangeName string, callback func(msg []byte), consumerName string, routingKeys []string) error {
	var err error

	if err = rabbit.QueueBind(exchangeName, routingKeys); err != nil {
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

	log.Printf("Consumer %s successfully registered with routing key(s) %v.", strconv.Quote(consumerName), routingKeysString)
	rabbit.listen(messages, callback)

	return err
}

func (rabbit *RabbitMQ) listen (messages <- chan amqp.Delivery, callback func(msg []byte)) {

	go func() {
		for msg := range messages {
			callback(msg.Body)
		}
	}()
}