package handlers

import (
	"fmt"

	"nats_node/http/client/request"
	"nats_node/utils/logger"

	"github.com/valyala/fasthttp"
)

var TokenHandler fasthttp.RequestHandler = func(ctx *fasthttp.RequestCtx) {
	defer CatchPanic(ctx)
	token, err := request.GetToken()
	logger.Logger.PrintError(err)
	bytes, err := getBytes(token)
	fmt.Println(string(bytes))
}
