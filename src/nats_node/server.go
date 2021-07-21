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

	server.ApiServer.GetRouter().GET("/mfc/swagger/{filename*}", handlers.MfcSwaggerHandler)
	server.ApiServer.GetRouter().GET("/medal/swagger/{filename*}", handlers.MedalSwaggerHandler)
	server.ApiServer.GetRouter().GET("/ourspb/swagger/{filename*}", handlers.OurSpbSwaggerHandler)
	server.ApiServer.GetRouter().GET("/health/swagger/{filename*}", handlers.CovidSwaggerHandler)
	server.ApiServer.GetRouter().GET("/statistic/swagger/{filename*}", handlers.StatisticSwaggerHandler)

	//	server.ApiServer.GetRouter().GET("/mfc/GetAppState", handlers.GetAppStateHandler)
	//	server.ApiServer.GetRouter().GET("/mfc/GetMfcList", handlers.BranchesHandler)
	//	server.ApiServer.GetRouter().GET("/mfc/GetMfcServices", handlers.BranchServiceHandler)
	//	server.ApiServer.GetRouter().GET("/mfc/GetDates", handlers.DatesHandler)
	//	server.ApiServer.GetRouter().GET("/mfc/GetTimes", handlers.TimesHandler)
	//	server.ApiServer.GetRouter().POST("/mfc/ReserveTime", handlers.ReservationHandler)
	//	server.ApiServer.GetRouter().POST("/mfc/TimeConfirmation", handlers.TimeConfirmationHandler)
	//	server.ApiServer.GetRouter().GET("/mfc/GetReservationCode", handlers.ReservationCodeHandler)
	//
	//  server.ApiServer.GetRouter().GET("/medal/GetPersonsCount", handlers.GetTotalPersonsCountHandler)
	//  server.ApiServer.GetRouter().GET("/medal/GetPersonsCountByName", handlers.GetPersonsCountByNameHandler)
	//	server.ApiServer.GetRouter().GET("/medal/SearchPerson", handlers.SearchPersonHandler)
	//	server.ApiServer.GetRouter().GET("/medal/GetAllStory", handlers.GetAllStoryHandler)

	//  server.ApiServer.GetRouter().POST("/ourspb/GetAllProblems", handlers.GetAllProblemsHandler)
	//  server.ApiServer.GetRouter().GET("/ourspb/GetProblem", handlers.GetProblemHandler)
	//  server.ApiServer.GetRouter().POST("/ourspb/GetFile", handlers.GetFileHandler)

	//  server.ApiServer.GetRouter().GET("/health/GetToken", handlers.TokenHandler)
	//  server.ApiServer.GetRouter().GET("/health/CheckToken", handlers.TokenExpirationHandler)
	//  server.ApiServer.GetRouter().GET("/health/GetDistrictList", handlers.DistrictListHandler)
	//  server.ApiServer.GetRouter().GET("/health/GetDistricts", handlers.DistrictsHandler)
	//  server.ApiServer.GetRouter().GET("/health/GetCovidLpuList", handlers.CovidLpuListHandler)
	//  server.ApiServer.GetRouter().GET("/health/GetSpecialityList", handlers.SpecialityListHandler)
	//  server.ApiServer.GetRouter().GET("/health/GetDoctorList", handlers.DoctorListHandler)
	//  server.ApiServer.GetRouter().GET("/health/GetAppointmentList", handlers.AvailableAppointmentHandler)
	//  server.ApiServer.GetRouter().POST("/health/CheckPatient", handlers.CheckPatientHandler)
	//  server.ApiServer.GetRouter().POST("/health/AddPatient", handlers.AddPatientHandler)
	//  server.ApiServer.GetRouter().POST("/health/UpdatePhone", handlers.UpdatePhoneHandler)
	//  server.ApiServer.GetRouter().POST("/health/SetAppointment", handlers.SetAppointmentHandler)
	//  server.ApiServer.GetRouter().POST("/health/DeleteAppointment", handlers.DeleteAppointmentHandler)
	//  server.ApiServer.GetRouter().GET("/health/GetLpuList", handlers.LpuListHandler)
	//  server.ApiServer.GetRouter().GET("/health/GetCovidLpuNames", handlers.CovidLpuNamesHandler)
	//  server.ApiServer.GetRouter().GET("/health/GetCovidLpuIds", handlers.CovidLpuIdsHandler)
	//  server.ApiServer.GetRouter().GET("/health/GetCovidLpuIdByName", handlers.CovidLpuIdByNameHandler)
	//  server.ApiServer.GetRouter().GET("/health/GetSpecialityNames", handlers.CovidSpecialityNamesHandler)
	//  server.ApiServer.GetRouter().GET("/health/GetSpecialityIds", handlers.CovidSpecialityIdsHandler)
	//  server.ApiServer.GetRouter().GET("/health/GetSpecialityId", handlers.CovidSpecialityIdHandler)
	//  server.ApiServer.GetRouter().GET("/health/GetDocNames", handlers.CovidDocNamesHandler)
	//  server.ApiServer.GetRouter().GET("/health/GetDocIds", handlers.CovidDocIdsHandler)
	//  server.ApiServer.GetRouter().GET("/health/GetDocId", handlers.CovidDocIdHandler)
	//  server.ApiServer.GetRouter().GET("/health/GetAppointmentTimes", handlers.CovidAppointmentTimesHandler)
	//  server.ApiServer.GetRouter().GET("/health/GetAppointmentIds", handlers.CovidAppointmentIdsHandler)
	//  server.ApiServer.GetRouter().GET("/health/GetAppointmentCount", handlers.CovidAppointmentCountHandler)

	server.ApiServer.GetRouter().POST("/statistic/send", handlers.StatisticSendHandler)

	if monitoring.Monitoring.WRITE_METRICS {
		server.MetricServer.GetRouter().GET("/state", fasthttpadaptor.NewFastHTTPHandler(promhttp.Handler()))
		server.MetricServer.GetRouter().GET("/metrics", fasthttpadaptor.NewFastHTTPHandler(promhttp.HandlerFor(monitoring.Registry, promhttp.HandlerOpts{})))
	}

	select {}
}
