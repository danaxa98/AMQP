package pkg

import (
	"github.com/streadway/amqp"
	"log"
	"strings"
)

//A wrapper for RabbitMQ
type RabbitMQ struct {
	Connection		*amqp.Connection
	Channel			*amqp.Channel
	Queue			*amqp.Queue
	ExchangeName	string
	RoutingKey		string
	Body			string
}

func (rabbit *RabbitMQ) Dial(url string) error {
	var err error
	rabbit.Connection, err = amqp.Dial(url)
	if err != nil {
		defer rabbit.Connection.Close()
		return err
	}
	return nil
}

func (rabbit *RabbitMQ) OpenChannel() error {
	var err error
	rabbit.Channel, err = rabbit.Connection.Channel()
	if err != nil {
		defer rabbit.Channel.Close()
		return err
	}
	return nil
}

func (rabbit *RabbitMQ) DeclareExchange(exchangeName string, kind string) error {
	err := rabbit.Channel.ExchangeDeclare(
		exchangeName,
		kind,
		true,
		false,
		false,
		false,
		nil,
	)
	rabbit.ExchangeName = exchangeName
	return err
}

func (rabbit *RabbitMQ) DeclareQueue(queueName string) error {

	temp, err := rabbit.Channel.QueueDeclare(
		queueName,
		false,
		false,
		true,
		false,
		nil,
	)
	if err != nil {
		return err
	}
	rabbit.Queue = &temp
	return nil
}

func (rabbit *RabbitMQ) QueueBind() error{
	if rabbit.Queue == nil {
		//todo check for empty queue!
	}
	if rabbit.ExchangeName == "" {
		//todo check for empty exchange!
	}
	if rabbit.RoutingKey == "" {
		//todo check for empty routing key!
	}

		return rabbit.Channel.QueueBind(
			rabbit.Queue.Name, // queue Name
			rabbit.RoutingKey, // routing key
			"logs_topics",
			false,
			nil,
		)
}

func (rabbit *RabbitMQ) GetRoutingKeyPublisher (args []string) {
	var s string
	if (len(args) < 2) || args[1] == "" {
		s = "anonymous.info"
		//todo check for invalid syntax!
	} else {
		s = args[1]
	}
	rabbit.RoutingKey = s
}

func (rabbit *RabbitMQ) GetRoutingKeyConsumer() {
	//if rabbit.Queue == nil {
	//	//	//todo check for empty queue
	//	//}
	//	//if rabbit.ExchangeName == "" {
	//	//	//todo check for empty exchange!
	//	//}
	//	//
	//	//if len(args) < 2 {
	//	//	log.Printf("Usage: %s [binding_key]...", args[0])
	//	//	os.Exit(0)
	//	//}
	//	//
	//	//log.Println(args)
	//	//for _, s := range args[:1] {
	//	//	rabbit.RoutingKey = s
	//	//	log.Printf("Binding queue %s to exchange %s with routing key %s",
	//	//		rabbit.Queue.Name, rabbit.ExchangeName, s)
	//	//	rabbit.queueBind()
	//	//}
	rabbit.RoutingKey = "hello.world"
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

func (rabbit *RabbitMQ) publish (contentType string) error {
	if rabbit.Channel == nil {
		//todo check for empty channel!
	}
	if rabbit.ExchangeName == "" {
		//todo check for empty exchange!
	}
	if rabbit.Body == "" {
		//todo check for empty body!
	}
	if rabbit.RoutingKey == "" {
		//todo check for empty routing key!
	}

	if contentType == "" { contentType = "text/plain" }
	err := rabbit.Channel.Publish(
		rabbit.ExchangeName,
		rabbit.RoutingKey,
		false,
		false,
		amqp.Publishing {
			ContentType: 	contentType,
			Body:			[]byte(rabbit.Body),
		})
	if err != nil{
		return err
	}
	log.Printf(" [x] sent %s", rabbit.Body)
	return nil
}

func (rabbit *RabbitMQ) register(consumerName string, callback func(msg []byte)) error {
	if rabbit.Channel == nil {
		//todo check for empty channel!
	}
	if rabbit.Queue == nil {
		//todo check for empty queue!
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
		return err
	}
	rabbit.listen(messages, callback)
	return nil
}

//todo if necessary, this function should return byte[] in order for Consumer to process the body message further
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