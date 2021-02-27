package handlers

import (
	"encoding/json"
	"fmt"
	"nats_node/http/client/model"
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
