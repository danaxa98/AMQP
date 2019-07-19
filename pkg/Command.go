package pkg

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
		viper.SetConfigFile("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("/loopline/")

		if err := viper.ReadInConfig(); err != nil { // Handle errors reading the config file
			if _, ok := err.(viper.ConfigFileNotFoundError); ok {
				err := RabbitError(NoConfigFileError)
				panic(err)
			} else {
				err := RabbitError(ConfigError)
				panic(err)
			}

		}

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