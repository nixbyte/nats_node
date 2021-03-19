package request

import (
	"bytes"
	"encoding/base64"
	"nats_node/http/client"
	"nats_node/utils/logger"
	"net/url"
	"strconv"
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

	servers = []string{"nats://192.168.49.91:4111", "nats://192.168.49.92:4222", "nats://127.0.0.1:4111"}

	Nc, err = nats.Connect(strings.Join(servers, ","), nats.NoEcho())
	if err != nil {
		logger.Logger.PrintError(err)
	}

	NatsConnection, err = nats.NewEncodedConn(Nc, nats.JSON_ENCODER)
	if err != nil {
		logger.Logger.PrintError(err)
	}
}

func GetAppState() {
	sub, err := NatsConnection.Conn.SubscribeSync("GetAppState")
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
			request.Endpoint = "/mfc/GetAppState"
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

func GetBranches() {

	sub, err := NatsConnection.Conn.SubscribeSync("GetMfcList")
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
			request.Endpoint = "/mfc/GetMfcList"

			response, err := client.SendRequest(request)
			if err != nil {
				logger.Logger.PrintError(err)
			}
			err = msg.Respond(response)
		}
	}
	defer NatsConnection.Close()
}

func GetBranchServices() {
	sub, err := NatsConnection.Conn.SubscribeSync("GetMfcServices")
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
			request.Endpoint = "/mfc/GetMfcServices"
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
func GetDates() {
	sub, err := NatsConnection.Conn.SubscribeSync("GetDates")
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
			request.Endpoint = "/mfc/GetDates"
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

func GetTimes() {
	sub, err := NatsConnection.Conn.SubscribeSync("GetTimes")
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
			request.Endpoint = "/mfc/GetTimes"
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

func ReserveTime() {
	sub, err := NatsConnection.Conn.SubscribeSync("ReserveTime")
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
			request.Endpoint = "/mfc/ReserveTime"

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
func TimeConfirmation() {
	sub, err := NatsConnection.Conn.SubscribeSync("ConfirmeTime")
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
			request.Endpoint = "/mfc/TimeConfirmation"

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
func GetReservationCode() {
	sub, err := NatsConnection.Conn.SubscribeSync("ReservationCode")
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
			request.Endpoint = "/mfc/GetReservationCode"
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
