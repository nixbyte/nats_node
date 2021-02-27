package handlers

import (
	"nats_node/http/client/request"
	"nats_node/utils/logger"

	"github.com/valyala/fasthttp"
)

var TokenHandler fasthttp.RequestHandler = func(ctx *fasthttp.RequestCtx) {
	defer CatchPanic(ctx)
	tk, err := request.GetToken()
	logger.Logger.PrintError(err)
	getResult(ctx, tk)
}
