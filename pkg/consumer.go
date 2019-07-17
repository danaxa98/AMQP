package pkg

import "fmt"

type Consumer struct {}

func (consumer Consumer) handle (msg []byte) {
	fmt.Printf("Message \"%s\" is received.\n", string(msg))
}