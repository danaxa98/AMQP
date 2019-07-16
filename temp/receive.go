package temp

import "Loopline/pkg"

func main() {
	consumer := pkg.Consumer{}
	pkg.UseRabbitMQ(pkg.RabbitMQ{})
	if err := pkg.Dial("amqp://guest:guest@localhost:5672/"); err != nil {

	}
	if err := pkg.OpenChannel(); err != nil {

	}
	if err := pkg.DeclareExchange("logs_topics","topic"); err != nil {

	}
	if err := pkg.DeclareQueue("test"); err != nil {

	}
	pkg.GetRoutingKeyConsumer()
	if err := pkg.QueueBind(); err != nil {

	}
	pkg.Register()
}

