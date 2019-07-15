package main

import (
	"github.com/streadway/amqp"
	"log"
)

type publisher struct {
	Name         string
	ExchangeName string
	RoutingKey  string
	Message      string
}

func (p publisher) changeRoutingKey(keys string) {
	p.RoutingKey = keys
}

func (p publisher) publish(ch *amqp.Channel, exchangeName string, body string) {
	err := ch.Publish(
		exchangeName,
		p.RoutingKey,
		false,
		false,
		amqp.Publishing {
			ContentType: 	"text/plain",
			Body:			[]byte(body),
		})
	p.error(err)
	log.Printf(" [x] message: \"%s\" is sent", body)
}

func (p publisher) error(err error) {
	if err != nil {
		log.Fatalf("Fail to publish a Message %s: %s", p.Message, err)
	}
}