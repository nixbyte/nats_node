package main

import (
	request "nats_node/http/client/requests"
	"nats_node/http/server"
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

	if monitoring.Monitoring.WRITE_METRICS {
		server.MetricApi.AddHandlerToRoute("/state", fasthttpadaptor.NewFastHTTPHandler(promhttp.Handler()))
		server.MetricApi.AddHandlerToRoute("/metrics", fasthttpadaptor.NewFastHTTPHandler(promhttp.HandlerFor(monitoring.Registry, promhttp.HandlerOpts{})))
	}

	request.GetToken()
}
