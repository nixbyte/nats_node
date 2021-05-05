package main

import (
	"nats_node/http/server"
	"nats_node/http/server/handlers"
	"nats_node/utils/monitoring"

	"github.com/fasthttp/router"
	docs "github.com/kecci/fasthttp-swagger/docs"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

func init() {
	monitoring.Monitoring.StartMonitoring()
	go func() {
		<-monitoring.Monitoring.StopChan
		//@todo: will be made graceful stop here
		monitoring.Monitoring.ExitChan <- monitoring.ExitCodeInterrupted
	}()
}

func main() {

	server.Start()

	//	server.Api.AddHandlerToRoute("/mfc/GetAppState", handlers.GetAppStateHandler)
	//	server.Api.AddHandlerToRoute("/mfc/GetMfcList", handlers.BranchesHandler)
	//	server.Api.AddHandlerToRoute("/mfc/GetMfcServices", handlers.BranchServiceHandler)
	//	server.Api.AddHandlerToRoute("/mfc/GetDates", handlers.DatesHandler)
	//	server.Api.AddHandlerToRoute("/mfc/GetTimes", handlers.TimesHandler)
	//	server.Api.AddHandlerToRoute("/mfc/ReserveTime", handlers.ReservationHandler)
	//	server.Api.AddHandlerToRoute("/mfc/TimeConfirmation", handlers.TimeConfirmationHandler)
	//	server.Api.AddHandlerToRoute("/mfc/GetReservationCode", handlers.ReservationCodeHandler)
	server.ApiServer.AddHandlerToRoute("/medal/GetPersonsCount", handlers.GetTotalPersonsCountHandler)
	server.ApiServer.AddHandlerToRoute("/medal/GetPersonsCountByName", handlers.GetPersonsCountByNameHandler)
	server.ApiServer.AddHandlerToRoute("/medal/SearchPerson", handlers.SearchPersonHandler)
	server.ApiServer.AddHandlerToRoute("/medal/GetAllStory", handlers.GetAllStoryHandler)
	//server.Api.AddHandlerToRoute("/medal/widget/add", handlers.AddWidgetHandler)
	//server.Api.AddHandlerToRoute("/medal/notification/unsubscribe", handlers.NotificationUnsubscribeHandler)
	//server.Api.AddHandlerToRoute("/medal/notification/add", handlers.NotificationAddHandler)
	//server.Api.AddHandlerToRoute("/medal/post/add", handlers.PostAddHandler)
	server.SwaggerServer.AddHandlerToRoute("/docs/swagger.yaml", handlers.SwaggerYamlHandler)
	//server.ApiServer.AddHandlerToRoute("/swagger/index.html", handlers.SwaggerHandler)
	//	server.ApiServer.AddHandlerToRoute("/swagger/swagger-ui.css", handlers.FilesHandler)
	//	server.ApiServer.AddHandlerToRoute("/swagger/swagger-ui-bundle.js", handlers.FilesHandler)
	//	server.ApiServer.AddHandlerToRoute("/swagger/swagger-ui-standalone-preset.js", handlers.FilesHandler)

	if monitoring.Monitoring.WRITE_METRICS {
		server.MetricServer.AddHandlerToRoute("/state", fasthttpadaptor.NewFastHTTPHandler(promhttp.Handler()))
		server.MetricServer.AddHandlerToRoute("/metrics", fasthttpadaptor.NewFastHTTPHandler(promhttp.HandlerFor(monitoring.Registry, promhttp.HandlerOpts{})))
	}

	docs.SwaggerInfo.Title = "Fasthttp Swagger"
	docs.SwaggerInfo.Description = "Fasthttp Swagger"
	docs.SwaggerInfo.Version = "1.0.0-0"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	docs.SwaggerInfo.Host = ""
	docs.SwaggerInfo.BasePath = ""

	r := router.New()
	r.GET("/docs/{filepath:*}", fasthttpadaptor.NewFastHTTPHandlerFunc(httpSwagger.WrapHandler))
	r.GET("/hello", Hello())

	if err := fasthttp.ListenAndServe(":9191", r.Handler); err != nil {
		panic(err)
	}
	select {}

}

// Hello godoc
// @Summary Show a Hello
// @Description get hello
// @Tags hello
// @Accept json
// @Produce json
// @Success 200 {object} string
// @Router /hello [get]
func Hello() fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.SetStatusCode(200)
		ctx.SetBody([]byte(`{ "message" : "HELLO" }`))
	}
}
