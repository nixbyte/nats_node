package model

import (
	model "nats_node/http/model/soap"
	"time"
)

type SoapEnvelope interface {
	GetSoapEnvelopeRequest() *model.EnvelopeRequest
}

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

type GetProblemRequest struct {
	ProblemId string `json:"problem_id,omitempty"`
}

type GetFileRequest struct {
	Filename string `json:"filename,omitempty"`
}

func (object GetAllProblemsRequest) GetSoapEnvelopeRequest() *model.EnvelopeRequest {
	problemsEnvelope := &model.EnvelopeRequest{}
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

func (object GetProblemRequest) GetSoapEnvelopeRequest() *model.EnvelopeRequest {
	problemEnvelope := &model.EnvelopeRequest{}
	problemEnvelope.Soapenv = "http://schemas.xmlsoap.org/soap/envelope/"
	problemEnvelope.Gorod = "https://gorod.gov.spb.ru/smev/gorod"
	problemEnvelope.Rev = "http://smev.gosuslugi.ru/rev120315"
	problemRequest := model.MessageBody{}

	problemRequest.Message.TypeCode = "GSRV"
	problemRequest.Message.Status = "REQUEST"
	problemRequest.Message.TestMsg = "FALSE"
	problemRequest.Message.ExchangeType = "2"
	currentTime := time.Now()
	problemRequest.Message.Date = currentTime.Format("2006-01-02T15:04:05Z")

	problemRequest.Message.Sender = model.Sender{
		"",
		"SPB010000",
		"Система классификаторов",
	}
	problemRequest.Message.Recipient = model.Recipient{
		"",
		"SPB010000",
		"Наш Санкт-Петербург",
	}

	problem := model.ProblemDetailRequest{
		Id: object.ProblemId,
	}

	problemRequest.MessageData = model.MessageData{}
	problemRequest.MessageData.AppData = model.ProblemDetailAppData{
		"",
		problem,
	}

	problemEnvelope.Body = model.GetProblemBody{
		"",
		"http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-utility-1.0.xsd",
		problemRequest,
	}

	return problemEnvelope

}

func (object GetFileRequest) GetSoapEnvelopeRequest() *model.EnvelopeRequest {
	Envelope := createEnvelope()
	MessageBody := createMessageBody()

	filename := model.GetFileRequest{
		URL: object.Filename,
	}

	MessageBody.MessageData.AppData = model.FileAppData{
		"",
		filename,
	}

	Envelope.Body = model.GetFileBody{
		"",
		"http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-utility-1.0.xsd",
		*MessageBody,
	}

	return Envelope

}

func createEnvelope() *model.EnvelopeRequest {
	problemEnvelope := &model.EnvelopeRequest{}
	problemEnvelope.Soapenv = "http://schemas.xmlsoap.org/soap/envelope/"
	problemEnvelope.Gorod = "https://gorod.gov.spb.ru/smev/gorod"
	problemEnvelope.Rev = "http://smev.gosuslugi.ru/rev120315"
	return problemEnvelope
}

func createMessageBody() *model.MessageBody {

	MessageBody := &model.MessageBody{}

	MessageBody.Message.TypeCode = "GSRV"
	MessageBody.Message.Status = "REQUEST"
	MessageBody.Message.TestMsg = "FALSE"
	MessageBody.Message.ExchangeType = "2"
	currentTime := time.Now()
	MessageBody.Message.Date = currentTime.Format("2006-01-02T15:04:05Z")

	MessageBody.Message.Sender = model.Sender{
		"",
		"SPB010000",
		"Система классификаторов",
	}
	MessageBody.Message.Recipient = model.Recipient{
		"",
		"SPB010000",
		"Наш Санкт-Петербург",
	}

	MessageBody.MessageData = model.MessageData{}
	return MessageBody
}
