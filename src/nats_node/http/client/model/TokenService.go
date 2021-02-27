package model

type Token struct {
	Status  bool   `json:"success"`
	Code    int    `json:"resultcode"`
	Message string `json:"message"`
	Content string `json:"content"`
}
