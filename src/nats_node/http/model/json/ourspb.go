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

type GetAllProblemsResponse struct {
	Text        string `json:"Text"`
	Count       string `json:"Count"`
	Currentpage string `json:"CurrentPage"`
	Islastpage  string `json:"IsLastPage"`
	Results     struct {
		Text    string `json:"Text"`
		Problem []struct {
			Text        string `json:"Text"`
			ID          string `json:"ID"`
			Reason      string `json:"Reason"`
			Status      string `json:"Status"`
			Updatetime  string `json:"UpdateTime"`
			Image       string `json:"Image"`
			Fulladdress string `json:"FullAddress"`
		} `json:"Problem"`
	} `json:"Results"`
}

type GetProblemRequest struct {
	ProblemId string `json:"problem_id,omitempty"`
}

type GetProblemResponse struct {
	Text                string `json:"Text"`
	ID                  string `json:"ID"`
	Cityobjectname      string `json:"CityObjectName"`
	Categoryname        string `json:"CategoryName"`
	Reasonname          string `json:"ReasonName"`
	Reason              string `json:"Reason"`
	Fulladdress         string `json:"FullAddress"`
	Latitude            string `json:"Latitude"`
	Longitude           string `json:"Longitude"`
	Districtid          string `json:"DistrictID"`
	Expectedanswerdt    string `json:"ExpectedAnswerDt"`
	Earlierorganization string `json:"EarlierOrganization"`
	Controller          struct {
		Text             string `json:"Text"`
		Name             string `json:"Name"`
		Personposition   string `json:"PersonPosition"`
		Personname       string `json:"PersonName"`
		Personemail      string `json:"PersonEmail"`
		Personphone      string `json:"PersonPhone"`
		Organizationinn  string `json:"OrganizationInn"`
		Organizationogrn string `json:"OrganizationOgrn"`
	} `json:"Controller"`
	Coordinator struct {
		Text             string `json:"Text"`
		Name             string `json:"Name"`
		Personposition   string `json:"PersonPosition"`
		Personname       string `json:"PersonName"`
		Personemail      string `json:"PersonEmail"`
		Personphone      string `json:"PersonPhone"`
		Organizationinn  string `json:"OrganizationInn"`
		Organizationogrn string `json:"OrganizationOgrn"`
	} `json:"Coordinator"`
	Feed struct {
		Text   string `json:"Text"`
		Widget []struct {
			Text   string `json:"Text"`
			Status struct {
				Text        string `json:"Text"`
				Dt          string `json:"Dt"`
				Status      string `json:"Status"`
				Closereason string `json:"CloseReason"`
			} `json:"Status"`
			Answer struct {
				Text       string `json:"Text"`
				Dt         string `json:"Dt"`
				Authorname string `json:"AuthorName"`
				Author     struct {
					Text             string `json:"Text"`
					Name             string `json:"Name"`
					Personposition   string `json:"PersonPosition"`
					Personname       string `json:"PersonName"`
					Personemail      string `json:"PersonEmail"`
					Personphone      string `json:"PersonPhone"`
					Organizationinn  string `json:"OrganizationInn"`
					Organizationogrn string `json:"OrganizationOgrn"`
				} `json:"Author"`
				Body   string `json:"Body"`
				Photos struct {
					Text  string      `json:"Text"`
					Photo interface{} `json:"Photo"`
				} `json:"Photos"`
				Files struct {
					Text string `json:"Text"`
					File struct {
						Text         string `json:"Text"`
						ID           string `json:"ID"`
						Originalname string `json:"OriginalName"`
						URL          string `json:"URL"`
					} `json:"File"`
				} `json:"Files"`
				Performer   string `json:"Performer"`
				Interimdate string `json:"InterimDate"`
			} `json:"Answer"`
			Petition struct {
				Text       string `json:"Text"`
				Dt         string `json:"Dt"`
				Authorname string `json:"AuthorName"`
				Body       string `json:"Body"`
				Photos     struct {
					Text  string      `json:"Text"`
					Photo interface{} `json:"Photo"`
				} `json:"Photos"`
				Files string `json:"Files"`
			} `json:"Petition"`
		} `json:"Widget"`
	} `json:"Feed"`
	Updatetime string `json:"UpdateTime"`
}

type GetFileRequest struct {
	Filename string `json:"filename,omitempty"`
}

type GetFileResponse struct {
	Text string `json:"Text"`
	File string `json:"File"`
}

func NewProblemListRequest(object GetAllProblemsRequest) interface{} {
	if object.Status != "" {
		return model.ProblemListRequest{
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
	} else {
		return model.ProblemListRequestWithoutStatus{
			Page:         object.Page,
			ItemsPerPage: object.Size,
			Query:        object.Query,
			District:     object.District,
			Latitude:     object.Latitude,
			Longitude:    object.Longitude,
			CityObject:   object.CityObject,
			Reason:       object.CityObject,
			UpdatedAfter: object.UpdateAfter,
			SortBy:       object.SortBy,
		}

	}
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

	problemsListRequest := NewProblemListRequest(object)

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
