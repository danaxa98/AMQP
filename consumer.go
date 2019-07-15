package main

import(
	"github.com/streadway/amqp"
	"log"
)

type consumer struct {
	Name		string
	Log 		[]string
}

func (c consumer) register(queue amqp.Queue, ch amqp.Channel) {
	msg, err := ch.Consume(
		queue.Name,	//queue
		c.Name,		//consumer
		false,
		false,
		false,
		false,
		nil,
	)
	c.error(err)
	c.receive(msg)
}

func (c consumer) receive(msg <-chan amqp.Delivery) {
	go func() {
		for d := range msg {
			c.Log = append(c.Log, string(d.Body))
		}
	}()
}

func (c consumer) printAll() {
	for _, v := range c.Log {
		log.Println(v)
	}
}

func (c consumer) deleteAll() {
	c.Log = nil
}

func (c consumer) error(err error) {
	if err != nil {
		log.Fatalf("Fail to register consumer %s: %s", c.Name, err)
	}
}
