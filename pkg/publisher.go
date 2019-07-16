package pkg


type Publisher struct {
	rabbitMQ RabbitMQ
	Name     string
}

func (publisher Publisher) UseRabbitMQ (r RabbitMQ) {
	publisher.rabbitMQ = r
}

func (publisher Publisher) rename (name string) {
	publisher.Name = name
}