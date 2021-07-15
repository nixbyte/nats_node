package model

type KomsportModel struct {
	Firstname             string `json:"firstname"`
	Lastname              string `json:"lastname"`
	Middlename            string `json:"middlename"`
	ChangeDate            int64  `json:"change_date"`
	ChangeReason          string `json:"change_reason"`
	BirtDate              int64  `json:"birth_date"`
	BirthPlace            string `json:"birth_place"`
	Gender                string `json:"gender"`
	DisableCategory       string `json:"disable_category"`
	DisableFrom           int64  `json:"disable_from"`
	DisableTill           int64  `json:"disable_till"`
	SportRank             string `json:"sport_rank"`
	SportType             string `json:"sport_type"`
	SportDiscipline       string `json:"sport_discipline"`
	SportOrganization     string `json:"sport_organization"`
	QualificationCategory string `json:"qualification_category"`
	Disqualifications     string `json:"disqualifications"`
	Achivements           string `json:"achivements"`
	Activities            string `json:"activities"`
	SportObjects          string `json:"sport_objects"`
}
