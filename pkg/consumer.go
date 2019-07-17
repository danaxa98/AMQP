package pkg

import "fmt"

type Consumer struct {
	RabbitMQ RabbitMQ
	Name     string
}

func (consumer Consumer) UseRabbitMQ(r RabbitMQ) {
	consumer.RabbitMQ = r
}

func (consumer Consumer) Register() {
	consumer.RabbitMQ.register(consumer.Name, consumer.handle)
}

func (consumer Consumer) handle (msg []byte) {
	fmt.Printf("Message \"%s\" is received.\n", string(msg))
}