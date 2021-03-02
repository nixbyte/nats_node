package nats_client

import (
	nats "github.com/nats-io/nats.go"
)

var NatsConnection *nats.EncodedConn

func init() {
	nc, _ := nats.Connect("127.0.0.1:4222")
	NatsConnection, _ = nats.NewEncodedConn(nc, nats.JSON_ENCODER)
}
