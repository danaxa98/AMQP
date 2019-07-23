package pkg

import "fmt"

type Consumer struct {
	name 		string
	rep			repository
}

func (consumer Consumer) handle (msg []byte) {
	consumer.store(msg)
}

func (consumer Consumer) store (msg [] byte) {
	consumer.rep.insert(msg)
	fmt.Printf("Message \"%s\" is stored in the database.\n", string(msg))
}