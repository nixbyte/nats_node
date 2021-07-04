package model

type Token struct {
	Status  bool   `json:"success"`
	Code    int    `json:"resultcode"`
	Message string `json:"message"`
	Content string `json:"content"`
}

type TokenExpiration struct {
	Status  bool `json:"success"`
	Code    int  `json:"resultcode"`
	Content struct {
		Token     string `json:"token"`
		StartDate string `json:"startDate"`
		EndDate   string `json:"endDate"`
	} `json:"content"`
}

type Patient struct {
	Firstname     string `json:"firstname"`
	Lastname      string `json:"lastname"`
	Middlename    string `json:"middlename"`
	Birthday      int64  `json:"birthday"`
	Phone         string `json:"phone"`
	IdLpu         string `json:"idLpu"`
	IdPat         string `json:"idPat"`
	IdAppointment string `json:"idAppointment"`
	AriaNumber    string `json:"AreaNumber"`
	DocumentN     string `json:"documentNumber"`
	DocumentS     string `json:"documentSerial"`
	PolisN        string `json:"polisNumber"`
	PolisS        string `json:"polisSerial"`
	Snils         string `json:"snils"`
}

type CovidLpuNames struct {
	Names []string `json:"lpu_names"`
}

type CovidLpuIds struct {
	Ids []string `json:"lpu_ids"`
}

type CovidSpecialityNames struct {
	Names []string `json:"speciality_names"`
}

type CovidSpecialityIds struct {
	Ids []string `json:"speciality_ids"`
}

type CovidDocNames struct {
	Names []string `json:"doc_names"`
}

type CovidDocIds struct {
	Ids []string `json:"doc_ids"`
}

type CovidAppointmentTimes struct {
	Times []string `json:"appointment_times"`
}

type CovidAppointmentIds struct {
	Ids []string `json:"appointment_ids"`
}

type CovidAppointmentCount struct {
	DistrictName     string `json:"district_name"`
	DistrictId       string `json:"district_id"`
	AppointmentCount int    `json:"ticket_count"`
}

type CovidAppointmentCountResponse struct {
	Districts []CovidAppointmentCount `json:"ditricts"`
}
