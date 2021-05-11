package handlers

import (
	"errors"
	jsonmodel "nats_node/http/model/json"
	"time"

	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

var OurSpbSwaggerHandler fasthttp.RequestHandler = fasthttpadaptor.NewFastHTTPHandler(
	httpSwagger.Handler(
		httpSwagger.URL("/ourspb"),
	))

var GetAllProblemsHandler fasthttp.RequestHandler = func(ctx *fasthttp.RequestCtx) {
	defer CatchPanic(ctx)

	state := &jsonmodel.ApiResponse{}

	if ctx.IsGet() == true {
		err = errors.New("method GET not supported")
	} else {
		err = NatsConnection.Request("GetAllProblems", ctx.Request.Body(), state, 10*time.Minute)
	}
	sendModelIfExist(ctx, state.Model, err)
}

var GetProblemHandler fasthttp.RequestHandler = func(ctx *fasthttp.RequestCtx) {
	defer CatchPanic(ctx)

	state := &jsonmodel.ApiResponse{}

	if ctx.IsPost() == true {
		err = errors.New("method POST not supported")
	} else {
		exist, err = validateParameters(ctx, []string{"problemId"})

		if exist {
			err = NatsConnection.Request("GetProblem", ctx.QueryArgs().QueryString(), state, 10*time.Minute)
		}
	}
	sendModelIfExist(ctx, state.Model, err)
}
