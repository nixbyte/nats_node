package handlers

import (
	"encoding/xml"
	"errors"
	"fmt"
	soapmodel "nats_node/http/model/soap"
	"nats_node/utils/logger"
	"time"

	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

var OurSpbSwaggerHandler fasthttp.RequestHandler = fasthttpadaptor.NewFastHTTPHandler(
	httpSwagger.Handler(
		httpSwagger.URL("/ourspb.json"),
	))

var GetAllProblemsHandler fasthttp.RequestHandler = func(ctx *fasthttp.RequestCtx) {
	defer CatchPanic(ctx)

	var state string
	var envelope soapmodel.GetAllProblemsEnvelopeResponse

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

	var state string
	var envelope soapmodel.GetProblemEnvelopeResponse

	if ctx.IsPost() == true {
		err = errors.New("method POST not supported")
	} else {
		exist, err = validateParameters(ctx, []string{"problemId"})

		if exist {
			err = NatsConnection.Request("GetProblem", ctx.QueryArgs().QueryString(), &state, 10*time.Minute)

			ctx.Response.Header.Set("Content-Type", "text/xml; charset=utf-8")

			err = xml.Unmarshal([]byte(state), &envelope)
			if err != nil {
				logger.Logger.PrintError(err)
			}

			fmt.Println(ctx, state)
			if envelope.Body.Fault.Faultstring != "" {
				sendModelIfExist(ctx, envelope.Body.Fault, errors.New("Fault Code: "+envelope.Body.Fault.Faultcode+" Fault String: "+envelope.Body.Fault.Faultstring))
			} else {
				sendModelIfExist(ctx, envelope.Body.GetProblemResponse.MessageData.AppData.GetProblemResult, err)
			}
		}
	}
}

var GetFileHandler fasthttp.RequestHandler = func(ctx *fasthttp.RequestCtx) {
	defer CatchPanic(ctx)

	var state string
	var envelope soapmodel.GetFileEnvelopeResponse

	if ctx.IsGet() == true {
		err = errors.New("method GET not supported")
	} else {
		err = NatsConnection.Request("GetFile", ctx.Request.Body(), &state, 10*time.Minute)
	}

	ctx.Response.Header.Set("Content-Type", "text/xml; charset=utf-8")

	err = xml.Unmarshal([]byte(state), &envelope)
	if err != nil {
		logger.Logger.PrintError(err)
	}

	//fmt.Fprint(ctx, state)
	if envelope.Body.Fault.Faultstring != "" {
		sendModelIfExist(ctx, envelope.Body.Fault, errors.New("Fault Code: "+envelope.Body.Fault.Faultcode+" Fault String: "+envelope.Body.Fault.Faultstring))
	} else if envelope.Body.GetFileResponse.MessageData.AppData.Error.ErrorMessage != "" {
		sendModelIfExist(ctx, envelope.Body.GetFileResponse.MessageData.AppData.Error, errors.New("Error Code: "+envelope.Body.GetFileResponse.MessageData.AppData.Error.ErrorCode+" Fault String: "+envelope.Body.GetFileResponse.MessageData.AppData.Error.ErrorMessage))
	} else {
		sendModelIfExist(ctx, envelope.Body.GetFileResponse.MessageData.AppData.GetFileResult, err)
	}
}
