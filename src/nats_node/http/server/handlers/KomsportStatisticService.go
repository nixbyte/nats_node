package handlers

import (
	"errors"
	model "nats_node/http/model/json"
	context "nats_node/nats/model"
	"nats_node/utils/logger"
	"time"

	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

var StatisticSwaggerHandler fasthttp.RequestHandler = fasthttpadaptor.NewFastHTTPHandler(
	httpSwagger.Handler(
		httpSwagger.URL("/statistic.json"),
	))

var StatisticSendHandler fasthttp.RequestHandler = func(ctx *fasthttp.RequestCtx) {
	defer CatchPanic(ctx)

	var token model.Token

	if ctx.IsGet() == true {
		err = errors.New("method GET not supported")
	} else {

		rc := new(context.RequestContext)
		rc.New(ctx)

		bytes, err := requestContextToBytesArray(rc)
		if err != nil {
			logger.Logger.PrintError(err)
		}
		err = NatsConnection.Request("KomsportStatistic", bytes, &token, 10*time.Minute)
		if err != nil {
			logger.Logger.PrintError(err)
		}

	}
	sendModelIfExist(ctx, token, err)
}
