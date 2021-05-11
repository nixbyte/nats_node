package model

type PersonsCountResponse struct {
	Count int64 `json:"count"`
}

type AllStoryResponse []struct {
	PersonName string `json:"personName"`
	PortalLink string `json:"portalLink"`
	PostLink   string `json:"postLink"`
	Story      string `json:"story"`
}

type PersonsListResponse struct {
	Content []struct {
		AwardingDocs []struct {
			Cipher  string `json:"cipher"`
			DocType string `json:"docType"`
			Images  []struct {
				ArchiveID              int64  `json:"archiveId"`
				AvailableForView       bool   `json:"availableForView"`
				AvailableImagesForView bool   `json:"availableImagesForView"`
				EntityID               int64  `json:"entityId"`
				EntityImageID          int64  `json:"entityImageId"`
				EntityKind             string `json:"entityKind"`
				Height                 int64  `json:"height"`
				Img                    struct {
					ID   int64  `json:"id"`
					Path string `json:"path"`
				} `json:"img"`
				Name        string `json:"name"`
				ObjectOrder int64  `json:"objectOrder"`
				Size        int64  `json:"size"`
				View        struct {
					ID   int64  `json:"id"`
					Path string `json:"path"`
				} `json:"view"`
				Width int64 `json:"width"`
			} `json:"images"`
			Name string `json:"name"`
		} `json:"awardingDocs"`
		AwardingPlace  string `json:"awardingPlace"`
		BirthYear      int64  `json:"birthYear"`
		DecisionDate   string `json:"decisionDate"`
		DecisionNumber string `json:"decisionNumber"`
		DocID          string `json:"docId"`
		DocumentNumber string `json:"documentNumber"`
		EntityID       int64  `json:"entityId"`
		Fio            string `json:"fio"`
		FirstName      string `json:"firstName"`
		JoinedFields   string `json:"joinedFields"`
		LastName       string `json:"lastName"`
		NextDocID      string `json:"nextDocId"`
		Patronymic     string `json:"patronymic"`
		PlaceOfWork    string `json:"placeOfWork"`
		Sex            string `json:"sex"`
	} `json:"content"`
	Total int64 `json:"total"`
}

type AddWidgetRequest struct {
	AccessToken string `json:"accessToken"`
	CommunityID string `json:"communityId"`
	WidgetType  string `json:"widgetType"`
}

type AddWidgetResponse struct {
	Body            struct{} `json:"body"`
	StatusCode      string   `json:"statusCode"`
	StatusCodeValue int64    `json:"statusCodeValue"`
}

type NotificationAddRequest struct {
	SearchCondition struct {
		BirthYear      int64  `json:"birthYear"`
		BirthYearFrom  int64  `json:"birthYearFrom"`
		BirthYearTo    int64  `json:"birthYearTo"`
		DocID          string `json:"docId"`
		DocumentNumber string `json:"documentNumber"`
		Fio            string `json:"fio"`
		FirstName      string `json:"firstName"`
		Fuzziness      string `json:"fuzziness"`
		LastName       string `json:"lastName"`
		Limit          int64  `json:"limit"`
		MaxExpansions  int64  `json:"maxExpansions"`
		NotEmpty       bool   `json:"notEmpty"`
		Offset         int64  `json:"offset"`
		Orders         []struct {
			Direction string `json:"direction"`
			Property  string `json:"property"`
		} `json:"orders"`
		Page           int64  `json:"page"`
		Patronymic     string `json:"patronymic"`
		PlaceOfWork    string `json:"placeOfWork"`
		PrefixLength   int64  `json:"prefixLength"`
		Query          string `json:"query"`
		Sex            string `json:"sex"`
		Size           int64  `json:"size"`
		Transpositions bool   `json:"transpositions"`
	} `json:"searchCondition"`
	Type  string `json:"type"`
	Value string `json:"value"`
}

type NotificationAddResponse struct {
	Body            struct{} `json:"body"`
	StatusCode      string   `json:"statusCode"`
	StatusCodeValue int64    `json:"statusCodeValue"`
}

type NotificationUnsubscribeRequest struct {
	Id    int    `json:"id"`
	Token string `json:"token"`
}

type PostAddRequest struct {
	PostID      int64  `json:"postId"`
	PostLink    string `json:"postLink"`
	PostMessage string `json:"postMessage"`
	VkUserID    int64  `json:"vkUserId"`
}
