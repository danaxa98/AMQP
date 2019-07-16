package main

import (
	"Loopline/pkg"
)

func main() {
	consumer := pkg.Consumer{}
	consumer.UseRabbitMQ(pkg.RabbitMQ{})
	if err := consumer.RabbitMQ.Dial("amqp://guest:guest@localhost:5672/"); err != pkg.Default {
		panic(err)
	}
	if err := consumer.RabbitMQ.OpenChannel(); err != pkg.Default {
		panic(err)
	}
	if err := consumer.RabbitMQ.DeclareExchange("logs_topics","topic"); err != pkg.Default {
		panic(err)
	}
	if err := consumer.RabbitMQ.DeclareQueue("test"); err != pkg.Default {
		panic(err)
	}
	consumer.RabbitMQ.GetRoutingKeyConsumer()
	if err := consumer.RabbitMQ.QueueBind(); err != pkg.Default {
		panic(err)
	}
	consumer.Register()
}
