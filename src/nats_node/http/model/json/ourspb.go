package model

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
