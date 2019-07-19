package pkg

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
)


var RootCommand = &cobra.Command{
	Version:	"1.0",
}

func Execute() {
	if err := RootCommand.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func init() {
	RootCommand.AddCommand(Listen)
}

var Listen = &cobra.Command{
	Use:		"listen",
	Short:		"Listen for message",
	Args:		cobra.PositionalArgs(cobra.ExactArgs(0)),
	Run:		func(cmd *cobra.Command, args[] string) {
		viper.SetConfigType("yaml")
		viper.AddConfigPath("..")
		if err := viper.ReadInConfig(); err != nil { // Handle errors reading the config file
			fmt.Println(err)
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
		if err := rabbit.Dial(viper.GetString("port")); err != Default {
			panic(err)
		}
		if err := rabbit.OpenChannel(); err != Default {
			panic(err)
		}
		if err := rabbit.DeclareExchange(viper.GetString("exchange.name"),viper.GetString("exchange.type")); err != Default {
			panic(err)
		}
		if err := rabbit.DeclareQueue(viper.GetString("queue_name")); err != Default {
			panic(err)
		}
		rabbit.Register(viper.GetString("consumer.routing_key"), viper.GetString("consumer.name"), consumer.handle)
	},
}