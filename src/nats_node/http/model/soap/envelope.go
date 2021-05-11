package model

import "encoding/xml"

type GetAllProblemsEnvelopeResponse struct {
	XMLName xml.Name    `xml:"Envelope"`
	Text    string      `xml:",chardata"`
	Soapenv string      `xml:"xmlns:soapenv,attr"`
	Rev     string      `xml:"xmlns:rev,attr"`
	Gorod   string      `xml:"xmlns:gorod,attr"`
	Header  Header      `xml:"soapenv:Header"`
	Body    interface{} `xml:"soapenv:Body"`
}
type GetAllProblemsEnvelope struct {
	XMLName xml.Name    `xml:"soapenv:Envelope"`
	Text    string      `xml:",chardata"`
	Soapenv string      `xml:"xmlns:soapenv,attr"`
	Rev     string      `xml:"xmlns:rev,attr"`
	Gorod   string      `xml:"xmlns:gorod,attr"`
	Header  interface{} `xml:"soapenv:Header"`
	Body    interface{} `xml:"soapenv:Body"`
}
type GetProblemsListBody struct {
	Text            string          `xml:",chardata"`
	Ns0             string          `xml:"xmlns:ns0,attr"`
	GetProblemsList GetProblemsList `xml:"gorod:GetProblemsList"`
}
type Sender struct {
	Text string `xml:",chardata"`
	Code string `xml:"rev:Code"`
	Name string `xml:"rev:Name"`
}
type Message struct {
	Text      string `xml:",chardata"`
	Sender    Sender `xml:"rev:Sender"`
	Recipient struct {
		Text string `xml:",chardata"`
		Code string `xml:"rev:Code"`
		Name string `xml:"rev:Name"`
	} `xml:"rev:Recipient"`
	TypeCode     string `xml:"rev:TypeCode"`
	Status       string `xml:"rev:Status"`
	Date         string `xml:"rev:Date"`
	ExchangeType string `xml:"rev:ExchangeType"`
	TestMsg      string `xml:"rev:TestMsg"`
}
type GetProblemsList struct {
	Text        string  `xml:",chardata"`
	Message     Message `xml:"rev:Message"`
	MessageData struct {
		Text    string `xml:",chardata"`
		AppData struct {
			Text               string `xml:",chardata"`
			ProblemListRequest struct {
				Text         string `xml:",chardata"`
				Page         string `xml:"page"`
				ItemsPerPage string `xml:"items_per_page"`
				Query        string `xml:"query"`
				Status       string `xml:"status"`
				District     string `xml:"district"`
				Latitude     string `xml:"latitude"`
				Longitude    string `xml:"longitude"`
				CityObject   string `xml:"city_object"`
				Category     string `xml:"category"`
				Reason       string `xml:"reason"`
				UpdatedAfter string `xml:"updated_after"`
				SortBy       string `xml:"sort_by"`
			} `xml:"gorod:ProblemListRequest"`
		} `xml:"rev:AppData"`
	} `xml:"rev:MessageData"`
}

type Header struct {
	XMLName          xml.Name `xml:"soapenv:Header"`
	Text             string   `xml:",chardata"`
	ApplicationToken string   `xml:"ApplicationToken"`
	Security         struct {
		Text                string `xml:",chardata"`
		Actor               string `xml:"actor,attr"`
		BinarySecurityToken struct {
			Text         string `xml:",chardata"`
			EncodingType string `xml:"EncodingType,attr"`
			ValueType    string `xml:"ValueType,attr"`
			ID           string `xml:"Id,attr"`
		} `xml:"BinarySecurityToken"`
		Signature struct {
			Text       string `xml:",chardata"`
			Xmlns      string `xml:"xmlns,attr"`
			ID         string `xml:"Id,attr"`
			SignedInfo struct {
				Text                   string `xml:",chardata"`
				CanonicalizationMethod struct {
					Text      string `xml:",chardata"`
					Algorithm string `xml:"Algorithm,attr"`
				} `xml:"CanonicalizationMethod"`
				SignatureMethod struct {
					Text      string `xml:",chardata"`
					Algorithm string `xml:"Algorithm,attr"`
				} `xml:"SignatureMethod"`
				Reference struct {
					Text       string `xml:",chardata"`
					URI        string `xml:"URI,attr"`
					Transforms struct {
						Text      string `xml:",chardata"`
						Transform struct {
							Text      string `xml:",chardata"`
							Algorithm string `xml:"Algorithm,attr"`
						} `xml:"Transform"`
					} `xml:"Transforms"`
					DigestMethod struct {
						Text      string `xml:",chardata"`
						Algorithm string `xml:"Algorithm,attr"`
					} `xml:"DigestMethod"`
					DigestValue string `xml:"DigestValue"`
				} `xml:"Reference"`
			} `xml:"SignedInfo"`
			SignatureValue string `xml:"SignatureValue"`
			KeyInfo        struct {
				Text                   string `xml:",chardata"`
				SecurityTokenReference struct {
					Text      string `xml:",chardata"`
					ID        string `xml:"Id,attr"`
					Reference struct {
						Text      string `xml:",chardata"`
						URI       string `xml:"URI,attr"`
						ValueType string `xml:"ValueType,attr"`
					} `xml:"Reference"`
				} `xml:"SecurityTokenReference"`
			} `xml:"KeyInfo"`
		} `xml:"Signature"`
	} `xml:"Security"`
}
