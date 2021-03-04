package nats_client

import (
	"log"

	"github.com/nats-io/nats.go"
)

var NatsConnection *nats.EncodedConn

func init() {

	nc, err := nats.Connect("localhost:4222", nats.NoEcho())
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	NatsConnection, err = nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	if err != nil {
		log.Fatal(err)
	}
}
