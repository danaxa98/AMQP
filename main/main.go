package main

import (
	"Loopline/pkg"
)

func main() {
	consumer := pkg.Consumer{}
	consumer.UseRabbitMQ(pkg.RabbitMQ{})
	if err := consumer.RabbitMQ.Dial("amqp://guest:guest@localhost:5672/"); err != nil {
		panic(pkg.RabbitError(pkg.ServerError))
	}
	if err := consumer.RabbitMQ.OpenChannel(); err != nil {
		panic(pkg.RabbitError(pkg.ChannelError))
	}
	if err := consumer.RabbitMQ.DeclareExchange("logs_topics","topic"); err != nil {
		panic(pkg.RabbitError(pkg.ExchangeError))
	}
	if err := consumer.RabbitMQ.DeclareQueue("test"); err != nil {
		panic(pkg.RabbitError(pkg.QueueError))
	}
	consumer.RabbitMQ.GetRoutingKeyConsumer()
	if err := consumer.RabbitMQ.QueueBind(); err != nil {
		panic(pkg.RabbitError(pkg.BindError))
	}
	consumer.Register()
}
