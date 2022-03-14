package request

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
	context "nats_node/nats/model"
	"nats_node/utils/logger"
	"strconv"
)

func GetBytesFromNatsBase64Msg(natsMsgData []byte) ([]byte, error) {

	unquoteString, err := strconv.Unquote(string(natsMsgData))
	if err != nil {
		logger.Logger.PrintError(err)
	}

	queryBytes, err := base64.StdEncoding.DecodeString(unquoteString)
	if err != nil {
		logger.Logger.PrintError(err)
	}

	return queryBytes, err
}

func GetRequestContextFromBytesArray(b []byte) (ctx *context.RequestContext, err error) {

	c := new(context.RequestContext)

	network := bytes.NewBuffer(b)
	bytesDecoder := gob.NewDecoder(network)
	err = bytesDecoder.Decode(&c)
	return c, err
}
