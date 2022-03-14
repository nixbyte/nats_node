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

	if monitoring.Monitoring.WRITE_METRICS {
		server.MetricServer.GetRouter().GET("/state", fasthttpadaptor.NewFastHTTPHandler(promhttp.Handler()))
		server.MetricServer.GetRouter().GET("/metrics", fasthttpadaptor.NewFastHTTPHandler(promhttp.HandlerFor(monitoring.Registry, promhttp.HandlerOpts{})))
	}

	select {}
}
