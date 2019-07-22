package pkg

import (
	"database/sql"
	_"github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
	"strings"
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

	//Configures yaml file
		viper.SetConfigType("yaml")
		viper.AddConfigPath("..")
		err := viper.ReadInConfig()
		checkError(err)


		db, err := sql.Open("sqlite3", "../amqp")
		checkError(err)

		defer db.Close()

		var rabbit = RabbitMQ{}
		checkError(rabbit.Dial(viper.GetString("port")))
		checkError(rabbit.OpenChannel())
		checkError(rabbit.DeclareExchange(viper.GetString("exchange.name"),viper.GetString("exchange.type")))
		checkError(rabbit.DeclareQueue(viper.GetString("queue_name")))


		//Read configuration of consumers
		consumerMap := viper.Get("consumer").([]interface{})
		for _, v := range consumerMap {
			eachConsumerMap := v.(map[interface{}]interface{})
			for k, v := range eachConsumerMap {
				consumer := Consumer{}
				consumerName := k.(string)
				consumerRoutingKeys := strings.Fields(v.(string))
				registerError := rabbit.Register(consumer.handle, consumerName, consumerRoutingKeys)
				checkError(registerError)
			}
		}

		//debugging variable
		var message = "example"

		tx, err := db.Begin()
		checkError(err)

		_, err = tx.Exec("INSERT INTO amqp(message) VALUES(?)", message)
		if err != nil {
			tx.Rollback()
			log.Fatal(err)
		}
		checkError(tx.Commit())

		forever := make(chan bool)

		log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
		<- forever
	},
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}