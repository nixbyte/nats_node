package nats_client

import (
	"fmt"

	nats "github.com/nats-io/nats.go"
)

var NatsConnection *nats.Conn

func init() {
	NatsConnection, _ = nats.Connect("127.0.0.1:4222")
	fmt.Println("Test")
}
