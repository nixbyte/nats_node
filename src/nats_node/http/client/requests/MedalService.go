package request

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"nats_node/http/client"
	"nats_node/http/model"
	"nats_node/utils/logger"
	"net/url"
	"strconv"
	"time"
)

func GetTotalPersonsCount() {
	sub, err := NatsConnection.Conn.SubscribeSync("GetTotalPersonsCount")
	if err != nil {
		logger.Logger.PrintError(err)
	}

	for {
		// Wait for a message
		msg, err := sub.NextMsg(10 * time.Second)
		if err != nil {
			logger.Logger.PrintError(err)
		} else {

			request := client.NewRequest()
			request.Rt = client.GET
			request.Endpoint = "/api/v1/person/total/count"

			response, err := client.SendRequest(request)
			if err != nil {
				logger.Logger.PrintError(err)
			}
			err = msg.Respond(response)
		}
	}
	defer NatsConnection.Close()
}

func GetPersonsCountByName() {
	sub, err := NatsConnection.Conn.SubscribeSync("GetPersonsCountByName")
	if err != nil {
		logger.Logger.PrintError(err)
	}

	for {
		// Wait for a message
		msg, err := sub.NextMsg(10 * time.Second)
		if err != nil {
			logger.Logger.PrintError(err)
		} else {

			unquoteString, err := strconv.Unquote(string(msg.Data))
			if err != nil {
				logger.Logger.PrintError(err)
			}

			queryBytes, err := base64.StdEncoding.DecodeString(unquoteString)
			if err != nil {
				logger.Logger.PrintError(err)
			}

			request := client.NewRequest()

			values, err := url.ParseQuery(string(queryBytes))

			request.Rt = client.GET
			request.Endpoint = "/api/v1/person/count/" + values.Get("name")

			response, err := client.SendRequest(request)
			if err != nil {
				logger.Logger.PrintError(err)
			}
			err = msg.Respond(response)
		}
	}
	defer NatsConnection.Close()
}

func SearchPerson() {
	sub, err := NatsConnection.Conn.SubscribeSync("SearchPerson")
	if err != nil {
		logger.Logger.PrintError(err)
	}

	for {
		// Wait for a message
		msg, err := sub.NextMsg(10 * time.Second)
		if err != nil {
			logger.Logger.PrintError(err)
		} else {

			unquoteString, err := strconv.Unquote(string(msg.Data))
			if err != nil {
				logger.Logger.PrintError(err)
			}

			queryBytes, err := base64.StdEncoding.DecodeString(unquoteString)
			if err != nil {
				logger.Logger.PrintError(err)
			}

			request := client.NewRequest()

			values, err := url.ParseQuery(string(queryBytes))

			request.Rt = client.GET
			request.Endpoint = "/api/v1/person/search/"
			request.Parameters = values

			response, err := client.SendRequest(request)
			if err != nil {
				logger.Logger.PrintError(err)
			}
			err = msg.Respond(response)
		}
	}
	defer NatsConnection.Close()
}

func GetAllStory() {
	sub, err := NatsConnection.Conn.SubscribeSync("GetAllStory")
	if err != nil {
		logger.Logger.PrintError(err)
	}

	for {
		// Wait for a message
		msg, err := sub.NextMsg(10 * time.Second)
		if err != nil {
			logger.Logger.PrintError(err)
		} else {
			request := client.NewRequest()

			request.Rt = client.GET
			request.Endpoint = "/api/v1/mini-story/get/all"

			response, err := client.SendRequest(request)
			if err != nil {
				logger.Logger.PrintError(err)
			}
			err = msg.Respond(response)
		}
	}
	defer NatsConnection.Close()
}

func AddWidget() {
	sub, err := NatsConnection.Conn.SubscribeSync("AddWidget")
	if err != nil {
		logger.Logger.PrintError(err)
	}

	for {
		// Wait for a message
		msg, err := sub.NextMsg(10 * time.Second)
		if err != nil {
			logger.Logger.PrintError(err)
		} else {

			request := client.NewRequest()

			request.Rt = client.POST
			request.Endpoint = "/api/v1/security/community/widget/add"

			unquoteString, err := strconv.Unquote(string(msg.Data))
			if err != nil {
				logger.Logger.PrintError(err)
			}

			queryBytes, err := base64.StdEncoding.DecodeString(unquoteString)
			if err != nil {
				logger.Logger.PrintError(err)
			}
			request.Body = bytes.NewReader(queryBytes)

			response, err := client.SendRequest(request)
			if err != nil {
				logger.Logger.PrintError(err)
			}
			err = msg.Respond(response)
		}
	}
	defer NatsConnection.Close()
}

func NotificationUnsubscribe() {
	sub, err := NatsConnection.Conn.SubscribeSync("NotificationUnsubscribe")
	if err != nil {
		logger.Logger.PrintError(err)
	}

	for {
		// Wait for a message
		msg, err := sub.NextMsg(10 * time.Second)
		if err != nil {
			logger.Logger.PrintError(err)
		} else {

			request := client.NewRequest()

			request.Rt = client.POST

			unquoteString, err := strconv.Unquote(string(msg.Data))
			if err != nil {
				logger.Logger.PrintError(err)
			}

			queryBytes, err := base64.StdEncoding.DecodeString(unquoteString)
			if err != nil {
				logger.Logger.PrintError(err)
			}

			var unsubscribeRequest model.NotificationUnsubscribeRequest
			if err := json.Unmarshal(queryBytes, &unsubscribeRequest); err != nil {
				panic(err)
			}

			request.Endpoint = "/api/v1/security/notification/" + strconv.Itoa(unsubscribeRequest.Id) + "/unsubscribe"
			request.Parameters.Add("token", unsubscribeRequest.Token)

			response, err := client.SendRequest(request)
			if err != nil {
				logger.Logger.PrintError(err)
			}
			err = msg.Respond(response)
		}
	}
	defer NatsConnection.Close()
}

func NotificationAdd() {
	sub, err := NatsConnection.Conn.SubscribeSync("NotificationAdd")
	if err != nil {
		logger.Logger.PrintError(err)
	}

	for {
		// Wait for a message
		msg, err := sub.NextMsg(10 * time.Second)
		if err != nil {
			logger.Logger.PrintError(err)
		} else {

			request := client.NewRequest()

			request.Rt = client.POST
			request.Endpoint = "/api/v1/security/notification/add"

			unquoteString, err := strconv.Unquote(string(msg.Data))
			if err != nil {
				logger.Logger.PrintError(err)
			}

			queryBytes, err := base64.StdEncoding.DecodeString(unquoteString)
			if err != nil {
				logger.Logger.PrintError(err)
			}
			request.Body = bytes.NewReader(queryBytes)

			response, err := client.SendRequest(request)
			if err != nil {
				logger.Logger.PrintError(err)
			}
			err = msg.Respond(response)
		}
	}
	defer NatsConnection.Close()
}

func PostAdd() {
	sub, err := NatsConnection.Conn.SubscribeSync("PostAdd")
	if err != nil {
		logger.Logger.PrintError(err)
	}

	for {
		// Wait for a message
		msg, err := sub.NextMsg(10 * time.Second)
		if err != nil {
			logger.Logger.PrintError(err)
		} else {

			request := client.NewRequest()

			request.Rt = client.POST
			request.Endpoint = "/api/v1/security/vk/post/add"

			unquoteString, err := strconv.Unquote(string(msg.Data))
			if err != nil {
				logger.Logger.PrintError(err)
			}

			queryBytes, err := base64.StdEncoding.DecodeString(unquoteString)
			if err != nil {
				logger.Logger.PrintError(err)
			}
			request.Body = bytes.NewReader(queryBytes)

			response, err := client.SendRequest(request)
			if err != nil {
				logger.Logger.PrintError(err)
			}
			err = msg.Respond(response)
		}
	}
	defer NatsConnection.Close()
}
