package pkg

import (
	_"database/sql"
	_"github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
	"strings"
	_"time"
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
	//
	//db, err := sql.Open("sqlite3", "./foo.db")
	//CheckError(err)
	//defer db.Close()
	//
	//tx, err := db.Begin()
}

var Listen = &cobra.Command{
	Use:		"listen",
	Short:		"Listen for message",
	Args:		cobra.PositionalArgs(cobra.ExactArgs(0)),
	Run:		func(cmd *cobra.Command, args[] string) {
	//Configures yaml file
		viper.SetConfigType("yaml")
		viper.AddConfigPath("..")
		err := viper.ReadInConfig()
		CheckError(err)

		var rabbit = RabbitMQ{}
		CheckError(rabbit.Dial(viper.GetString("port")))
		CheckError(rabbit.OpenChannel())
		CheckError(rabbit.DeclareExchange(viper.GetString("exchange.name"),viper.GetString("exchange.type")))
		CheckError(rabbit.DeclareQueue(viper.GetString("queue_name")))


		//Read configuration of consumers
		consumerMap := viper.Get("consumer").([]interface{})
		for _, v := range consumerMap {
			eachConsumerMap := v.(map[interface{}]interface{})
			for k, v := range eachConsumerMap {
				consumer := Consumer{}
				consumerName := k.(string)
				consumerRoutingKeys := strings.Fields(v.(string))
				rabbit.Register(consumer.handle, consumerName, consumerRoutingKeys)
			}
		}

		forever := make(chan bool)

		log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
		<- forever
	},
}

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}