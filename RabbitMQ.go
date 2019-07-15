package main

import (
	"github.com/streadway/amqp"
	"log"
	"os"
	"strings"
)

//A wrapper for RabbitMQ
type RabbitMQ struct {
	Connection		*amqp.Connection
	Channel			*amqp.Channel
	Queue			*amqp.Queue
	ExchangeName	string
	RoutingKey		string
	Error			error
	Body			string
}

func (r RabbitMQ) dial (url string) {
	r.Connection, r.Error = amqp.Dial(url)
	r.failOnError("Failed to connect to RabbitMQ")
	defer r.Connection.Close()
}

func (r RabbitMQ) openChannel () {
	r.Channel, r.Error = r.Connection.Channel()
	r.failOnError("Failed to open a channel")
	defer r.Channel.Close()
}

func (r RabbitMQ) declareExchange (exchangeName string) {
	r.Error = r.Channel.ExchangeDeclare(
		exchangeName,
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)
	r.failOnError("Failed to declare a exchange")
	r.ExchangeName = exchangeName
}

func (r RabbitMQ) declareQueue (queueName string) {
	*r.Queue, r.Error = r.Channel.QueueDeclare(
		queueName,
		false,
		false,
		true,
		false,
		nil,
	)
	r.failOnError("Failed to declare a queue")
}

func (r RabbitMQ) bindQueue () {
	if r.Queue == nil {
		//todo check for empty queue!
	}
	if r.ExchangeName == "" {
		//todo check for empty exchange!
	}
	if r.RoutingKey == "" {
		//todo check for empty routing key!
	}

		r.Error = r.Channel.QueueBind(
			r.Queue.Name, 			// queue Name
			r.RoutingKey,          			// routing key
			"logs_topic",
			false,
			nil,
		)
		r.failOnError( "Failed to bind a queue")
}

func (r RabbitMQ) getRoutingKeyPublisher (args []string) {
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
		s = "anonymous.info"
		//todo check for invalid syntax!
	} else {
		s = os.Args[1]
	}
	r.RoutingKey = s
}

func (r RabbitMQ) getRoutingKeyConsumer (args []string) {
	if r.Queue == nil {
		//todo check for empty queue!
	}
	if r.ExchangeName == "" {
		//todo check for empty exchange!
	}

	for _, s := range args {
		r.RoutingKey = s
		log.Printf("Binding queue %s to exchange %s with routing key %s",
			r.Queue.Name, r.ExchangeName, s)
	}
}

func (r RabbitMQ) getBody (args []string) {
	var s string
	if len(args) < 3 || os.Args[2] == "" {
		s = "Hello"
	} else {
		s = strings.Join(args[2:], " ")
	}

	r.Body = s
}

func (r RabbitMQ) publish (contentType string) {
	if r.Channel == nil {
		//todo check for empty channel!
	}
	if r.ExchangeName == "" {
		//todo check for empty exchange!
	}
	if r.Body == "" {
		//todo check for empty body!
	}
	if r.RoutingKey == "" {
		//todo check for empty routing key!
	}

	if contentType == "" { contentType = "text/plain" }
	r.Error = r.Channel.Publish(
		r.ExchangeName,
		r.RoutingKey,
		false,
		false,
		amqp.Publishing {
			ContentType: 	contentType,
			Body:			[]byte(r.Body),
		})

	r.failOnError("Failed to publish a Message")
	log.Printf(" [x] sent %s", r.Body)
}

func (r RabbitMQ) register(consumerName string) {
	if r.Channel == nil {
		//todo check for empty channel!
	}
	if r.Queue == nil {
		//todo check for empty queue!
	}

	var message <- chan amqp.Delivery
	message, r.Error = r.Channel.Consume(
		r.Queue.Name,
		consumerName,
		false,
		false,
		false,
		false,
		nil,
	)
	r.failOnError("Failed to register a Consumer")
	r.listen(message)
}

func (r RabbitMQ) listen (message <- chan amqp.Delivery) {
	forever := make(chan bool)

	go func() {
		for d := range message {
			log.Printf(" [x] %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<- forever
}

func (r RabbitMQ) failOnError (message string) {
	if r.Error != nil {
		log.Fatalf("%s: %s", message, r.Error)
	}
}