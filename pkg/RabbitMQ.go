package pkg

import (
	"github.com/streadway/amqp"
	"log"
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

func (rabbit *RabbitMQ) Dial(url string) RabbitError {
	var err error
	rabbit.Connection, err = amqp.Dial(url)
	if err != nil {
		defer rabbit.Connection.Close()
		return ConnectServerError
	}
	return Default
}

func (rabbit *RabbitMQ) OpenChannel() RabbitError {
	var err error
	rabbit.Channel, err = rabbit.Connection.Channel()
	if err != nil {
		defer rabbit.Channel.Close()
		return OpenChannelError
	}
	return Default
}

func (rabbit *RabbitMQ) DeclareExchange(exchangeName string, kind string) RabbitError {
	err := rabbit.Channel.ExchangeDeclare(
		exchangeName,
		kind,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return DeclareExchangeError
	}
	rabbit.ExchangeName = exchangeName
	return Default
}

func (rabbit *RabbitMQ) DeclareQueue(queueName string) RabbitError {

	temp, err := rabbit.Channel.QueueDeclare(
		queueName,
		false,
		false,
		true,
		false,
		nil,
	)
	if err != nil {
		return DeclareQueueError
	}
	rabbit.Queue = &temp
	return Default
}

func (rabbit *RabbitMQ) QueueBind(routingKey string) RabbitError {
	if rabbit.Queue == nil {
		return EmptyQueue
	}
	if rabbit.ExchangeName == "" {
		return EmptyExchange
	}

		err := rabbit.Channel.QueueBind(
			rabbit.Queue.Name,  // queue Name
			routingKey, // routing key
			rabbit.ExchangeName,
			false,
			nil,
		)
		if err != nil {
			return BindQueueError
		}
		return Default
}

func (rabbit *RabbitMQ) getBody (args []string) {
	var s string
	if len(args) < 3 || args[2] == "" {
		s = "Hello"
	} else {
		s = strings.Join(args[2:], " ")
	}

	rabbit.Body = s
}

func (rabbit *RabbitMQ) Register(routingKey string, consumerName string, callback func(msg []byte)) RabbitError {
	rabbit.QueueBind(routingKey)
	if rabbit.Channel == nil {
		return EmptyChannel
	}
	if rabbit.Queue == nil {
		return EmptyQueue
	}

	var messages <- chan amqp.Delivery
	var err error
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
		return RegistryError
	}
	rabbit.listen(messages, callback)
	return Default
}

func (rabbit *RabbitMQ) listen (messages <- chan amqp.Delivery, callback func(msg []byte)) {
	forever := make(chan bool)

	go func() {
		for msg := range messages {
			callback(msg.Body)

		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<- forever
}