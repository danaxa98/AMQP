package pkg

import (
	"github.com/spf13/cobra"
)

type CommandFactory struct {
	Dial				cobra.Command
	OpenChannel 		cobra.Command
	DeclareExchange		cobra.Command
	DeclareQueue		cobra.Command
	SetRoutingKey		cobra.Command
	QueueBind			cobra.Command
	Register			cobra.Command
}

func (cf *CommandFactory) config() {
	cf.Dial = cobra.Command{
		Use:		"-dial",
		Version:	"1.0",
		Hidden:		false,
		Args:		cobra.PositionalArgs(cobra.ExactArgs(1)),
		Short:		"Command \"-dial\" makes a connection to RabbitMQ server",
		Long:		"Command \"-dial\" requires one argument and makes a connection to the RabbitMQ server",
		Example:	"\"-dial amqp://guest:guest@localhost:5672/\"",
	}

	cf.OpenChannel = cobra.Command{
		Use:		"-oc",
		Version:	"1.0",
		Hidden:		false,
		Args:		cobra.PositionalArgs(cobra.ExactArgs(0)),
		Short:		"Command \"-oc\" opens a channel",
		Long:		"Command \"-oc\" opens a channel",
		Example:	"-oc",
	}

	cf.DeclareExchange = cobra.Command{
		Use:		"-de",
		Version:	"1.0",
		Hidden:		false,
		Args:		cobra.PositionalArgs(cobra.ExactArgs(2)),
		Short:		"Command \"-de\" declares an exchange",
		Long:		"Command \"-de\" requires two arguments, name and type, and declares an exchange " +
			"with said attributes.",
		Example:	"-de test direct",
	}

	cf.DeclareQueue = cobra.Command{
		Use:		"-dq",
		Version: 	"1.0",
		Hidden:		false,
		Args:		cobra.PositionalArgs(cobra.ExactArgs(1)),
		Short:		"Command \"-dq\" declares a queue with the argument's name",
		Long:		"Command \"-dq\" requires one argument and declares a queue with the argument's name",
		Example:	"-dq test",
	}

	cf.SetRoutingKey = cobra.Command{
		Use:		"-srk",
		Version:	"1.0",
		Hidden:		false,
		Args:		cobra.PositionalArgs(cobra.ExactArgs(1)),
		Short:		"Command \"-srk\" sets the routing key",
		Long:		"Command \"-srk\" requires one argument and sets the routing key to that argument",
		Example:	"-srk hello.world",
	}

	cf.QueueBind = cobra.Command{
		Use:		"-qb",
		Version: 	"1.0",
		Hidden:		false,
		Args:		cobra.PositionalArgs(cobra.ExactArgs(0)),
		Short:		"Command \"-qb\" binds queue to the current routing keys",
		Long:		"Command \"-qb\" binds queue to the current routing keys",
		Example:	"-qb",
	}

	cf.Register = cobra.Command{
		Use:		"-reg",
		Version: 	"1.0",
		Hidden:		false,
		Args:		cobra.PositionalArgs(cobra.ExactArgs(0)),
		Short:		"Command \"-reg\" register a consumer",
		Long:		"Command \"-reg\" register a consumer and is read to listen to any incoming message",
		Example:	"-reg",
	}
}