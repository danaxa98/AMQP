package main

import (
	"Loopline/pkg"
	"bufio"
	"os"
)

func main() {
	cf := pkg.CommandFactory{}
	cf.Config()

	consumer := pkg.Consumer{}
	consumer.UseRabbitMQ(pkg.RabbitMQ{})
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

}