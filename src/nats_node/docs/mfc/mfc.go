// Package classification Mfc Service.
//
// API для сервиса - Предварительная запись в МФЦ
//
//     Schemes: http
//     BasePath: /mfc
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
//     - none
//
//    SecurityDefinitions:
//    basic:
//      type: basic
//
// swagger:meta
package mfc

import (
	model "nats_node/http/model/json"
)

type ApiResponseMfcList struct {
	Success bool                  `json:"success"`
	Message string                `json:"message"`
	Model   model.MfcListResponse `json:"model"`
}

// swagger:route GET /GetMfcList idOfGetMfcListEndpoint
// GetMfcList позволяет получить список доступных для записи МФЦ
// responses:
//   200: ApiResponseMfcList

// Модель данных возвращает список МФЦ
// swagger:response ApiResponseMfcList
type GetMfcListResponseWrapper struct {
	// in:body
	Body ApiResponseMfcList
}

type ApiResponseMfcServicesList struct {
	Success bool                      `json:"success"`
	Message string                    `json:"message"`
	Model   model.MfcServicesResponse `json:"model"`
}

// swagger:route GET /GetMfcServices idOfGetMfcServicesEndpoint
// GetMfcServices позволяет получить список сервисов МФЦ
// responses:
//   200: ApiResponseMfcServicesList

// Модель данных возвращает список услуг и сервисов МФЦ
// swagger:response ApiResponseMfcList
type GetMfcServicesResponseWrapper struct {
	// in:body
	Body ApiResponseMfcServicesList
}

// swagger:parameters GetMfcServices idOfGetMfcServicesEndpoint
type BranchParam struct {
	// - publicId полученный из метода GetMfcList
	// unique: true
	// required: true
	// in: query
	// example: "50c339b67c0cb6da05fcd7320bbd049d2551184be1ec934f94bc2891c3380f1c"
	Branch string `json:"branch"`
}

type ApiResponseDatesList struct {
	Success bool                `json:"success"`
	Message string              `json:"message"`
	Model   model.DatesResponse `json:"model"`
}

// swagger:route GET /GetDates idOfGetDatesEndpoint
// GetDates позволяет получить список дат доступных для записи в МФЦ
// responses:
//   200: ApiResponseDatesList

// Модель данных возвращает список дат доступных для ззаписи
// swagger:response ApiResponseDatesList
type GetDatesResponseWrapper struct {
	// in:body
	Body ApiResponseDatesList
}

// swagger:parameters GetDates idOfGetDatesEndpoint
type DatesParam struct {
	// - publicId полученный из метода GetMfcList
	// unique: true
	// required: true
	// in: query
	// example: "50c339b67c0cb6da05fcd7320bbd049d2551184be1ec934f94bc2891c3380f1c"
	Branch string `json:"branch"`
	// - publicId полученный из метода GetMfcServices
	// unique: true
	// required: true
	// in: query
	// example: "049d2551184be1ec934fsfdf34225wef94bc2891c3380f1c"
	ServiceId string `json:"serviceId"`
}

type ApiResponseTimesList struct {
	Success bool                `json:"success"`
	Message string              `json:"message"`
	Model   model.TimesResponse `json:"model"`
}

// swagger:route GET /GetTimes idOfGetTimesEndpoint
// GetTimes позволяет получить список времен доступных для записи в МФЦ
// responses:
//   200: ApiResponseTimesList

// Модель данных возвращает список времен доступных для записи
// swagger:response ApiResponseTimesList
type GetTimesResponseWrapper struct {
	// in:body
	Body ApiResponseTimesList
}

// swagger:parameters GetTimes idOfGetTimesEndpoint
type TimesParam struct {
	// - publicId полученный из метода GetMfcList
	// unique: true
	// required: true
	// in: query
	// example: "50c339b67c0cb6da05fcd7320bbd049d2551184be1ec934f94bc2891c3380f1c"
	Branch string `json:"branch"`
	// - publicId полученный из метода GetMfcServices
	// unique: true
	// required: true
	// in: query
	// example: "049d2551184be1ec934fsfdf34225wef94bc2891c3380f1c"
	ServiceId string `json:"serviceId"`
	// - дата полученный из метода GetDates
	// unique: true
	// required: true
	// in: query
	// example: "2019-07-12"
	Date string `json:"date"`
}

type ApiResponseGetAppState struct {
	Success bool                   `json:"success"`
	Message string                 `json:"message"`
	Model   model.AppStateResponse `json:"model"`
}

// swagger:route GET /GetAppState idOfGetAppStateEndpoint
// GetAppState позволяет получить информацию по заявлению
// responses:
//   200: ApiResponseGetAppState

// Модель данных возвращает информацию по заявлению
// swagger:response ApiResponseGetAppState
type GetAppStateResponseWrapper struct {
	// in:body
	Body ApiResponseGetAppState
}

// swagger:parameters GetAppState idOfGetAppStateEndpoint
type ApplicationIdParam struct {
	// - applicationId - номер заявления
	// unique: true
	// required: true
	// in: query
	// example: "19232323"
	ApplicationId string `json:"applicationId"`
}

type ApiResponseReserveTime struct {
	Success bool                  `json:"success"`
	Message string                `json:"message"`
	Model   model.ReserveResponse `json:"model"`
}

// swagger:route POST /ReserveTime idOfReserveTimeEndpoint
// ReserverTime - метод резервирования времени в МФЦ
// responses:
//   200: ApiResponseReserveTime

// swagger:parameters ReserveTime idOfReserveTimeEndpoint
type ReserveRequest struct {
	// unique: true
	// required: true
	// in: body
	Body model.ReservationRequestData
}

// Модель данных возвращает информацию о зарезервированном времени, мфц и услуге
// swagger:response ApiResponseReserveTime
type ReserveTimeResponseWrapper struct {
	// in:body
	Body ApiResponseReserveTime
}

type ApiResponseTimeConfirmation struct {
	Success bool                       `json:"success"`
	Message string                     `json:"message"`
	Model   model.ConfirmationResponse `json:"model"`
}

// swagger:route POST /TimeConfirmation idOfTimeConfirmationEndpoint
// TimeConfirmation - метод подтверждения записи
// responses:
//   200: ApiResponseTimeConfirmation

// swagger:parameters TimeConfirmation idOfTimeConfirmationEndpoint
type ConfirmationRequest struct {
	// unique: true
	// required: true
	// in: body
	Body model.ConfirmationRequest
}

// Модель данных возвращает информацию о подтверждении записи
// swagger:response ApiResponseTimeConfirmation
type TimeConfirmationResponseWrapper struct {
	// in:body
	Body ApiResponseTimeConfirmation
}

type ApiResponseGetReservationCode struct {
	Success bool                          `json:"success"`
	Message string                        `json:"message"`
	Model   model.ReservationCodeResponse `json:"model"`
}

// swagger:route GET /GetReservationCode idOfGetReservationCodeEndpoint
// GetReservatonCode позволяет получить информацию по записи
// responses:
//   200: ApiResponseGetReservationCode

// Модель данных возвращает информацию по записи
// swagger:response ApiResponseGetReservationCode
type GetReservationCodeResponseWrapper struct {
	// in:body
	Body ApiResponseGetReservationCode
}

// swagger:parameters GetReservationCode idOfGetReservationCodeEndpoint
type PublicIdParam struct {
	// - publicId - параметр из запроса TimeConfirmation
	// unique: true
	// required: true
	// in: query
	// example: "19232323"
	PublicId string `json:"publicId"`
}
