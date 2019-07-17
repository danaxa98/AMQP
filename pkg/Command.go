package pkg

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)


var RootCommand = &cobra.Command{
	Version:	"1.0",
}

func Execute() {
	if err := RootCommand.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	RootCommand.AddCommand(Listen)
}

var Listen = &cobra.Command{
	Use:		"listen",
	Short:		"Listen for message",
	Args:		cobra.PositionalArgs(cobra.ExactArgs(2)),
	Run:		func(cmd *cobra.Command, args[] string) {

		var rabbit = RabbitMQ{}
		consumer := Consumer{}
		if err := rabbit.Dial("amqp://guest:guest@localhost:5672/"); err != Default {
			panic(err)
		}
		if err := rabbit.OpenChannel(); err != Default {
			panic(err)
		}
		if err := rabbit.DeclareExchange(args[0],"topic"); err != Default {
			panic(err)
		}
		if err := rabbit.DeclareQueue(args[1]); err != Default {
			panic(err)
		}
		rabbit.Register("hello.world", "", consumer.handle)
	},
}