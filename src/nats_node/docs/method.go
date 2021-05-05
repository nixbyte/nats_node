package docs

import (
	"nats_node/http/model"
)

type ApiResponsePersonCount struct {
	Success bool                       `json:"success"`
	Message string                     `json:"message"`
	Model   model.PersonsCountResponse `json:"model"`
}

type ApiResponseSearchPerson struct {
	Success bool                      `json:"success"`
	Message string                    `json:"message"`
	Model   model.PersonsListResponse `json:"model"`
}

type ApiResponseStory struct {
	Success bool                   `json:"success"`
	Message string                 `json:"message"`
	Model   model.AllStoryResponse `json:"model"`
}

// swagger:route GET /GetPersonsCount Person idOfGetPersonsCountEndpoint
// GetPersonsCount позволяет получить общее количество персоналий
// responses:
//   200: ApiResponsePersonCount

// Модель данных возвращает количество персоналий
// swagger:response ApiResponsePersonCount
type GetPersonsCountResponseWrapper struct {
	// in:body
	Body ApiResponsePersonCount
}

// swagger:route GET /GetPersonsCountByName Person idOfGetPersonsCountByNameEndpoint
// GetPersonsCountByName позволяет получить количество персоналий по фамилии
// responses:
//   200: ApiResponsePersonCount

// swagger:parameters GetPersonsCountByName idOfGetPersonsCountByNameEndpoint
type NameParam struct {
	//  - фамилия
	//
	// min items: 3
	// max items: 10
	// unique: true
	// required: true
	// in: query
	// example: "Никонов"
	Name string `json:"name"`
}

// swagger:route GET /SearchPerson Person idOfSearchPersonEndpoint
// SearchPerson позволяет получить информацию о персоне по поисковому запросу
// responses:
//   200: ApiResponseSearchPerson

// swagger:parameters SearchPerson idOfSearchPersonEndpoint
type SearchParam struct {
	// - страница
	//
	// unique: true
	// required: true
	// in: query
	Page int `json:"page"`
	// - размер страницы
	//
	// unique: true
	// required: true
	// in: query
	Size int `json:"size"`
	// - ФИО человека
	//
	// unique: true
	// in: query
	Fio string `json:"fio"`
	// - Год рождения человека
	//
	// unique: true
	// in: query
	BirthYear int `json:"birthYear"`
	// - Год рождения человека начиная с
	//
	// unique: true
	// in: query
	BirthYearFrom int `json:"birthYearFrom"`
	// - Год рождения человека по
	//
	// unique: true
	// in: query
	BirthYearTo int `json:"birthYearTo"`
	// - Место работы человека
	//
	// unique: true
	// in: query
	PlaceOfWork string `json:"placeOfWork"`
	// - Номер документа
	//
	// unique: true
	// in: query
	DocumentNumber string `json:"documentNumber"`
}

// Модель данных постранично возвращает информацию о персоне по поисковому запросу
// swagger:response ApiResponseSearchPerson
type SearchPersonResponseWrapper struct {
	// in:body
	Body ApiResponseSearchPerson
}

// swagger:route GET /GetAllStory Person idOfGetAllStoryEndpoint
// GetAllStory позволяет получить весь список мини историй
// responses:
//   200: ApiResponseStory

// Модель данных возвращает массив историй
// swagger:response ApiResponseStory
type GetAllStoryResponseWrapper struct {
	// in:body
	Body ApiResponseStory
}
