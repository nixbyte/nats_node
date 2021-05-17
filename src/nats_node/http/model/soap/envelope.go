package model

import (
	"encoding/xml"
)

type EnvelopeRequest struct {
	XMLName xml.Name    `xml:"soapenv:Envelope"`
	Text    string      `xml:",chardata"`
	Soapenv string      `xml:"xmlns:soapenv,attr"`
	Rev     string      `xml:"xmlns:rev,attr"`
	Gorod   string      `xml:"xmlns:gorod,attr"`
	Header  Header      `xml:"soapenv:Header"`
	Body    interface{} `xml:"soapenv:Body"`
}

type GetProblemsListBody struct {
	Text        string      `xml:",chardata"`
	Ns0         string      `xml:"xmlns:ns0,attr"`
	MessageBody MessageBody `xml:"gorod:GetProblemsList"`
}

type ProblemListAppData struct {
	Text               string      `xml:",chardata"`
	ProblemListRequest interface{} `xml:"gorod:ProblemListRequest"`
}

type ProblemListRequestWithoutStatus struct {
	Text         string `xml:",chardata"`
	Page         string `xml:"page"`
	ItemsPerPage string `xml:"items_per_page"`
	Query        string `xml:"query"`
	District     string `xml:"district"`
	Latitude     string `xml:"latitude"`
	Longitude    string `xml:"longitude"`
	CityObject   string `xml:"city_object"`
	Category     string `xml:"category"`
	Reason       string `xml:"reason"`
	UpdatedAfter string `xml:"updated_after"`
	SortBy       string `xml:"sort_by"`
}

type ProblemListRequest struct {
	Text         string      `xml:",chardata"`
	Page         string      `xml:"page"`
	ItemsPerPage string      `xml:"items_per_page"`
	Query        string      `xml:"query"`
	Status       interface{} `xml:"status"`
	District     string      `xml:"district"`
	Latitude     string      `xml:"latitude"`
	Longitude    string      `xml:"longitude"`
	CityObject   string      `xml:"city_object"`
	Category     string      `xml:"category"`
	Reason       string      `xml:"reason"`
	UpdatedAfter string      `xml:"updated_after"`
	SortBy       string      `xml:"sort_by"`
}

type GetProblemBody struct {
	Text        string      `xml:",chardata"`
	Ns0         string      `xml:"xmlns:ns0,attr"`
	MessageBody MessageBody `xml:"gorod:GetProblem"`
}

type ProblemDetailAppData struct {
	Text                 string               `xml:",chardata"`
	ProblemDetailRequest ProblemDetailRequest `xml:"gorod:ProblemDetailRequest"`
}

type ProblemDetailRequest struct {
	Text string `xml:",chardata"`
	Id   string `xml:"id"`
}

type GetFileBody struct {
	Text        string      `xml:",chardata"`
	Ns0         string      `xml:"xmlns:ns0,attr"`
	MessageBody MessageBody `xml:"gorod:GetFile"`
}

type GetFileRequest struct {
	XMLName xml.Name `xml:"gorod:FileRequest"`
	Text    string   `xml:",chardata"`
	URL     string   `xml:"url"`
}

type FileAppData struct {
	Text           string         `xml:",chardata"`
	GetFileRequest GetFileRequest `xml:"gorod:FileRequest"`
}

type GetAllProblemsEnvelopeResponse struct {
	XMLName   xml.Name `xml:"Envelope"`
	Text      string   `xml:",chardata"`
	Soap11env string   `xml:"soap11env,attr"`
	Tns       string   `xml:"tns,attr"`
	Header    struct {
		Text     string `xml:",chardata"`
		Security struct {
			Text                string `xml:",chardata"`
			Wsse                string `xml:"wsse,attr"`
			Actor               string `xml:"actor,attr"`
			BinarySecurityToken struct {
				Text         string `xml:",chardata"`
				Ns0          string `xml:"ns0,attr"`
				EncodingType string `xml:"EncodingType,attr"`
				ValueType    string `xml:"ValueType,attr"`
				ID           string `xml:"Id,attr"`
			} `xml:"BinarySecurityToken"`
			Signature struct {
				Text       string `xml:",chardata"`
				Ds         string `xml:"ds,attr"`
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
						Reference struct {
							Text      string `xml:",chardata"`
							URI       string `xml:"URI,attr"`
							ValueType string `xml:"ValueType,attr"`
						} `xml:"Reference"`
					} `xml:"SecurityTokenReference"`
				} `xml:"KeyInfo"`
			} `xml:"Signature"`
		} `xml:"Security"`
	} `xml:"Header"`
	Body struct {
		Text                    string `xml:",chardata"`
		Ns0                     string `xml:"ns0,attr"`
		ID                      string `xml:"Id,attr"`
		GetProblemsListResponse struct {
			Text    string `xml:",chardata"`
			Message struct {
				Text   string `xml:",chardata"`
				Smev   string `xml:"smev,attr"`
				Sender struct {
					Text string `xml:",chardata"`
					Code string `xml:"Code"`
					Name string `xml:"Name"`
				} `xml:"Sender"`
				Recipient struct {
					Text string `xml:",chardata"`
					Code string `xml:"Code"`
					Name string `xml:"Name"`
				} `xml:"Recipient"`
				Service struct {
					Text     string `xml:",chardata"`
					Mnemonic string `xml:"Mnemonic"`
					Version  string `xml:"Version"`
				} `xml:"Service"`
				TypeCode     string `xml:"TypeCode"`
				Status       string `xml:"Status"`
				Date         string `xml:"Date"`
				ExchangeType string `xml:"ExchangeType"`
				TestMsg      string `xml:"TestMsg"`
			} `xml:"Message"`
			MessageData struct {
				Text    string `xml:",chardata"`
				Smev    string `xml:"smev,attr"`
				AppData struct {
					Text  string `xml:",chardata"`
					Error struct {
						XMLName      xml.Name `xml:"Error"`
						Text         string   `xml:",chardata"`
						ErrorCode    string   `xml:"errorCode"`
						ErrorMessage string   `xml:"errorMessage"`
					} `xml:"Error"`
					GetProblemsListResult struct {
						Text        string `xml:",chardata"`
						Count       string `xml:"count"`
						CurrentPage string `xml:"current_page"`
						IsLastPage  string `xml:"is_last_page"`
						Results     struct {
							Text    string `xml:",chardata"`
							Problem []struct {
								Text        string `xml:",chardata"`
								ID          string `xml:"id"`
								Reason      string `xml:"reason"`
								Status      string `xml:"status"`
								UpdateTime  string `xml:"update_time"`
								Image       string `xml:"image"`
								FullAddress string `xml:"full_address"`
							} `xml:"Problem"`
						} `xml:"results"`
					} `xml:"GetProblemsListResult"`
				} `xml:"AppData"`
			} `xml:"MessageData"`
		} `xml:"GetProblemsListResponse"`
		Fault struct {
			XMLName     xml.Name `xml:"Fault"`
			Text        string   `xml:",chardata"`
			Faultcode   string   `xml:"faultcode"`
			Faultstring string   `xml:"faultstring"`
			Faultactor  string   `xml:"faultactor"`
		} `xml:"Fault"`
	} `xml:"Body"`
}

type GetProblemEnvelopeResponse struct {
	XMLName   xml.Name `xml:"Envelope"`
	Text      string   `xml:",chardata"`
	Soap11env string   `xml:"soap11env,attr"`
	Tns       string   `xml:"tns,attr"`
	Header    struct {
		Text     string `xml:",chardata"`
		Security struct {
			Text                string `xml:",chardata"`
			Wsse                string `xml:"wsse,attr"`
			Actor               string `xml:"actor,attr"`
			BinarySecurityToken struct {
				Text         string `xml:",chardata"`
				Ns0          string `xml:"ns0,attr"`
				EncodingType string `xml:"EncodingType,attr"`
				ValueType    string `xml:"ValueType,attr"`
				ID           string `xml:"Id,attr"`
			} `xml:"BinarySecurityToken"`
			Signature struct {
				Text       string `xml:",chardata"`
				Ds         string `xml:"ds,attr"`
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
						Reference struct {
							Text      string `xml:",chardata"`
							URI       string `xml:"URI,attr"`
							ValueType string `xml:"ValueType,attr"`
						} `xml:"Reference"`
					} `xml:"SecurityTokenReference"`
				} `xml:"KeyInfo"`
			} `xml:"Signature"`
		} `xml:"Security"`
	} `xml:"Header"`
	Body struct {
		Text               string `xml:",chardata"`
		Ns0                string `xml:"ns0,attr"`
		ID                 string `xml:"Id,attr"`
		GetProblemResponse struct {
			Text    string `xml:",chardata"`
			Message struct {
				Text   string `xml:",chardata"`
				Smev   string `xml:"smev,attr"`
				Sender struct {
					Text string `xml:",chardata"`
					Code string `xml:"Code"`
					Name string `xml:"Name"`
				} `xml:"Sender"`
				Recipient struct {
					Text string `xml:",chardata"`
					Code string `xml:"Code"`
					Name string `xml:"Name"`
				} `xml:"Recipient"`
				Service struct {
					Text     string `xml:",chardata"`
					Mnemonic string `xml:"Mnemonic"`
					Version  string `xml:"Version"`
				} `xml:"Service"`
				TypeCode     string `xml:"TypeCode"`
				Status       string `xml:"Status"`
				Date         string `xml:"Date"`
				ExchangeType string `xml:"ExchangeType"`
				TestMsg      string `xml:"TestMsg"`
			} `xml:"Message"`
			MessageData struct {
				Text    string `xml:",chardata"`
				Smev    string `xml:"smev,attr"`
				AppData struct {
					Text             string `xml:",chardata"`
					GetProblemResult struct {
						Text                string `xml:",chardata"`
						ID                  string `xml:"id"`
						CityObjectName      string `xml:"city_object_name"`
						CategoryName        string `xml:"category_name"`
						ReasonName          string `xml:"reason_name"`
						Reason              string `xml:"reason"`
						FullAddress         string `xml:"full_address"`
						Latitude            string `xml:"latitude"`
						Longitude           string `xml:"longitude"`
						DistrictID          string `xml:"district_id"`
						ExpectedAnswerDt    string `xml:"expected_answer_dt"`
						EarlierOrganization string `xml:"earlier_organization"`
						Controller          struct {
							Text             string `xml:",chardata"`
							Name             string `xml:"name"`
							PersonPosition   string `xml:"person_position"`
							PersonName       string `xml:"person_name"`
							PersonEmail      string `xml:"person_email"`
							PersonPhone      string `xml:"person_phone"`
							OrganizationInn  string `xml:"organization_inn"`
							OrganizationOgrn string `xml:"organization_ogrn"`
						} `xml:"controller"`
						Coordinator struct {
							Text             string `xml:",chardata"`
							Name             string `xml:"name"`
							PersonPosition   string `xml:"person_position"`
							PersonName       string `xml:"person_name"`
							PersonEmail      string `xml:"person_email"`
							PersonPhone      string `xml:"person_phone"`
							OrganizationInn  string `xml:"organization_inn"`
							OrganizationOgrn string `xml:"organization_ogrn"`
						} `xml:"coordinator"`
						Feed struct {
							Text   string `xml:",chardata"`
							Widget []struct {
								Text   string `xml:",chardata"`
								Status struct {
									Text        string `xml:",chardata"`
									Dt          string `xml:"dt"`
									Status      string `xml:"status"`
									CloseReason string `xml:"close_reason"`
								} `xml:"status"`
								Answer struct {
									Text       string `xml:",chardata"`
									Dt         string `xml:"dt"`
									AuthorName string `xml:"author_name"`
									Author     struct {
										Text             string `xml:",chardata"`
										Name             string `xml:"name"`
										PersonPosition   string `xml:"person_position"`
										PersonName       string `xml:"person_name"`
										PersonEmail      string `xml:"person_email"`
										PersonPhone      string `xml:"person_phone"`
										OrganizationInn  string `xml:"organization_inn"`
										OrganizationOgrn string `xml:"organization_ogrn"`
									} `xml:"author"`
									Body   string `xml:"body"`
									Photos struct {
										Text  string `xml:",chardata"`
										Photo []struct {
											Text         string `xml:",chardata"`
											ID           string `xml:"id"`
											OriginalName string `xml:"original_name"`
											URL          string `xml:"url"`
										} `xml:"Photo"`
									} `xml:"photos"`
									Files struct {
										Text string `xml:",chardata"`
										File struct {
											Text         string `xml:",chardata"`
											ID           string `xml:"id"`
											OriginalName string `xml:"original_name"`
											URL          string `xml:"url"`
										} `xml:"File"`
									} `xml:"files"`
									Performer   string `xml:"performer"`
									InterimDate string `xml:"interim_date"`
								} `xml:"answer"`
								Petition struct {
									Text       string `xml:",chardata"`
									Dt         string `xml:"dt"`
									AuthorName string `xml:"author_name"`
									Body       string `xml:"body"`
									Photos     struct {
										Text  string `xml:",chardata"`
										Photo []struct {
											Text         string `xml:",chardata"`
											ID           string `xml:"id"`
											OriginalName string `xml:"original_name"`
											URL          string `xml:"url"`
										} `xml:"Photo"`
									} `xml:"photos"`
									Files string `xml:"files"`
								} `xml:"petition"`
							} `xml:"Widget"`
						} `xml:"feed"`
						UpdateTime string `xml:"update_time"`
					} `xml:"GetProblemResult"`
				} `xml:"AppData"`
			} `xml:"MessageData"`
		} `xml:"GetProblemResponse"`
		Fault struct {
			XMLName     xml.Name `xml:"Fault"`
			Text        string   `xml:",chardata"`
			Faultcode   string   `xml:"faultcode"`
			Faultstring string   `xml:"faultstring"`
			Faultactor  string   `xml:"faultactor"`
		} `xml:"Fault"`
	} `xml:"Body"`
}

type GetFileEnvelopeResponse struct {
	XMLName   xml.Name `xml:"Envelope"`
	Text      string   `xml:",chardata"`
	Soap11env string   `xml:"soap11env,attr"`
	Tns       string   `xml:"tns,attr"`
	Header    struct {
		Text     string `xml:",chardata"`
		Security struct {
			Text                string `xml:",chardata"`
			Wsse                string `xml:"wsse,attr"`
			Actor               string `xml:"actor,attr"`
			BinarySecurityToken struct {
				Text         string `xml:",chardata"`
				Ns0          string `xml:"ns0,attr"`
				EncodingType string `xml:"EncodingType,attr"`
				ValueType    string `xml:"ValueType,attr"`
				ID           string `xml:"Id,attr"`
			} `xml:"BinarySecurityToken"`
			Signature struct {
				Text       string `xml:",chardata"`
				Ds         string `xml:"ds,attr"`
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
						Reference struct {
							Text      string `xml:",chardata"`
							URI       string `xml:"URI,attr"`
							ValueType string `xml:"ValueType,attr"`
						} `xml:"Reference"`
					} `xml:"SecurityTokenReference"`
				} `xml:"KeyInfo"`
			} `xml:"Signature"`
		} `xml:"Security"`
	} `xml:"Header"`
	Body struct {
		Text            string `xml:",chardata"`
		Ns0             string `xml:"ns0,attr"`
		ID              string `xml:"Id,attr"`
		GetFileResponse struct {
			Text    string `xml:",chardata"`
			Message struct {
				Text   string `xml:",chardata"`
				Smev   string `xml:"smev,attr"`
				Sender struct {
					Text string `xml:",chardata"`
					Code string `xml:"Code"`
					Name string `xml:"Name"`
				} `xml:"Sender"`
				Recipient struct {
					Text string `xml:",chardata"`
					Code string `xml:"Code"`
					Name string `xml:"Name"`
				} `xml:"Recipient"`
				Service struct {
					Text     string `xml:",chardata"`
					Mnemonic string `xml:"Mnemonic"`
					Version  string `xml:"Version"`
				} `xml:"Service"`
				TypeCode     string `xml:"TypeCode"`
				Status       string `xml:"Status"`
				Date         string `xml:"Date"`
				ExchangeType string `xml:"ExchangeType"`
				TestMsg      string `xml:"TestMsg"`
			} `xml:"Message"`
			MessageData struct {
				Text    string `xml:",chardata"`
				Smev    string `xml:"smev,attr"`
				AppData struct {
					Text          string `xml:",chardata"`
					GetFileResult struct {
						Text string `xml:",chardata"`
						File string `xml:"file"`
					} `xml:"GetFileResult"`
					Error struct {
						Text         string `xml:",chardata"`
						ErrorCode    string `xml:"errorCode"`
						ErrorMessage string `xml:"errorMessage"`
					} `xml:"Error"`
				} `xml:"AppData"`
			} `xml:"MessageData"`
		} `xml:"GetFileResponse"`
		Fault struct {
			XMLName     xml.Name `xml:"Fault"`
			Text        string   `xml:",chardata"`
			Faultcode   string   `xml:"faultcode"`
			Faultstring string   `xml:"faultstring"`
			Faultactor  string   `xml:"faultactor"`
		} `xml:"Fault"`
	} `xml:"Body"`
}
