package context

import (
	"bytes"

	"github.com/valyala/fasthttp"
)

type RequestContext struct {
	Request   []byte
	Body      []byte
	Headers   map[string]string
	QueryArgs map[string]string
}

func (rc *RequestContext) New(ctx *fasthttp.RequestCtx) {
	var requestBuffer bytes.Buffer
	ctx.Request.WriteTo(&requestBuffer)
	rc.Request = requestBuffer.Bytes()

	var bodyBuffer bytes.Buffer
	ctx.Request.BodyWriteTo(&bodyBuffer)
	rc.Body = bodyBuffer.Bytes()

	rc.Headers = make(map[string]string)
	rc.QueryArgs = make(map[string]string)

	ctx.Request.Header.VisitAll(func(key, value []byte) {
		rc.Headers[string(key)] = string(value)
	})

	ctx.QueryArgs().VisitAll(func(key, value []byte) {
		rc.QueryArgs[string(key)] = string(value)
	})
}
