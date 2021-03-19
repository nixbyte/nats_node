package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"nats_node/http/model"
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
		resp := model.ApiResponse{
			Success: false,
			Message: err.Error(),
			Model:   nil,
		}
		getResult(ctx, resp)
	} else {
		resp := model.ApiResponse{
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
