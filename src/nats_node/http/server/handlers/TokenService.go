package handlers

import (
	"log"
	"nats_node/http/model"
	"nats_node/utils/logger"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/valyala/fasthttp"
)

var TokenHandler fasthttp.RequestHandler = func(ctx *fasthttp.RequestCtx) {
	defer CatchPanic(ctx)

	m := &model.Token{}

	//servers := []string{"nats://192.168.49.91:4222", "nats://192.168.49.92:4222"}
	//nc, err := nats.Connect(strings.Join(servers, ","), nats.NoEcho())
	nc, err := nats.Connect("localhost:4222", nats.NoEcho())
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	NatsConnection, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	if err != nil {
		log.Fatal(err)
	}

	err = NatsConnection.Request("token", ctx.Request.Body(), m, 10*time.Minute)
	if err != nil {
		logger.Logger.PrintError(err)
	}
	getResult(ctx, m)
}
