package handlers

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"errors"
	"fmt"
	jsonmodel "nats_node/http/model/json"
	model "nats_node/http/model/json"
	context "nats_node/nats/model"
	"nats_node/utils/logger"
	"runtime/debug"

	"github.com/valyala/fasthttp"
)

func CatchPanic(ctx *fasthttp.RequestCtx) {
	if r := recover(); r != nil {
		logger.Logger.PrintRecover(r)
		logger.Logger.PrintStack(debug.Stack())
		ctx.Error("Internal Error", fasthttp.StatusInternalServerError)
	}
}

func getBytes(obj interface{}) ([]byte, error) {
	bodyBytes, err := json.Marshal(obj)
	if err != nil {
		logger.Logger.PrintError(err)
		return nil, err
	}
	return bodyBytes, nil
}

func getResult(ctx *fasthttp.RequestCtx, obj interface{}) {
	ctx.SetContentType("application/json")

	bodyBytes, err := getBytes(obj)
	logger.Logger.PrintError(err)

	fmt.Fprint(ctx, string(bodyBytes))
}

func sendModelIfExist(ctx *fasthttp.RequestCtx, m interface{}, err error) {
	if err != nil {
		resp := jsonmodel.ApiResponse{
			Success: false,
			Message: err.Error(),
			Model:   nil,
		}
		getResult(ctx, resp)
	} else {
		resp := jsonmodel.ApiResponse{
			Success: true,
			Message: "",
			Model:   m,
		}
		getResult(ctx, resp)
	}
}
func checkIfExistParameter(ctx *fasthttp.RequestCtx, name string) (bool, error) {
	parameter := ctx.QueryArgs().Peek(name)
	if len(parameter) == 0 {
		return false, errors.New("Query parameter " + string(parameter) + " not found")
	} else {
		return true, nil
	}
}

func checkHeader(ctx *fasthttp.RequestCtx, h string) (bool, error) {

	header := ctx.Request.Header.Peek(h)

	if len(header) == 0 {
		return false, errors.New(h + " Header not found or empty")
	} else {
		return true, nil
	}
}
func validateHeaders(ctx *fasthttp.RequestCtx, h []string) (bool, error) {
	var buffer bytes.Buffer
	for _, header := range h {
		exist, _ := checkHeader(ctx, header)
		if exist != true {
			buffer.WriteString(" " + header + " ")
		}
	}

	if len(buffer.Bytes()) != 0 {
		return false, errors.New("Headers: " + buffer.String() + " not found")
	} else {
		return true, nil
	}
}

func validateParameters(ctx *fasthttp.RequestCtx, params []string) (bool, error) {
	var buffer bytes.Buffer
	for _, param := range params {
		exist, _ := checkIfExistParameter(ctx, param)
		if exist != true {
			buffer.WriteString(" " + param + " ")
		}
	}

	if len(buffer.Bytes()) != 0 {
		return false, errors.New("Query parameters" + buffer.String() + " not found")
	} else {
		return true, nil
	}
}

func requestContextToBytesArray(context *context.RequestContext) ([]byte, error) {
	var bytesBuffer bytes.Buffer
	bytesEncoder := gob.NewEncoder(&bytesBuffer)
	err := bytesEncoder.Encode(&context)
	return bytesBuffer.Bytes(), err
}

func validatePatient(patient model.Patient) (bool, error) {
	isValid := true

	if len(patient.Firstname) == 0 ||
		len(patient.Lastname) == 0 ||
		len(patient.Middlename) == 0 ||
		len(patient.Phone) == 0 ||
		patient.Birthday == 0 {

		isValid = false
	}
	if isValid == true {
		return isValid, nil
	} else {
		return false, errors.New("Person is not valid or empty")
	}
}
