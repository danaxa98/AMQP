package pkg

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)


var RootCommand = &cobra.Command{
	Use:		"-listen",
	Version:	"1.0",
	Hidden:		false,
	Args:		cobra.PositionalArgs(cobra.ExactArgs(0)),
	Run: 		func(cmd *cobra.Command, args []string){
		consumer := Consumer{}
		consumer.UseRabbitMQ(RabbitMQ{})
		if err := consumer.RabbitMQ.Dial("amqp://guest:guest@localhost:5672/"); err != Default {
			panic(err)
		}
		if err := consumer.RabbitMQ.OpenChannel(); err != Default {
			panic(err)
		}
		if err := consumer.RabbitMQ.DeclareExchange("logs_topics","topic"); err != Default {
			panic(err)
		}
		if err := consumer.RabbitMQ.DeclareQueue("test"); err != Default {
			panic(err)
		}
		if err := consumer.RabbitMQ.SetRoutingKeyConsumer("hello.*"); err != Default {
			panic(err)
		}

		if err := consumer.RabbitMQ.QueueBind(); err != Default {
			panic(err)
		}
		consumer.Register()
	},
}

func Execute() {
	if err := RootCommand.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	//RootCommand.AddCommand(
	//	Dial,
	//	OpenChannel,
	//	DeclareExchange,
	//	DeclareQueue,
	//	SetRoutingKey,
	//	QueueBind,
	//	Register)
}

var Dial = &cobra.Command{
	Use:		"-dial",
	Version:	"1.0",
	Hidden:		false,
	Args:		cobra.PositionalArgs(cobra.ExactArgs(1)),
	Short:		"Command \"-dial\" makes a connection to RabbitMQ server",
	Long:		"Command \"-dial\" requires one argument and makes a connection to the RabbitMQ server",
	Example:	"\"-dial amqp://guest:guest@localhost:5672/\"",
	Run:		func(cmd *cobra.Command, args[] string) {
		//todo
	},
}

var OpenChannel = &cobra.Command{
	Use:		"-oc",
	Version:	"1.0",
	Hidden:		false,
	Args:		cobra.PositionalArgs(cobra.ExactArgs(0)),
	Short:		"Command \"-oc\" opens a channel",
	Long:		"Command \"-oc\" opens a channel",
	Example:	"-oc",
	Run:		func(cmd *cobra.Command, args[] string) {
		//todo
	},
}

var DeclareExchange = &cobra.Command{
	Use:		"-de",
	Version:	"1.0",
	Hidden:		false,
	Args:		cobra.PositionalArgs(cobra.ExactArgs(2)),
	Short:		"Command \"-de\" declares an exchange",
	Long:		"Command \"-de\" requires two arguments, name and type, and declares an exchange " +
		"with said attributes.",
	Example:	"-de test direct",
	Run:		func(cmd *cobra.Command, args[] string) {
		//todo
	},
}

var DeclareQueue = &cobra.Command{
	Use:		"-dq",
	Version: 	"1.0",
	Hidden:		false,
	Args:		cobra.PositionalArgs(cobra.ExactArgs(1)),
	Short:		"Command \"-dq\" declares a queue with the argument's name",
	Long:		"Command \"-dq\" requires one argument and declares a queue with the argument's name",
	Example:	"-dq test",
	Run:		func(cmd *cobra.Command, args[] string) {
		//todo
	},
}

var SetRoutingKey = &cobra.Command{
	Use:		"-srk",
	Version:	"1.0",
	Hidden:		false,
	Args:		cobra.PositionalArgs(cobra.ExactArgs(1)),
	Short:		"Command \"-srk\" sets the routing key",
	Long:		"Command \"-srk\" requires one argument and sets the routing key to that argument",
	Example:	"-srk hello.world",
	Run:		func(cmd *cobra.Command, args[] string) {
		//todo
	},
}

var QueueBind = &cobra.Command{
	Use:		"-qb",
	Version: 	"1.0",
	Hidden:		false,
	Args:		cobra.PositionalArgs(cobra.ExactArgs(0)),
	Short:		"Command \"-qb\" binds queue to the current routing keys",
	Long:		"Command \"-qb\" binds queue to the current routing keys",
	Example:	"-qb",
	Run:		func(cmd *cobra.Command, args[] string) {
		//todo
	},
}

var Register = &cobra.Command{
	Use:		"-reg",
	Version: 	"1.0",
	Hidden:		false,
	Args:		cobra.PositionalArgs(cobra.ExactArgs(0)),
	Short:		"Command \"-reg\" register a consumer",
	Long:		"Command \"-reg\" register a consumer and is read to listen to any incoming message",
	Example:	"-reg",
	Run:		func(cmd *cobra.Command, args[] string) {
		//todo
	},
}