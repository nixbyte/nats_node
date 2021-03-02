package request

import (
	"nats_node/http/client"
	nats_client "nats_node/nats"
	"nats_node/utils/logger"
	"time"
)

func GetMfcList() {

	request := client.NewRequest()

	request.Rt = client.GET
	request.Endpoint = "/calendar-backend/public/api/v1/branches"

	response, err := client.SendRequest(request)

	if err != nil {
		logger.Logger.PrintError(err)
	}

	sub, err := nats_client.NatsConnection.Conn.SubscribeSync("token")
	if err != nil {
		logger.Logger.PrintError(err)
	}

	// Wait for a message
	msg, err := sub.NextMsg(10 * time.Minute)
	if err != nil {
		logger.Logger.PrintError(err)
	}
	defer nats_client.NatsConnection.Close()
	msg.Respond(response)
}
