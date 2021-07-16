package model

type AppStateResponse struct {
	Status    string `json:"status"`
	State     int    `json:"state"`
	StateName string `json:"state_name"`
	Error     int    `json:"error"`
}

type MfcListResponse struct {
	BranchList []struct {
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
}

type MfcServicesResponse struct {
	ServiceList []struct {
		AdditionalCustomerDuration int    `json:"additionalCustomerDuration"`
		Duration                   int    `json:"duration"`
		Updated                    int64  `json:"updated"`
		Created                    int64  `json:"created"`
		Name                       string `json:"name"`
		PublicID                   string `json:"publicId"`
		Active                     bool   `json:"active"`
		PublicEnabled              bool   `json:"publicEnabled"`
		Custom                     string `json:"custom"`
	} `json:"serviceList"`
}

type DatesResponse struct {
	Dates []string `json:"dates"`
}

type TimesResponse struct {
	Times []string `json:"times"`
}

type ReserveRequest struct {
	Services []ServiceId `json:"services"`
}

type ServiceId struct {
	PublicID string `json:"publicId"`
}

type ReservationRequestData struct {
	Branch    string `json:"branch"`
	ServiceId string `json:"serviceId"`
	Date      string `json:"date"`
	Time      string `json:"time"`
}

type ReserveResponse struct {
	Services []struct {
		AdditionalCustomerDuration int    `json:"additionalCustomerDuration"`
		Duration                   int    `json:"duration"`
		Updated                    int64  `json:"updated"`
		Created                    int64  `json:"created"`
		Name                       string `json:"name"`
		PublicID                   string `json:"publicId"`
		Active                     bool   `json:"active"`
		PublicEnabled              bool   `json:"publicEnabled"`
		Custom                     string `json:"custom"`
	} `json:"services"`
	AllDay   bool `json:"allDay"`
	Status   int  `json:"status"`
	Resource struct {
		Name   string `json:"name"`
		Custom string `json:"custom"`
	} `json:"resource"`
	Customers []interface{} `json:"customers"`
	Blocking  bool          `json:"blocking"`
	Title     string        `json:"title"`
	Start     string        `json:"start"`
	Created   int64         `json:"created"`
	Updated   int64         `json:"updated"`
	PublicID  string        `json:"publicId"`
	Branch    struct {
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
	} `json:"branch"`
	Notes  interface{} `json:"notes"`
	End    string      `json:"end"`
	Custom interface{} `json:"custom"`
}
type CustomersRequest struct {
	Customers []Customer `json:"customers"`
}

type Customer struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

type ConfirmationRequest struct {
	FirstName         string `json:"firstName"`
	LastName          string `json:"lastName"`
	Email             string `json:"email"`
	Phone             string `json:"phone"`
	ReservationTimeId string `json:"reservationTimeId"`
	PersonId          string `json:"personId"`
}
type ConfirmationResponse struct {
	Services []struct {
		AdditionalCustomerDuration int    `json:"additionalCustomerDuration"`
		Duration                   int    `json:"duration"`
		Updated                    int64  `json:"updated"`
		Created                    int64  `json:"created"`
		Name                       string `json:"name"`
		PublicID                   string `json:"publicId"`
		Active                     bool   `json:"active"`
		PublicEnabled              bool   `json:"publicEnabled"`
		Custom                     string `json:"custom"`
	} `json:"services"`
	AllDay   bool `json:"allDay"`
	Status   int  `json:"status"`
	Resource struct {
		Name   string `json:"name"`
		Custom string `json:"custom"`
	} `json:"resource"`
	Customers []struct {
		DateOfBirth              interface{} `json:"dateOfBirth"`
		DeletionTimestamp        string      `json:"deletionTimestamp"`
		AddressState             interface{} `json:"addressState"`
		LastName                 string      `json:"lastName"`
		Phone                    string      `json:"phone"`
		LastInteractionTimestamp string      `json:"lastInteractionTimestamp"`
		AddressCity              interface{} `json:"addressCity"`
		RetentionPolicy          string      `json:"retentionPolicy"`
		ExternalID               interface{} `json:"externalId"`
		AddressLine2             interface{} `json:"addressLine2"`
		AddressLine1             interface{} `json:"addressLine1"`
		ConsentTimestamp         interface{} `json:"consentTimestamp"`
		Updated                  interface{} `json:"updated"`
		Created                  int64       `json:"created"`
		Email                    string      `json:"email"`
		ConsentIdentifier        interface{} `json:"consentIdentifier"`
		Name                     string      `json:"name"`
		PublicID                 string      `json:"publicId"`
		FirstName                string      `json:"firstName"`
		AddressCountry           interface{} `json:"addressCountry"`
		DateOfBirthWithoutTime   interface{} `json:"dateOfBirthWithoutTime"`
		Custom                   interface{} `json:"custom"`
		IdentificationNumber     interface{} `json:"identificationNumber"`
		AddressZip               interface{} `json:"addressZip"`
	} `json:"customers"`
	Blocking bool   `json:"blocking"`
	Title    string `json:"title"`
	Start    string `json:"start"`
	Created  int64  `json:"created"`
	Updated  int64  `json:"updated"`
	PublicID string `json:"publicId"`
	Branch   struct {
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
	} `json:"branch"`
	Notes  string `json:"notes"`
	End    string `json:"end"`
	Custom string `json:"custom"`
}
type ReservationCodeResponse struct {
	Notifications []interface{} `json:"notifications"`
	Meta          struct {
		Start        string      `json:"start"`
		End          string      `json:"end"`
		TotalResults int         `json:"totalResults"`
		Offset       interface{} `json:"offset"`
		Limit        interface{} `json:"limit"`
		Fields       string      `json:"fields"`
		Arguments    struct {
		} `json:"arguments"`
	} `json:"meta"`
	Appointment struct {
		Services []struct {
			ID       int    `json:"id"`
			Name     string `json:"name"`
			PublicID string `json:"publicId"`
		} `json:"services"`
		AllDay   bool `json:"allDay"`
		Status   int  `json:"status"`
		Resource struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"resource"`
		Customers []struct {
			ID        int    `json:"id"`
			LastName  string `json:"lastName"`
			Name      string `json:"name"`
			PublicID  string `json:"publicId"`
			FirstName string `json:"firstName"`
		} `json:"customers"`
		ID       int    `json:"id"`
		Blocking bool   `json:"blocking"`
		Title    string `json:"title"`
		QpID     int    `json:"qpId"`
		Start    string `json:"start"`
		Created  int64  `json:"created"`
		Updated  int64  `json:"updated"`
		PublicID string `json:"publicId"`
		Branch   struct {
			ID       int    `json:"id"`
			Name     string `json:"name"`
			PublicID string `json:"publicId"`
		} `json:"branch"`
		Notes  string `json:"notes"`
		End    string `json:"end"`
		Custom string `json:"custom"`
	} `json:"appointment"`
}
