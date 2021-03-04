package request

import (
	"nats_node/http/client"
	"nats_node/utils/logger"
	"time"

	"github.com/nats-io/nats.go"
)

func GetToken() {
	//	nats_sender.Configure()
	//	nats_sender.Connect()
	//	nats_sender.Subscribe()
	request := client.NewRequest()

	request.Rt = client.GET
	request.Endpoint = "/health/GetToken"

	response, err := client.SendRequest(request)
	if err != nil {
		logger.Logger.PrintError(err)
	}
	//servers := []string{"nats://192.168.49.91:4222", "nats://192.168.49.92:4222"}
	//nc, err := nats.Connect(strings.Join(servers, ","), nats.NoEcho())
	nc, err := nats.Connect("localhost:4222", nats.Name("api"), nats.NoEcho())
	if err != nil {
		logger.Logger.PrintError(err)
	}
	defer nc.Close()

	NatsConnection, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	if err != nil {
		logger.Logger.PrintError(err)
	}

	sub, err := NatsConnection.Conn.SubscribeSync("token")
	if err != nil {
		logger.Logger.PrintError(err)
	}

	for {
		// Wait for a message
		msg, err := sub.NextMsg(10 * time.Minute)
		if err != nil {
			logger.Logger.PrintError(err)
		}
		msg.Respond(response)
		break
	}
	sub.Unsubscribe()
	defer NatsConnection.Close()
}
