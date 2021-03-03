package request

import (
	"nats_node/http/client"
	nats_client "nats_node/nats"
	"nats_node/utils/logger"
)

func GetMfcList() {
	nats_client.Configure()
	nats_client.Connect()
	nats_client.Subscribe()

	request := client.NewRequest()

	request.Rt = client.GET
	request.Endpoint = "/calendar-backend/public/api/v1/branches"

	response, err := client.SendRequest(request)
	if err != nil {
		logger.Logger.PrintError(err)
	}

	nats_client.SendResponse(response)

	defer nats_client.Disconnect()
}
