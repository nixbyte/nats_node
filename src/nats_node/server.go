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

	server.ApiServer.GetRouter().ServeFiles("/{filepath:*}", "./docs/swagger/")
	server.ApiServer.GetRouter().GET("/GetTest", handlers.GetTestHandler)
	server.ApiServer.GetRouter().POST("/PostTest", handlers.PostTestHandler)

	server.ApiServer.GetRouter().GET("/mfc/GetAppState", handlers.GetAppStateHandler)
	server.ApiServer.GetRouter().GET("/mfc/GetMfcList", handlers.BranchesHandler)
	server.ApiServer.GetRouter().GET("/mfc/GetMfcServices", handlers.BranchServiceHandler)
	server.ApiServer.GetRouter().GET("/mfc/GetDates", handlers.DatesHandler)
	server.ApiServer.GetRouter().GET("/mfc/GetTimes", handlers.TimesHandler)
	server.ApiServer.GetRouter().POST("/mfc/ReserveTime", handlers.ReservationHandler)
	server.ApiServer.GetRouter().POST("/mfc/TimeConfirmation", handlers.TimeConfirmationHandler)
	server.ApiServer.GetRouter().POST("/mfc/GetReservationCode", handlers.ReservationCodeHandler)

	//  server.ApiServer.GetRouter().GET("/medal/GetPersonsCount", handlers.GetTotalPersonsCountHandler)
	//  server.ApiServer.GetRouter().GET("/medal/GetPersonsCountByName", handlers.GetPersonsCountByNameHandler)
	//	server.ApiServer.GetRouter().GET("/medal/SearchPerson", handlers.SearchPersonHandler)
	//	server.ApiServer.GetRouter().GET("/medal/GetAllStory", handlers.GetAllStoryHandler)
	//	server.ApiServer.GetRouter().GET("/medal/swagger/{filename*}", handlers.MedalSwaggerHandler)

	//  server.ApiServer.GetRouter().GET("/ourspb/swagger/{filename*}", handlers.OurSpbSwaggerHandler)
	//  server.ApiServer.GetRouter().POST("/ourspb/GetAllProblems", handlers.GetAllProblemsHandler)
	//  server.ApiServer.GetRouter().GET("/ourspb/GetProblem", handlers.GetProblemHandler)
	//  server.ApiServer.GetRouter().POST("/ourspb/GetFile", handlers.GetFileHandler)

	if monitoring.Monitoring.WRITE_METRICS {
		server.MetricServer.GetRouter().GET("/state", fasthttpadaptor.NewFastHTTPHandler(promhttp.Handler()))
		server.MetricServer.GetRouter().GET("/metrics", fasthttpadaptor.NewFastHTTPHandler(promhttp.HandlerFor(monitoring.Registry, promhttp.HandlerOpts{})))
	}

	select {}
}
