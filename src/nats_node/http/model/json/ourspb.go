package model

import (
	model "nats_node/http/model/soap"
	"time"
)

type GetAllProblemsRequest struct {
	Page        string `json:"page"`
	Size        string `json:"size"`
	Query       string `json:"query"`
	Status      string `json:"status"`
	District    string `json:"district"`
	Latitude    string `json:"latitude"`
	Longitude   string `json:"longitude"`
	CityObject  string `json:"city_object"`
	Category    string `json:"category"`
	Reason      string `json:"reason"`
	UpdateAfter string `json:"update_after"`
	SortBy      string `json:"sort_by"`
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

	problemsListRequest := model.ProblemListRequest{
		Page:         object.Page,
		ItemsPerPage: object.Size,
		Query:        object.Query,
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
