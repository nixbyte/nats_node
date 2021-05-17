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

type ApiResponseAllProblems struct {
	Success bool                         `json:"success"`
	Message string                       `json:"message"`
	Model   model.GetAllProblemsResponse `json:"model"`
}

type ApiResponseFile struct {
	Success bool                  `json:"success"`
	Message string                `json:"message"`
	Model   model.GetFileResponse `json:"model"`
}

// swagger:route POST /GetAllProblems AllProblems idOfGetAllProblemsEndpoint
// GetAllProblems возвращает список проблем удовлетворяющих параметрам фильтрации
// responses:
//   200: ApiResponseAllProblems

// swagger:parameters GetAllProblems idOfGetAllProblemsEndpoint
type AllProblemsParam struct {
	//  - параметры фильтрации
	//
	// unique: true
	// required: true
	// in: body
	// example: "3429642"
	Body model.GetAllProblemsRequest
}

// Модель данных списка проблем
// swagger:response ApiResponseAllProblems
type GetAllProblemsResponseWrapper struct {
	// in:body
	Body ApiResponseAllProblems
}

// swagger:route GET /GetProblem Problem idOfGetProblemEndpoint
// GetProblem принимает на вход номер проблемы и возвращает детальную информацию о проблеме
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

// Модель данных детальной информации о проблеме
// swagger:response ApiResponseProblem
type GetProblemResponseWrapper struct {
	// in:body
	Body ApiResponseProblem
}

// swagger:route POST /GetFile File idOfGetFileEndpoint
// GetFile возвращает возвращает содержимое файла в кодировке base64
// responses:
//   200: ApiResponseFile

// swagger:parameters GetFile idOfGetFileEndpoint
type FileParam struct {
	//  - имя файла
	//
	// unique: true
	// required: true
	// in: body
	// example: "/storage/1/61e8bf31-4ca6-4a05-bcbc-d95ce038d3bd.jpg"
	Body model.GetFileRequest
}

// Модель данных файла
// swagger:response ApiResponseFile
type GetFileResponseWrapper struct {
	// in:body
	Body ApiResponseFile
}
