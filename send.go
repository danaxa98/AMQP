package main

import (
	"github.com/streadway/amqp"
	"log"
	"os"
	"strings"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func bodyFrom(args []string) string {
	var s string
	if len(args) < 3 || os.Args[2] == "" {
		s = "Hello"
	} else {
		s = strings.Join(args[2:], " ")
	}

	return s
}

func getRoutingKey(args []string) string{
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
		s = "anonymous.info"
	} else {
		s = os.Args[1]
	}
	return s
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	exchangeName := "logs_topic"
	err = ch.ExchangeDeclare(
		exchangeName,
		"topic",
		true,
		false,
		false,
		false,
		nil,
		)
	failOnError(err, "Failed to declare a exchange")

	body := bodyFrom(os.Args)

	//c := publisher{
	//	Name: "publisher 1",
	//}
	//c.changeRoutingKey(getRoutingKey(os.Args))
	//c.publish(ch, exchangeName, body)

	err = ch.Publish(
		exchangeName,
		getRoutingKey(os.Args),	//Routing key
		false,
		false,
		amqp.Publishing {
			ContentType: 	"text/plain",
			Body:			[]byte(body),
		})

	failOnError(err, "Failed to publish a Message")
	log.Printf(" [x] sent %s", body)
}