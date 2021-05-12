package request

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	configs "nats_node/configs/http"
	"nats_node/http/client"
	jsonmodel "nats_node/http/model/json"
	"nats_node/utils/logger"
	"regexp"
	"time"
)

func GetAllMessages() {
	sub, err := NatsConnection.Conn.SubscribeSync("GetAllProblems")
	if err != nil {
		logger.Logger.PrintError(err)
	}

	for {
		// Wait for a message
		msg, err := sub.NextMsg(10 * time.Minute)
		if err != nil {
			logger.Logger.PrintError(err)
		} else {

			requestBody, err := GetBytesFromNatsBase64Msg(msg.Data)
			if err != nil {
				logger.Logger.PrintError(err)
			}

			problemsRequest := &jsonmodel.GetAllProblemsRequest{}
			err = json.Unmarshal(requestBody, problemsRequest)
			if err != nil {
				logger.Logger.PrintError(err)
			}

			problemsEnvelopeRequest := problemsRequest.GetSoapEnvelopeRequest()

			signServiceConfig := &configs.ClientConfig{
				"http://localhost:8181",
				60,
				60,
				65500,
				65500,
				20,
				60,
				120,
			}

			signService := client.NewWorker(signServiceConfig)

			envelopeBody, err := xml.Marshal(problemsEnvelopeRequest)
			if err != nil {
				logger.Logger.PrintError(err)
			}

			request := client.NewRequest()
			request.Rt = client.POST
			request.Headers["Content-Type"] = "application/xml"
			request.Endpoint = "/api/smev2/sign/wss"
			request.Body = bytes.NewReader(envelopeBody)

			response, err := signService.SendRequest(request)
			if err != nil {
				logger.Logger.PrintError(err)
			}

			problemsEnvelopeResponseString := string(response)
			var tokenString = "<soapenv:Header><gorod:ApplicationToken>JcYJslfFh2kQw9XbMsSFds3EQFMUy7miqXf4LSdYKFgCwdWFfUQszRxgH</gorod:ApplicationToken>"

			reg := regexp.MustCompile(`<soapenv:Header>`)
			envelopeParts := reg.Split(problemsEnvelopeResponseString, -1)
			envelopeWithToken := envelopeParts[0] + tokenString + envelopeParts[1]

			//problemsEnvelopeResponse.Header.ApplicationToken = "JcYJslfFh2kQw9XbMsSFds3EQFMUy7miqXf4LSdYKFgCwdWFfUQszRxgH"

			//request = client.NewRequest()
			//request.Rt = client.POST
			//request.Endpoint = "/smev/openspb/"
			//request.Headers["Content-Type"] = "text/xml charset=utf-8"
			//request.Body = bytes.NewReader(response)

			//apiResponse, err := client.Client.SendRequest(request)

			//responseBody, err := json.Marshal(apiResponse)
			//if err != nil {
			//	logger.Logger.PrintError(err)
			//}
			err = msg.Respond([]byte(envelopeWithToken))
		}
	}
	defer NatsConnection.Close()
}
