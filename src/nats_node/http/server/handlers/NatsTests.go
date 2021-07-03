package handlers

import (
	"errors"
	"fmt"
	"time"

	"github.com/valyala/fasthttp"
)

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
