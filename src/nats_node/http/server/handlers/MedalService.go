package handlers

import (
	"errors"
	"nats_node/http/model"
	"time"

	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

var SwaggerHandler fasthttp.RequestHandler = fasthttpadaptor.NewFastHTTPHandler(
	httpSwagger.Handler(
		httpSwagger.URL("/medal"),
	))

//var SwaggerYamlHandler fasthttp.RequestHandler = func(ctx *fasthttp.RequestCtx) {
//	ctx.SetContentType("application/json")
//
//	ctx.Response.Header.Set("Access-Control-Allow-Origin", "*")
//
//	bodyBytes, err := ioutil.ReadFile("../docs/swagger/swagger.yaml")
//	logger.Logger.PrintError(err)
//
//	fmt.Fprint(ctx, string(bodyBytes))
//}

var GetTotalPersonsCountHandler fasthttp.RequestHandler = func(ctx *fasthttp.RequestCtx) {
	defer CatchPanic(ctx)

	count := &model.PersonsCountResponse{}

	if ctx.IsPost() == true {
		err = errors.New("method POST not supported")
	} else {
		//exist, err = validateParameters(ctx, []string{"applicationId"})

		//if exist {
		err = NatsConnection.Request("GetTotalPersonsCount", ctx.QueryArgs().QueryString(), count, 10*time.Minute)
		//}
	}
	sendModelIfExist(ctx, count, err)
}

var GetPersonsCountByNameHandler fasthttp.RequestHandler = func(ctx *fasthttp.RequestCtx) {
	defer CatchPanic(ctx)

	count := &model.PersonsCountResponse{}

	if ctx.IsPost() == true {
		err = errors.New("method POST not supported")
	} else {
		exist, err = validateParameters(ctx, []string{"name"})

		if exist {
			err = NatsConnection.Request("GetPersonsCountByName", ctx.QueryArgs().QueryString(), count, 10*time.Minute)
		}
	}
	sendModelIfExist(ctx, count, err)
}

var SearchPersonHandler fasthttp.RequestHandler = func(ctx *fasthttp.RequestCtx) {
	defer CatchPanic(ctx)

	personsList := &model.PersonsListResponse{}

	if ctx.IsPost() == true {
		err = errors.New("method POST not supported")
	} else {
		exist, err = validateParameters(ctx, []string{"size", "page"})

		if exist {
			err = NatsConnection.Request("SearchPerson", ctx.QueryArgs().QueryString(), personsList, 10*time.Minute)
		}
	}
	sendModelIfExist(ctx, personsList, err)
}
var GetAllStoryHandler fasthttp.RequestHandler = func(ctx *fasthttp.RequestCtx) {
	defer CatchPanic(ctx)

	stories := &model.AllStoryResponse{}

	if ctx.IsPost() == true {
		err = errors.New("method POST not supported")
	} else {
		err = NatsConnection.Request("GetAllStory", ctx.QueryArgs().QueryString(), stories, 10*time.Minute)
	}
	sendModelIfExist(ctx, stories, err)
}

var AddWidgetHandler fasthttp.RequestHandler = func(ctx *fasthttp.RequestCtx) {
	defer CatchPanic(ctx)

	response := &model.AddWidgetResponse{}

	if ctx.IsGet() == true {
		err = errors.New("method GET not supported")
	} else {
		err = NatsConnection.Request("AddWidget", ctx.Request.Body(), response, 10*time.Minute)
	}
	sendModelIfExist(ctx, response, err)
}

var NotificationUnsubscribeHandler fasthttp.RequestHandler = func(ctx *fasthttp.RequestCtx) {
	defer CatchPanic(ctx)

	state := &model.ApiResponse{}

	if ctx.IsGet() == true {
		err = errors.New("method GET not supported")
	} else {
		err = NatsConnection.Request("NotificationUnsubscribe", ctx.Request.Body(), state, 10*time.Minute)
	}
	sendModelIfExist(ctx, state.Model, err)
}

var NotificationAddHandler fasthttp.RequestHandler = func(ctx *fasthttp.RequestCtx) {
	defer CatchPanic(ctx)

	state := &model.NotificationAddResponse{}

	if ctx.IsGet() == true {
		err = errors.New("method GET not supported")
	} else {
		err = NatsConnection.Request("NotificationAdd", ctx.Request.Body(), state, 10*time.Minute)
	}
	sendModelIfExist(ctx, state, err)
}

var PostAddHandler fasthttp.RequestHandler = func(ctx *fasthttp.RequestCtx) {
	defer CatchPanic(ctx)

	state := &model.ApiResponse{}

	if ctx.IsGet() == true {
		err = errors.New("method GET not supported")
	} else {
		err = NatsConnection.Request("AddWidget", ctx.Request.Body(), state, 10*time.Minute)
	}
	sendModelIfExist(ctx, state.Model, err)
}
