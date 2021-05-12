package model

import "encoding/xml"

type GetAllProblemsEnvelopeRequest struct {
	XMLName xml.Name            `xml:"soapenv:Envelope"`
	Text    string              `xml:",chardata"`
	Soapenv string              `xml:"xmlns:soapenv,attr"`
	Rev     string              `xml:"xmlns:rev,attr"`
	Gorod   string              `xml:"xmlns:gorod,attr"`
	Header  Header              `xml:"soapenv:Header"`
	Body    GetProblemsListBody `xml:"soapenv:Body"`
}

type GetProblemsListBody struct {
	Text        string      `xml:",chardata"`
	Ns0         string      `xml:"xmlns:ns0,attr"`
	MessageBody MessageBody `xml:"gorod:GetProblemsList"`
}

type ProblemListAppData struct {
	Text               string             `xml:",chardata"`
	ProblemListRequest ProblemListRequest `xml:"gorod:ProblemListRequest"`
}

type ProblemListRequest struct {
	Text         string `xml:",chardata"`
	Page         string `xml:"page"`
	ItemsPerPage string `xml:"items_per_page"`
	Query        string `xml:"query"`
	//	Status       string `xml:"status"`
	//	District     string `xml:"district"`
	//	Latitude     string `xml:"latitude"`
	//	Longitude    string `xml:"longitude"`
	//	CityObject   string `xml:"city_object"`
	//	Category     string `xml:"category"`
	//	Reason       string `xml:"reason"`
	//	UpdatedAfter string `xml:"updated_after"`
	//	SortBy       string `xml:"sort_by"`
}
