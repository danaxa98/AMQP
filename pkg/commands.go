
package pkg

import (
	"github.com/spf13/cobra"
)

func main() {
	SetRoutingKey := cobra.Command{
		Use:		"-srk",
		Version:	"1.0",
		Hidden:		false,
		Args:		cobra.PositionalArgs(cobra.ExactArgs(1)),
		Short:		"Command \"-srk\" sets the routing key",
		Long:		"Command \"-srk\" requires one argument and sets the routing key to that argument",
		Example:	"\"-srk hello.world\"",
	}

	DeclareExchange := cobra.Command{
		Use:		"-de",
		Version:	"1.0",
		Hidden:		false,
		Args:		cobra.PositionalArgs(cobra.ExactArgs(2)),
		Short:		"Command \"-de\" declares an exchange",
		Long:		"Command \"-de\" requires two arguments, name and type, and declares an exchange " +
			"with said attributes.",
		Example:	"-de test direct",
	}

	DeclareQueue := cobra.Command{
		Use:		"-dq",
		Version: 	"1.0",
		Hidden:		false,
		Args:		cobra.PositionalArgs(cobra.ExactArgs(1)),
		Short:		"Command \"-dq\" requires one argument and declares a queue with the argument's name",
		Long:		"Command \"-dq\" requires one argument and declares a queue with the argument's name",
		Example:	"\"-dq test\"",
	}

	QueueBind := cobra.Command{
		Use:		"-qb",
		Version: 	"1.0",
		Hidden:		false,
		Args:		cobra.PositionalArgs(cobra.ExactArgs(0)),
		Short:		"Command \"-qb\" binds queue to the current routing keys",
		Long:		"Command \"-qb\" binds queue to the current routing keys",
		Example:	"-qb",
	}
}