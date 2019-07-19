package pkg

import "fmt"

type Consumer struct {name string}

func (consumer Consumer) handle (msg []byte) {
	fmt.Printf("Message \"%s\" is received.\n", string(msg))
}