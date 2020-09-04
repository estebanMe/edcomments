package models

// Message to API client
type Message struct {
	Message string `json:"message"`
	Code int `json:"code"`
}