package handlers

import (
	"errors"
	jsonmodel "nats_node/http/model/json"
	"nats_node/utils/logger"
	"strings"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/valyala/fasthttp"
)

var servers []string
var Nc *nats.Conn
var NatsConnection *nats.EncodedConn
var err error
var exist bool

func init() {

	servers = []string{"nats://192.168.49.91:4111", "nats://192.168.49.92:4222", "nats://127.0.0.1:4111"}

	Nc, err = nats.Connect(strings.Join(servers, ","), nats.NoEcho())
	if err != nil {
		logger.Logger.PrintError(err)
	}

	NatsConnection, err = nats.NewEncodedConn(Nc, nats.JSON_ENCODER)
	if err != nil {
		logger.Logger.PrintError(err)
	}
}

var GetAppStateHandler fasthttp.RequestHandler = func(ctx *fasthttp.RequestCtx) {
	defer CatchPanic(ctx)

	state := &jsonmodel.ApiResponse{}

	if ctx.IsPost() == true {
		err = errors.New("method POST not supported")
	} else {
		exist, err = validateParameters(ctx, []string{"applicationId"})

		if exist {
			err = NatsConnection.Request("GetAppState", ctx.QueryArgs().QueryString(), state, 10*time.Minute)
		}
	}
	sendModelIfExist(ctx, state.Model, err)
}

var BranchesHandler fasthttp.RequestHandler = func(ctx *fasthttp.RequestCtx) {
	defer CatchPanic(ctx)

	state := &jsonmodel.ApiResponse{}

	if ctx.IsPost() == true {
		err = errors.New("method POST not supported")
	} else {
		err = NatsConnection.Request("GetMfcList", ctx.Request.Body(), state, 10*time.Minute)
	}
	sendModelIfExist(ctx, state.Model, err)
}

var BranchServiceHandler fasthttp.RequestHandler = func(ctx *fasthttp.RequestCtx) {
	defer CatchPanic(ctx)

	state := &jsonmodel.ApiResponse{}

	if ctx.IsPost() == true {
		err = errors.New("method POST not supported")
	} else {
		exist, err = validateParameters(ctx, []string{"branch"})

		if exist {
			err = NatsConnection.Request("GetMfcServices", ctx.QueryArgs().QueryString(), state, 10*time.Minute)
		}
	}
	sendModelIfExist(ctx, state.Model, err)
}

var DatesHandler fasthttp.RequestHandler = func(ctx *fasthttp.RequestCtx) {
	defer CatchPanic(ctx)

	state := &jsonmodel.ApiResponse{}

	if ctx.IsPost() == true {
		err = errors.New("method POST not supported")
	} else {
		exist, err = validateParameters(ctx, []string{"branch", "serviceId"})

		if exist {
			err = NatsConnection.Request("GetDates", ctx.QueryArgs().QueryString(), state, 10*time.Minute)
		}
	}
	sendModelIfExist(ctx, state.Model, err)
}

var TimesHandler fasthttp.RequestHandler = func(ctx *fasthttp.RequestCtx) {
	defer CatchPanic(ctx)

	state := &jsonmodel.ApiResponse{}

	if ctx.IsPost() == true {
		err = errors.New("method POST not supported")
	} else {
		exist, err = validateParameters(ctx, []string{"branch", "serviceId", "date"})

		if exist {
			err = NatsConnection.Request("GetTimes", ctx.QueryArgs().QueryString(), state, 10*time.Minute)
		}
	}
	sendModelIfExist(ctx, state.Model, err)
}

var ReservationHandler fasthttp.RequestHandler = func(ctx *fasthttp.RequestCtx) {
	defer CatchPanic(ctx)

	state := &jsonmodel.ApiResponse{}

	if ctx.IsGet() == true {
		err = errors.New("method GET not supported")
	} else {
		err = NatsConnection.Request("ReserveTime", ctx.Request.Body(), state, 10*time.Minute)
	}
	sendModelIfExist(ctx, state.Model, err)
}

var TimeConfirmationHandler fasthttp.RequestHandler = func(ctx *fasthttp.RequestCtx) {
	defer CatchPanic(ctx)

	state := &jsonmodel.ApiResponse{}

	if ctx.IsGet() == true {
		err = errors.New("method GET not supported")
	} else {
		err = NatsConnection.Request("ConfirmeTime", ctx.Request.Body(), state, 10*time.Minute)
	}
	sendModelIfExist(ctx, state.Model, err)
}

var ReservationCodeHandler fasthttp.RequestHandler = func(ctx *fasthttp.RequestCtx) {
	defer CatchPanic(ctx)

	state := &jsonmodel.ApiResponse{}

	if ctx.IsPost() == true {
		err = errors.New("method POST not supported")
	} else {
		exist, err = validateParameters(ctx, []string{"publicId"})

		if exist {
			err = NatsConnection.Request("ReservationCode", ctx.QueryArgs().QueryString(), state, 10*time.Minute)
		}
	}
	sendModelIfExist(ctx, state.Model, err)
}
