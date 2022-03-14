package request

import (
	"nats_node/utils/logger"
	"strings"
	"time"

	"github.com/nats-io/nats.go"
)

var servers []string
var Nc *nats.Conn
var NatsConnection *nats.EncodedConn
var err error
var exist bool

func init() {

	servers = []string{"nats://127.0.0.1:4111", "nats://127.0.0.1:4222", "nats://127.0.0.1:4333"}

	Nc, err = nats.Connect(strings.Join(servers, ","))
	if err != nil {
		logger.Logger.PrintError(err)
	}

	NatsConnection, err = nats.NewEncodedConn(Nc, nats.JSON_ENCODER)
	if err != nil {
		logger.Logger.PrintError(err)
	}
}

func GetTest() {
	sub, err := NatsConnection.Conn.SubscribeSync("GetTest")
	if err != nil {
		logger.Logger.PrintError(err)
	}

	for {
		// Wait for a message
		msg, err := sub.NextMsg(10 * time.Minute)

		if err != nil {
			logger.Logger.PrintError(err)
		} else {
			err = msg.Respond(msg.Data)
		}
	}
	defer NatsConnection.Close()
}

func PostTest() {
	sub, err := NatsConnection.Conn.SubscribeSync("PostTest")
	if err != nil {
		logger.Logger.PrintError(err)
	}

	for {
		// Wait for a message
		msg, err := sub.NextMsg(10 * time.Minute)

		if err != nil {
			logger.Logger.PrintError(err)
		} else {
			err = msg.Respond(msg.Data)
		}
	}
	defer NatsConnection.Close()
}
