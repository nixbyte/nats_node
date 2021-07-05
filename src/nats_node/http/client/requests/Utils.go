package request

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"errors"
	model "nats_node/http/model/json"
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

func validatePatient(patient model.Patient) (bool, error) {
	isValid := true

	if len(patient.Firstname) == 0 ||
		len(patient.Lastname) == 0 ||
		len(patient.Middlename) == 0 ||
		len(patient.Phone) == 0 ||
		len(patient.IdLpu) == 0 ||
		patient.Birthday == 0 {

		isValid = false
	}
	if isValid == true {
		return isValid, nil
	} else {
		return false, errors.New("Person is not valid or empty")
	}
}
