package model

type ApiResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Model   interface{} `json:"model"`
}
