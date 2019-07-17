package pkg

type RabbitError string

const(
	ConnectServerError  	RabbitError = "Unable to connect to RabbitMQ server!"
	OpenChannelError    	RabbitError = "Unable to open a channel!"
	DeclareExchangeError	RabbitError = "Unable to declare an exchange!"
	DeclareQueueError  		RabbitError = "Unable to declare a queue!"
	BindQueueError      	RabbitError = "Unable to bind the queue!"
	RegistryError      	  	RabbitError = "Unable to register a consumer!"
	PublishError       	 	RabbitError = "Unable to publish the message!"
	EmptyChannel			RabbitError = "Empty channel!"
	EmptyQueue			 	RabbitError = "Empty queue!"
	EmptyExchange		 	RabbitError = "Empty exchange!"
	Default              	RabbitError = ""
)