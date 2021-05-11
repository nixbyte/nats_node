package request

import (
	"encoding/base64"
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
