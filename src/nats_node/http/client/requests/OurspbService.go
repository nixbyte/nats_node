package request

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	configs "nats_node/configs/http"
	"nats_node/http/client"
	jsonmodel "nats_node/http/model/json"
	soapmodel "nats_node/http/model/soap"
	"nats_node/utils/logger"
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

			problemsEnvelope := &soapmodel.GetAllProblemsEnvelope{}
			problemsEnvelope.Soapenv = "http://schemas.xmlsoap.org/soap/envelope/"
			problemsEnvelope.Gorod = "https://gorod.gov.spb.ru/smev/gorod"
			problemsEnvelope.Rev = "http://smev.gosuslugi.ru/rev120315"
			problemsList := soapmodel.GetProblemsList{}
			problemsList.Message = soapmodel.Message{}
			problemsList.Message.Sender = soapmodel.Sender{
				"",
				"SPB010000",
				"Система классификаторов",
			}
			problemsEnvelope.Body = soapmodel.GetProblemsListBody{
				"",
				"http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-utility-1.0.xsd",
				problemsList,
			}

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

			envelopeBody, err := xml.MarshalIndent(problemsEnvelope, "", " ")
			if err != nil {
				logger.Logger.PrintError(err)
			}

			request := client.NewRequest()
			request.Rt = client.POST
			request.Endpoint = "/api/smev2/sign/wss"
			request.Body = bytes.NewReader(envelopeBody)

			response, err := signService.SendRequest(request)
			if err != nil {
				logger.Logger.PrintError(err)
			}

			problemsEnvelopeResponse := &soapmodel.GetAllProblemsEnvelopeResponse{}

			err = xml.Unmarshal([]byte(string(response)), problemsEnvelopeResponse)
			if err != nil {
				logger.Logger.PrintError(err)
			}

			problemsEnvelopeResponse.Header.ApplicationToken = "JcYJslfFh2kQw9XbMsSFds3EQFMUy7miqXf4LSdYKFgCwdWFfUQszRxgH"

			request = client.NewRequest()
			request.Rt = client.POST
			request.Endpoint = "/smev/openspb/"
			request.Headers["Content-Type"] = "text/xml charset=utf-8"
			request.Body = bytes.NewReader(response)

			apiResponse, err := client.Client.SendRequest(request)

			responseBody, err := json.Marshal(apiResponse)
			if err != nil {
				logger.Logger.PrintError(err)
			}

			err = msg.Respond(responseBody)
		}
	}
	defer NatsConnection.Close()
}
