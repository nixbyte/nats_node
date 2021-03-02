package model

type ApiResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Model   interface{} `json:"model"`
}

type MfcList struct {
	Notifications []interface{} `json:"notifications"`
	BranchList    []struct {
		AddressState   interface{} `json:"addressState"`
		Phone          string      `json:"phone"`
		AddressCity    string      `json:"addressCity"`
		FullTimeZone   string      `json:"fullTimeZone"`
		TimeZone       string      `json:"timeZone"`
		AddressLine2   interface{} `json:"addressLine2"`
		AddressLine1   string      `json:"addressLine1"`
		Updated        int64       `json:"updated"`
		Created        int64       `json:"created"`
		Email          string      `json:"email"`
		Name           string      `json:"name"`
		PublicID       string      `json:"publicId"`
		Longitude      float64     `json:"longitude"`
		BranchPrefix   interface{} `json:"branchPrefix"`
		Latitude       float64     `json:"latitude"`
		AddressCountry interface{} `json:"addressCountry"`
		Custom         string      `json:"custom"`
		AddressZip     interface{} `json:"addressZip"`
	} `json:"branchList"`
	Meta struct {
		Start        string      `json:"start"`
		End          string      `json:"end"`
		TotalResults int         `json:"totalResults"`
		Offset       interface{} `json:"offset"`
		Limit        interface{} `json:"limit"`
		Fields       string      `json:"fields"`
		Arguments    struct {
		} `json:"arguments"`
	} `json:"meta"`
}
