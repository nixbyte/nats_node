package main

import (
	"nats_node/http/server"
	"nats_node/http/server/handlers"
	"nats_node/utils/monitoring"

	"github.com/prometheus/client_golang/prometheus/promhttp"
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
	server.Api.AddHandlerToRoute("/medal/GetPersonsCount", handlers.GetTotalPersonsCountHandler)
	server.Api.AddHandlerToRoute("/medal/GetPersonsCountByName", handlers.GetPersonsCountByNameHandler)
	server.Api.AddHandlerToRoute("/medal/SearchPerson", handlers.SearchPersonHandler)
	server.Api.AddHandlerToRoute("/medal/GetAllStory", handlers.GetAllStoryHandler)
	server.Api.AddHandlerToRoute("/medal/widget/add", handlers.AddWidgetHandler)
	server.Api.AddHandlerToRoute("/medal/notification/unsubscribe", handlers.NotificationUnsubscribeHandler)
	server.Api.AddHandlerToRoute("/medal/notification/add", handlers.NotificationAddHandler)
	server.Api.AddHandlerToRoute("/medal/post/add", handlers.PostAddHandler)

	if monitoring.Monitoring.WRITE_METRICS {
		server.MetricApi.AddHandlerToRoute("/state", fasthttpadaptor.NewFastHTTPHandler(promhttp.Handler()))
		server.MetricApi.AddHandlerToRoute("/metrics", fasthttpadaptor.NewFastHTTPHandler(promhttp.HandlerFor(monitoring.Registry, promhttp.HandlerOpts{})))
	}

	select {}

}
