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

	go request.GetAppState()
	go request.GetBranches()
	go request.GetBranchServices()
	go request.GetDates()
	go request.GetTimes()
	go request.ReserveTime()
	go request.TimeConfirmation()
	go request.GetReservationCode()

	//  go request.GetAllProblems()
	//  go request.GetProblem()
	//  go request.GetFile()

	go request.GetCovidToken()
	go request.CheckCovidTokenExpiration()
	go request.GetTest()
	go request.PostTest()
	go request.CheckCovidTokenExpiration()
	go request.GetDistrictList()
	go request.GetDistricts()
	go request.GetLpuList()
	go request.GetCovidLpuList()
	go request.GetSpecialityList()
	go request.GetDoctorList()
	go request.GetAppointmentList()
	go request.CheckPatient()
	go request.AddPatient()
	go request.UpdatePhone()
	go request.SetAppointment()
	go request.DeleteAppointment()
	go request.GetCovidLpuNames()
	go request.GetCovidLpuIds()
	go request.GetCovidLpuIdByName()
	go request.GetCovidSpecialityNames()
	go request.GetCovidSpecialityIds()
	go request.GetCovidSpecialityId()
	go request.GetCovidDocNames()
	go request.GetCovidDocIds()
	go request.GetCovidDocId()
	go request.GetCovidAppointmentTimes()
	go request.GetCovidAppointmentIds()
	go request.GetCovidAppointmentCount()
	select {}
}
