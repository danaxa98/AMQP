package pkg

import "fmt"

type Consumer struct {name string}

func (consumer Consumer) handle (msg []byte) string {
	fmt.Printf("Message \"%s\" is received.\n", string(msg))
	return string(msg)
}