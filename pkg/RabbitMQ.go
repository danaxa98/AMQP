package pkg

import (
	"github.com/streadway/amqp"
	"log"
	"strings"
)

//A wrapper for RabbitMQ
type RabbitMQ struct {
	Connection   *amqp.Connection
	Channel      *amqp.Channel
	Queue        *amqp.Queue
	ExchangeName string
	RoutingKeys  string
	Body         string
}

func (rabbit *RabbitMQ) Dial(url string) RabbitError {
	var err error
	rabbit.Connection, err = amqp.Dial(url)
	if err != nil {
		defer rabbit.Connection.Close()
		return ConnectServerError
	}
	return Default
}

func (rabbit *RabbitMQ) OpenChannel() RabbitError {
	var err error
	rabbit.Channel, err = rabbit.Connection.Channel()
	if err != nil {
		defer rabbit.Channel.Close()
		return OpenChannelError
	}
	return Default
}

func (rabbit *RabbitMQ) DeclareExchange(exchangeName string, kind string) RabbitError {
	err := rabbit.Channel.ExchangeDeclare(
		exchangeName,
		kind,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return DeclareExchangeError
	}
	rabbit.ExchangeName = exchangeName
	return Default
}

func (rabbit *RabbitMQ) DeclareQueue(queueName string) RabbitError {

	temp, err := rabbit.Channel.QueueDeclare(
		queueName,
		false,
		false,
		true,
		false,
		nil,
	)
	if err != nil {
		return DeclareQueueError
	}
	rabbit.Queue = &temp
	return Default
}

func (rabbit *RabbitMQ) QueueBind() RabbitError {
	if rabbit.Queue == nil {
		return EmptyQueue
	}
	if rabbit.ExchangeName == "" {
		return EmptyExchange
	}
	if rabbit.RoutingKeys == "" {
		return EmptyRoutingKeys
	}

		err := rabbit.Channel.QueueBind(
			rabbit.Queue.Name,  // queue Name
			rabbit.RoutingKeys, // routing key
			"logs_topics",
			false,
			nil,
		)
		if err != nil {
			return BindQueueError
		}
		return Default
}

func (rabbit *RabbitMQ) GetRoutingKeyPublisher (args []string) {
	var s string
	if (len(args) < 2) || args[1] == "" {
		s = "anonymous.info"
		//todo check for invalid syntax!
	} else {
		s = args[1]
	}
	rabbit.RoutingKeys = s
}

func (rabbit *RabbitMQ) SetRoutingKeyConsumer(key string) RabbitError{
	//if rabbit.Queue == nil {
	//
	//	//}
	//	//if rabbit.ExchangeName == "" {
	//
	//	//}
	//	//
	//	//if len(args) < 2 {
	//	//	log.Printf("Usage: %s [binding_key]...", args[0])
	//	//	os.Exit(0)
	//	//}
	//	//
	//	//log.Println(args)
	//	//for _, s := range args[:1] {
	//	//	rabbit.RoutingKeys = s
	//	//	log.Printf("Binding queue %s to exchange %s with routing key %s",
	//	//		rabbit.Queue.Name, rabbit.ExchangeName, s)
	//	//	rabbit.queueBind()
	//	//}
	if key == "" {
		return EmptyRoutingKeys
	}
	rabbit.RoutingKeys = key
	return Default
}

func (rabbit *RabbitMQ) getBody (args []string) {
	var s string
	if len(args) < 3 || args[2] == "" {
		s = "Hello"
	} else {
		s = strings.Join(args[2:], " ")
	}

	rabbit.Body = s
}

func (rabbit *RabbitMQ) publish (contentType string) RabbitError {
	if rabbit.Channel == nil {
		return EmptyChannel
	}
	if rabbit.ExchangeName == "" {
		return EmptyExchange
	}
	if rabbit.Body == "" {
		return EmptyBody
	}
	if rabbit.RoutingKeys == "" {
		return EmptyRoutingKeys
	}

	if contentType == "" { contentType = "text/plain" }
	err := rabbit.Channel.Publish(
		rabbit.ExchangeName,
		rabbit.RoutingKeys,
		false,
		false,
		amqp.Publishing {
			ContentType: 	contentType,
			Body:			[]byte(rabbit.Body),
		})
	if err != nil{
		return PublishError
	}
	log.Printf(" [x] sent %s", rabbit.Body)
	return Default
}

func (rabbit *RabbitMQ) register(consumerName string, callback func(msg []byte)) RabbitError {
	if rabbit.Channel == nil {
		return EmptyChannel
	}
	if rabbit.Queue == nil {
		return EmptyQueue
	}

	var messages <- chan amqp.Delivery
	var err error
	messages, err = rabbit.Channel.Consume(
		rabbit.Queue.Name,
		consumerName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return RegistryError
	}
	rabbit.listen(messages, callback)
	return Default
}

func (rabbit *RabbitMQ) listen (messages <- chan amqp.Delivery, callback func(msg []byte)) {
	forever := make(chan bool)

	go func() {
		for msg := range messages {
			callback(msg.Body)

		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<- forever
}