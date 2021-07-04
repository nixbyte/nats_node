package model

import "encoding/xml"

type SoapDistrictRequest struct {
	XMLName xml.Name `xml:"tem:GetDistrictList"`
	Guid    string   `xml:"tem:guid"`
}

type SoapDistrictListResponse struct {
	Body struct {
		GetDistrictListResponse struct {
			GetDistrictListResult struct {
				Success   string `xml:"Success"`
				Districts struct {
					District []District `xml:"District"`
				} `xml:"Districts"`
			} `xml:"GetDistrictListResult"`
		} `xml:"GetDistrictListResponse"`
	} `xml:"Body"`
}

type District struct {
	DistrictName string `xml:"DistrictName"`
	IdDistrict   string `xml:"IdDistrict"`
	Okato        string `xml:"Okato"`
}

type SoapLpuListRequest struct {
	XMLName    xml.Name `xml:"tem:GetLPUList"`
	IdDistrict string   `xml:"tem:idDistrict"`
	Guid       string   `xml:"tem:guid"`
}

type SoapLpuListResponse struct {
	Body struct {
		GetLPUListResponse struct {
			GetLPUListResult struct {
				Success string `xml:"Success"`
				ListLPU struct {
					Clinic []struct {
						Text         string `xml:",chardata"`
						Description  string `xml:"Description"`
						District     string `xml:"District"`
						IdLPU        string `xml:"IdLPU"`
						IsActive     string `xml:"IsActive"`
						LPUFullName  string `xml:"LPUFullName"`
						LPUShortName string `xml:"LPUShortName"`
						LPUType      string `xml:"LPUType"`
						Oid          struct {
							Text string `xml:",chardata"`
							Nil  string `xml:"nil,attr"`
						} `xml:"Oid"`
						PartOf struct {
							Text string `xml:",chardata"`
							Nil  string `xml:"nil,attr"`
						} `xml:"PartOf"`
					} `xml:"Clinic"`
				} `xml:"ListLPU"`
			} `xml:"GetLPUListResult"`
		} `xml:"GetLPUListResponse"`
	} `xml:"Body"`
}

type SoapCovidLpuListRequest struct {
	XMLName xml.Name `xml:"tem:GetLpus"`
	Guid    string   `xml:"tem:guid"`
}

type SoapCovidLpuListResponse struct {
	Body struct {
		GetLpusResponse struct {
			GetLpusResult struct {
				Lpus struct {
					Lpu []Lpu `xml:"Lpu"`
				} `xml:"Lpus"`
				Success string `xml:"Success"`
			} `xml:"GetLpusResult"`
		} `xml:"GetLpusResponse"`
	} `xml:"Body"`
}
type Lpu struct {
	Text                              string `xml:",chardata"`
	Address                           string `xml:"Address"`
	AppointmentsLastUpdate            string `xml:"AppointmentsLastUpdate"`
	CountOfAvailableCovidAppointments string `xml:"CountOfAvailableCovidAppointments"`
	DistrictId                        string `xml:"DistrictId"`
	DistrictName                      string `xml:"DistrictName"`
	Email                             struct {
		Text string `xml:",chardata"`
		Nil  string `xml:"nil,attr"`
	} `xml:"Email"`
	FullName     string `xml:"FullName"`
	ID           string `xml:"Id"`
	Oid          string `xml:"Oid"`
	Organization string `xml:"Organization"`
	Phone        struct {
		Text string `xml:",chardata"`
		Nil  string `xml:"nil,attr"`
	} `xml:"Phone"`
	ShortName string `xml:"ShortName"`
}

type SoapSpecialityListRequest struct {
	XMLName xml.Name `xml:"tem:GetSpesialityList"`
	IdLpu   string   `xml:"tem:idLpu"`
	Guid    string   `xml:"tem:guid"`
}

type SoapSpecialityListResponse struct {
	Body struct {
		GetSpesialityListResponse struct {
			GetSpesialityListResult struct {
				Success        string `xml:"Success"`
				ListSpesiality struct {
					Spesiality []struct {
						Text                   string `xml:",chardata"`
						CountFreeParticipantIE string `xml:"CountFreeParticipantIE"`
						CountFreeTicket        string `xml:"CountFreeTicket"`
						FerIdSpesiality        string `xml:"FerIdSpesiality"`
						IdSpesiality           string `xml:"IdSpesiality"`
						LastDate               string `xml:"LastDate"`
						NameSpesiality         string `xml:"NameSpesiality"`
						NearestDate            string `xml:"NearestDate"`
					} `xml:"Spesiality"`
				} `xml:"ListSpesiality"`
			} `xml:"GetSpesialityListResult"`
		} `xml:"GetSpesialityListResponse"`
	} `xml:"Body"`
}

type SoapDoctorListRequest struct {
	XMLName      xml.Name `xml:"tem:GetDoctorList"`
	IdSpesiality string   `xml:"tem:idSpesiality"`
	IdLpu        string   `xml:"tem:idLpu"`
	Guid         string   `xml:"tem:guid"`
}

type SoapDoctorListResponse struct {
	Body struct {
		GetDoctorListResponse struct {
			GetDoctorListResult struct {
				Success string `xml:"Success"`
				Docs    struct {
					Doctor []struct {
						Comment                string `xml:"Comment"`
						CountFreeParticipantIE string `xml:"CountFreeParticipantIE"`
						CountFreeTicket        string `xml:"CountFreeTicket"`
						IdDoc                  string `xml:"IdDoc"`
						LastDate               string `xml:"LastDate"`
						Name                   string `xml:"Name"`
						NearestDate            string `xml:"NearestDate"`
						Snils                  string `xml:"Snils"`
					} `xml:"Doctor"`
				} `xml:"Docs"`
			} `xml:"GetDoctorListResult"`
		} `xml:"GetDoctorListResponse"`
	} `xml:"Body"`
}

type SoapAppointmentListRequest struct {
	XMLName    xml.Name `xml:"tem:GetAvaibleAppointments"`
	IdDoc      string   `xml:"tem:idDoc"`
	IdLpu      string   `xml:"tem:idLpu"`
	VisitStart string   `xml:"tem:visitStart"`
	VisitEnd   string   `xml:"tem:visitEnd"`
	Guid       string   `xml:"tem:guid"`
}

type SoapAppointmentListResponse struct {
	Body struct {
		GetAvaibleAppointmentsResponse struct {
			GetAvaibleAppointmentsResult struct {
				Success          string `xml:"Success"`
				ListAppointments struct {
					Appointment []struct {
						Text          string `xml:",chardata"`
						Address       string `xml:"Address"`
						IdAppointment string `xml:"IdAppointment"`
						Num           string `xml:"Num"`
						Room          string `xml:"Room"`
						VisitEnd      string `xml:"VisitEnd"`
						VisitStart    string `xml:"VisitStart"`
					} `xml:"Appointment"`
				} `xml:"ListAppointments"`
			} `xml:"GetAvaibleAppointmentsResult"`
		} `xml:"GetAvaibleAppointmentsResponse"`
	} `xml:"Body"`
}

type SoapPatient struct {
	Birthday   string `xml:"hub:Birthday"`
	Name       string `xml:"hub:Name"`
	SecondName string `xml:"hub:SecondName"`
	Surname    string `xml:"hub:Surname"`
	AriaNumber string `xml:"hub:AriaNumber"`
	CellPhone  string `xml:"hub:CellPhone"`
	Document_N string `xml:"hub:Document_N"`
	Document_S string `xml:"hub:Document_S"`
	HomePhone  string `xml:"hub:HomePhone"`
	IdPat      string `xml:"hub:IdPat"`
	Polis_N    string `xml:"hub:Polis_N"`
	Polis_S    string `xml:"hub:Polis_S"`
	Snils      string `xml:"hub:Snils"`
}

type SoapCheckPatientRequest struct {
	XMLName xml.Name    `xml:"tem:CheckPatient"`
	Pat     SoapPatient `xml:"tem:pat"`
	IdLpu   string      `xml:"tem:idLpu"`
	Guid    string      `xml:"tem:guid"`
}

type SoapCheckPatientResponse struct {
	Body struct {
		CheckPatientResponse struct {
			CheckPatientResult struct {
				ErrorList struct {
					Error struct {
						Text             string `xml:",chardata"`
						ErrorDescription string `xml:"ErrorDescription"`
						IdError          string `xml:"IdError"`
					} `xml:"Error"`
				} `xml:"ErrorList"`
				Success string `xml:"Success"`
				IdPat   string `xml:"IdPat"`
			} `xml:"CheckPatientResult"`
		} `xml:"CheckPatientResponse"`
	} `xml:"Body"`
}

type SoapAddPatientRequest struct {
	XMLName xml.Name    `xml:"tem:AddNewPatient"`
	Pat     SoapPatient `xml:"tem:patient"`
	IdLpu   string      `xml:"tem:idLpu"`
	Guid    string      `xml:"tem:guid"`
}

type SoapAddPatientResponse struct {
	Body struct {
		AddNewPatientResponse struct {
			AddNewPatientResult struct {
				ErrorList struct {
					Error struct {
						Text             string `xml:",chardata"`
						ErrorDescription string `xml:"ErrorDescription"`
						IdError          string `xml:"IdError"`
					} `xml:"Error"`
				} `xml:"ErrorList"`
				Success string `xml:"Success"`
				IdPat   struct {
					Text string `xml:",chardata"`
					Nil  string `xml:"nil,attr"`
				} `xml:"IdPat"`
			} `xml:"AddNewPatientResult"`
		} `xml:"AddNewPatientResponse"`
	} `xml:"Body"`
}

type SoapUpdatePhoneRequest struct {
	XMLName   xml.Name `xml:"tem:UpdatePhoneByIdPat"`
	IdLpu     string   `xml:"tem:idLpu"`
	IdPat     string   `xml:"tem:idPat"`
	HomePhone string   `xml:"tem:homePhone"`
	CellPhone string   `xml:"tem:cellPhone"`
	Guid      string   `xml:"tem:guid"`
}

type SoapUpdatePhoneResponse struct {
	Body struct {
		UpdatePhoneByIdPatResponse struct {
			UpdatePhoneByIdPatResult struct {
				ErrorList struct {
					Text string `xml:",chardata"`
					Nil  string `xml:"nil,attr"`
				} `xml:"ErrorList"`
				Success string `xml:"Success"`
			} `xml:"UpdatePhoneByIdPatResult"`
		} `xml:"UpdatePhoneByIdPatResponse"`
	} `xml:"Body"`
}

type SoapSetAppointmentRequest struct {
	XMLName       xml.Name `xml:"tem:SetAppointment"`
	IdAppointment string   `xml:"tem:idAppointment"`
	IdLpu         string   `xml:"tem:idLpu"`
	IdPat         string   `xml:"tem:idPat"`
	Guid          string   `xml:"tem:guid"`
}

type SoapSetAppointmentResponse struct {
	Body struct {
		SetAppointmentResponse struct {
			SetAppointmentResult struct {
				ErrorList struct {
					Text string `xml:",chardata"`
					Nil  string `xml:"nil,attr"`
				} `xml:"ErrorList"`
				Success string `xml:"Success"`
				Type    string `xml:"Type"`
			} `xml:"SetAppointmentResult"`
		} `xml:"SetAppointmentResponse"`
	} `xml:"Body"`
}

type SoapDeleteAppointmentRequest struct {
	XMLName       xml.Name `xml:"tem:CreateClaimForRefusal"`
	IdLpu         string   `xml:"tem:idLpu"`
	IdPat         string   `xml:"tem:idPat"`
	IdAppointment string   `xml:"tem:idAppointment"`
	Guid          string   `xml:"tem:guid"`
}

type SoapDeleteAppointmentResponse struct {
	Body struct {
		CreateClaimForRefusalResponse struct {
			CreateClaimForRefusalResult struct {
				ErrorList struct {
					Text string `xml:",chardata"`
					Nil  string `xml:"nil,attr"`
				} `xml:"ErrorList"`
				Success string `xml:"Success"`
			} `xml:"CreateClaimForRefusalResult"`
		} `xml:"CreateClaimForRefusalResponse"`
	} `xml:"Body"`
}
