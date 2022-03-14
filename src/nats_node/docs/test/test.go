// Package classification Test Service.
//
// API для сервиса - NatsTest
//
//     Schemes: http
//     BasePath: /
//     Version: 1.0.0
//     Host: localhost:8080
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Security:
//     - http
//
//    SecurityDefinitions:
//    basic:
//      type: basic
//
// swagger:meta
//package statistic
//
//import model "nats_node/http/model/json"
//
//type ApiResponseStatistic struct {
//	Success bool   `json:"success"`
//	Message string `json:"message"`
//	Model   string `json:"model"`
//}
//
//// swagger:route POST /send idOfSendEndpoint
//// send - метод для отправки статистики
//// responses:
////   200: ApiResponseStatistic
//
//// swagger:parameters  idOfSendEndpoint
//type StatisticRequest struct {
//	// unique: true
//	// required: true
//	// in: body
//	Body model.StatisticModel
//}
//
//// Модель данных возвращает временную метку загруженных в ClickHouse данных
//// swagger:response ApiResponseStatistic
//type StatisticResponseWrapper struct {
//	// in:body
//	Body ApiResponseStatistic
//}
