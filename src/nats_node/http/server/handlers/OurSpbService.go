package handlers

import (
	"encoding/xml"
	"errors"
	jsonmodel "nats_node/http/model/json"
	soapmodel "nats_node/http/model/soap"
	"nats_node/utils/logger"
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

	var state string
	var envelope soapmodel.EnvelopeResponse

	if ctx.IsGet() == true {
		err = errors.New("method GET not supported")
	} else {
		err = NatsConnection.Request("GetAllProblems", ctx.Request.Body(), &state, 10*time.Minute)
	}

	ctx.Response.Header.Set("Content-Type", "text/xml; charset=utf-8")

	err = xml.Unmarshal([]byte(state), &envelope)
	if err != nil {
		logger.Logger.PrintError(err)
	}

	if envelope.Body.Fault.Faultstring != "" {
		sendModelIfExist(ctx, envelope.Body.Fault, errors.New("Fault Code: "+envelope.Body.Fault.Faultcode+" Fault String: "+envelope.Body.Fault.Faultstring))
	} else {
		sendModelIfExist(ctx, envelope.Body.GetProblemsListResponse.MessageData.AppData.GetProblemsListResult, err)
	}
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
