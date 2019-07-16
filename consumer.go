package main

import "fmt"

type consumer struct {
	rabbitMQ RabbitMQ
	Name     string
}

func (consumer consumer) useRabbitMQ (r RabbitMQ) {
	consumer.rabbitMQ = r
}

func (consumer consumer) register () {
	//todo check for non-initialized field!
	consumer.rabbitMQ.register(consumer.Name, consumer.handle)
}

func (consumer consumer) handle (msg []byte) {
	fmt.Print("MESSAGE", string(msg))
}