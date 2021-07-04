package request

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	configs "nats_node/configs/http"
	"nats_node/http/client"
	jsonmodel "nats_node/http/model/json"
	"nats_node/utils/logger"
	"net/url"
	"regexp"
	"time"
)

func sendSignRequest(env jsonmodel.SoapEnvelope) ([]byte, error) {

	envelopeRequest := env.GetSoapEnvelopeRequest()

	signServiceConfig := &configs.ClientConfig{
		"http://localhost:8181",
		"",
		60,
		60,
		65500,
		65500,
		20,
		60,
		120,
	}

	signService := client.NewWorker(signServiceConfig)

	envelopeBody, err := xml.Marshal(envelopeRequest)
	if err != nil {
		logger.Logger.PrintError(err)
	}

	request := client.NewRequest()
	request.Rt = client.POST
	request.Headers["Content-Type"] = "application/xml"
	request.Headers["Connection"] = "close"
	request.Endpoint = "/api/smev2/sign/wss"
	request.Body = bytes.NewReader(envelopeBody)

	return signService.SendRequest(request)
}

func addTokenToEnvelope(response []byte) []byte {

	problemsEnvelopeResponseString := string(response)

	var tokenString = "<soapenv:Header><gorod:ApplicationToken>JcYJslfFh2kQw9XbMsSFds3EQFMUy7miqXf4LSdYKFgCwdWFfUQszRxgH</gorod:ApplicationToken>" //TODO remove govnokod разобраться с маршалингом в xml

	reg := regexp.MustCompile(`<soapenv:Header>`)
	envelopeParts := reg.Split(problemsEnvelopeResponseString, -1)
	envelopeWithToken := envelopeParts[0] + tokenString + envelopeParts[1]

	return []byte(envelopeWithToken)
}

func sendRequestToOurspb(body *[]byte) ([]byte, error) {
	request := client.NewRequest()
	request.Rt = client.POST
	request.Endpoint = "/smev/openspb/"
	request.Headers["Content-Type"] = "text/xml charset=utf-8"
	request.Body = bytes.NewReader(*body)

	return client.Client.SendRequest(request)
}

func GetAllProblems() {
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

			response, err := sendSignRequest(problemsRequest)
			envelopeWithToken := addTokenToEnvelope(response)

			apiResponse, err := sendRequestToOurspb(&envelopeWithToken)
			if err != nil {
				logger.Logger.PrintError(err)
			}

			err = msg.Respond(apiResponse)
		}
	}
	defer NatsConnection.Close()
}

func GetProblem() {
	sub, err := NatsConnection.Conn.SubscribeSync("GetProblem")
	if err != nil {
		logger.Logger.PrintError(err)
	}

	for {
		// Wait for a message
		msg, err := sub.NextMsg(10 * time.Minute)
		if err != nil {
			logger.Logger.PrintError(err)
		} else {

			query, err := GetBytesFromNatsBase64Msg(msg.Data)
			if err != nil {
				logger.Logger.PrintError(err)
			}

			parameters, err := url.ParseQuery(string(query))
			if err != nil {
				logger.Logger.PrintError(err)
			}

			problemRequest := &jsonmodel.GetProblemRequest{
				ProblemId: parameters.Get("problemId"),
			}

			response, err := sendSignRequest(problemRequest)
			if err != nil {
				logger.Logger.PrintError(err)
			}

			envelopeWithToken := addTokenToEnvelope(response)

			apiResponse, err := sendRequestToOurspb(&envelopeWithToken)
			if err != nil {
				logger.Logger.PrintError(err)
			}

			err = msg.Respond(apiResponse)
		}
	}
	defer NatsConnection.Close()
}

func GetFile() {
	sub, err := NatsConnection.Conn.SubscribeSync("GetFile")
	if err != nil {
		logger.Logger.PrintError(err)
	}

	for {
		// Wait for a message
		msg, err := sub.NextMsg(10 * time.Minute)

		if err != nil {
			logger.Logger.PrintError(err)
		} else {

			fmt.Println("Hello")
			requestBody, err := GetBytesFromNatsBase64Msg(msg.Data)
			if err != nil {
				logger.Logger.PrintError(err)
			}

			filesRequest := &jsonmodel.GetFileRequest{}
			err = json.Unmarshal(requestBody, filesRequest)
			if err != nil {
				logger.Logger.PrintError(err)
			}

			response, err := sendSignRequest(filesRequest)
			if err != nil {
				logger.Logger.PrintError(err)
			}

			envelopeWithToken := addTokenToEnvelope(response)

			apiResponse, err := sendRequestToOurspb(&envelopeWithToken)
			if err != nil {
				logger.Logger.PrintError(err)
			}

			err = msg.Respond(apiResponse)
		}
	}
	defer NatsConnection.Close()
}
