package model

import (
	model "nats_node/http/model/soap"
	"time"
)

type GetAllProblemsRequest struct {
	Page        string `json:"page,omitempty"`
	Size        string `json:"size,omitempty"`
	Query       string `json:"query,omitempty"`
	Status      string `json:"status,omitempty"`
	District    string `json:"district,omitempty"`
	Latitude    string `json:"latitude,omitempty"`
	Longitude   string `json:"longitude,omitempty"`
	CityObject  string `json:"city_object,omitempty"`
	Category    string `json:"category,omitempty"`
	Reason      string `json:"reason,omitempty"`
	UpdateAfter string `json:"update_after,omitempty"`
	SortBy      string `json:"sort_by,omitempty"`
}

func setEmptyStringParameter(value *string) {
	if value == nil {
		*value = ""
	}
}

func (object GetAllProblemsRequest) GetSoapEnvelopeRequest() *model.GetAllProblemsEnvelopeRequest {
	problemsEnvelope := &model.GetAllProblemsEnvelopeRequest{}
	problemsEnvelope.Soapenv = "http://schemas.xmlsoap.org/soap/envelope/"
	problemsEnvelope.Gorod = "https://gorod.gov.spb.ru/smev/gorod"
	problemsEnvelope.Rev = "http://smev.gosuslugi.ru/rev120315"
	problemsList := model.MessageBody{}

	problemsList.Message.TypeCode = "GSRV"
	problemsList.Message.Status = "REQUEST"
	problemsList.Message.TestMsg = "FALSE"
	problemsList.Message.ExchangeType = "2"
	currentTime := time.Now()
	problemsList.Message.Date = currentTime.Format("2006-01-02T15:04:05Z")

	problemsList.Message.Sender = model.Sender{
		"",
		"SPB010000",
		"Система классификаторов",
	}
	problemsList.Message.Recipient = model.Recipient{
		"",
		"SPB010000",
		"Наш Санкт-Петербург",
	}

	setEmptyStringParameter(&object.Status)

	problemsListRequest := model.ProblemListRequest{
		Page:         object.Page,
		ItemsPerPage: object.Size,
		Query:        object.Query,
		Status:       object.Status,
		District:     object.District,
		Latitude:     object.Latitude,
		Longitude:    object.Longitude,
		CityObject:   object.CityObject,
		Reason:       object.CityObject,
		UpdatedAfter: object.UpdateAfter,
		SortBy:       object.SortBy,
	}

	problemsList.MessageData = model.MessageData{}
	problemsList.MessageData.AppData = model.ProblemListAppData{
		"",
		problemsListRequest,
	}

	problemsEnvelope.Body = model.GetProblemsListBody{
		"",
		"http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-utility-1.0.xsd",
		problemsList,
	}

	return problemsEnvelope

}
