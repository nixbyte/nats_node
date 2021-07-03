package request

import (
	"nats_node/utils/logger"
	"time"
)

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
