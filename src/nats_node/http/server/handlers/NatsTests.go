package handlers

import (
	"errors"
	"fmt"
	"nats_node/utils/logger"
	"strings"
	"time"

	"github.com/nats-io/nats.go"
	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

var servers []string
var Nc *nats.Conn
var NatsConnection *nats.EncodedConn
var err error
var exist bool

func init() {

	servers = []string{"nats://localhost:4111", "nats://localhost:4222", "nats://localhost:5222"}

	Nc, err = nats.Connect(strings.Join(servers, ","), nats.NoEcho())
	if err != nil {
		logger.Logger.PrintError(err)
	}

	NatsConnection, err = nats.NewEncodedConn(Nc, nats.JSON_ENCODER)
	if err != nil {
		logger.Logger.PrintError(err)
	}
}

var TestSwaggerHandler fasthttp.RequestHandler = fasthttpadaptor.NewFastHTTPHandler(
	httpSwagger.Handler(
		httpSwagger.URL("/test.json"),
	))

var GetTestHandler fasthttp.RequestHandler = func(ctx *fasthttp.RequestCtx) {
	defer CatchPanic(ctx)

	var state string

	if ctx.IsPost() == true {
		err = errors.New("method POST not supported")
	} else {
		err = NatsConnection.Request("GetTest", ctx.QueryArgs().String(), &state, 10*time.Minute)
	}

	fmt.Fprint(ctx, state)
}

var PostTestHandler fasthttp.RequestHandler = func(ctx *fasthttp.RequestCtx) {
	defer CatchPanic(ctx)

	var state string

	if ctx.IsGet() == true {
		err = errors.New("method GET not supported")
	} else {
		err = NatsConnection.Request("PostTest", ctx.Request.Body(), &state, 10*time.Minute)
	}

	fmt.Fprint(ctx, state)
}
