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
	//todo check for non-initialized field!
	consumer.RabbitMQ.register(consumer.Name, consumer.handle)
}

func (consumer Consumer) handle (msg []byte) {
	fmt.Print("MESSAGE", string(msg))
}