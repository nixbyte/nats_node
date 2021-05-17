// Package classification Ourspb Service.
//
// API для сервиса - Получения данных через СМЭВ из системы НашСПб
//
//     Schemes: http
//     BasePath: /ourspb
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
package ourspb

import (
	model "nats_node/http/model/json"
)

type ApiResponseProblem struct {
	Success bool                     `json:"success"`
	Message string                   `json:"message"`
	Model   model.GetProblemResponse `json:"model"`
}

// swagger:route GET /GetProblem Problem idOfGetProblemEndpoint
// GetProblem Принимает на вход номер проблемы и возвращает детальную информацию о проблеме
// responses:
//   200: ApiResponseProblem

// swagger:parameters GetProblem idOfGetProblemEndpoint
type ProblemParam struct {
	//  - идентификтаор проблемы
	//
	// min items: 5
	// max items: 14
	// unique: true
	// required: true
	// in: query
	// example: "3429642"
	ProblemId string `json:"problemId"`
}

// Модель данных детальную информацию о проблеме
// swagger:response ApiResponseProblem
type GetProblemResponseWrapper struct {
	// in:body
	Body ApiResponseProblem
}
