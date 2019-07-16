package main

func main() {
	consumer := consumer {}
	consumer.useRabbitMQ(RabbitMQ{})
	if err := consumer.rabbitMQ.dial("amqp://guest:guest@localhost:5672/"); err != nil {

	}
	if err := consumer.rabbitMQ.openChannel(); err != nil {

	}
	if err := consumer.rabbitMQ.declareExchange("logs_topics","topic"); err != nil {

	}
	if err := consumer.rabbitMQ.declareQueue("test"); err != nil {

	}
	consumer.rabbitMQ.getRoutingKeyConsumer()
	if err := consumer.rabbitMQ.queueBind(); err != nil {

	}
	consumer.register()
}

