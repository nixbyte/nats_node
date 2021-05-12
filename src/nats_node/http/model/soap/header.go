package model

import "encoding/xml"

type Header struct {
	XMLName          xml.Name `xml:"soapenv:Header"`
	Text             string   `xml:",chardata"`
	ApplicationToken string   `xml:"gorod:ApplicationToken"`
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
