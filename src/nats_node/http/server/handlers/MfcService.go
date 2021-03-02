package handlers

import (
	"nats_node/http/model"
	nats_client "nats_node/nats"
	"nats_node/utils/logger"
	"time"

	"github.com/valyala/fasthttp"
)

var MfcListHandler fasthttp.RequestHandler = func(ctx *fasthttp.RequestCtx) {
	defer CatchPanic(ctx)

	m := &model.MfcList{}

	err := nats_client.NatsConnection.Request("token", ctx.Request.Body(), m, 10*time.Minute)
	if err != nil {
		logger.Logger.PrintError(err)
	}
	getResult(ctx, m)
}
