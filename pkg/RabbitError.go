package pkg

type RabbitError string

const(
	ServerError   RabbitError = "Unable to connect to RabbitMQ server!"
	ChannelError  RabbitError = "Unable to open a channel!"
	ExchangeError RabbitError = "Unable to declare an exchange!"
	QueueError    RabbitError = "Unable to declare a queue!"
	BindError     RabbitError = "Unable to bind the queue!"
	RegistryError RabbitError = "Unable to register a consumer!"
	PublishError  RabbitError = "Unable to publish the message!"
)