package main

import (
	request "nats_node/http/client/requests"
	"nats_node/utils/monitoring"
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
	//server.Start()

	//if monitoring.Monitoring.WRITE_METRICS {
	//	server.MetricServer.AddHandlerToRoute("/state", fasthttpadaptor.NewFastHTTPHandler(promhttp.Handler()))
	//	server.MetricServer.AddHandlerToRoute("/metrics", fasthttpadaptor.NewFastHTTPHandler(promhttp.HandlerFor(monitoring.Registry, promhttp.HandlerOpts{})))
	//}

	//  go request.GetTotalPersonsCount()
	//  go request.GetPersonsCountByName()
	//  go request.SearchPerson()
	//  go request.GetAllStory()
	//  go request.AddWidget()
	//  go request.NotificationUnsubscribe()
	//  go request.NotificationAdd()
	//  go request.PostAdd()

	//  go request.GetAppState()
	//	go request.GetBranches()
	//	go request.GetBranchServices()
	//	go request.GetDates()
	//	go request.GetTimes()
	//	go request.ReserveTime()
	//	go request.TimeConfirmation()
	//	go request.GetReservationCode()

	go request.GetAllProblems()
	go request.GetProblem()
	go request.GetFile()

	select {}
}
