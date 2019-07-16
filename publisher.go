package main


type publisher struct {
	rabbitMQ RabbitMQ
	Name     string
}

func (publisher publisher) UseRabbitMQ (r RabbitMQ) {
	publisher.rabbitMQ = r
}

func (publisher publisher) rename (name string) {
	publisher.Name = name
}