package models

type Response struct {
	Status        bool        `json:"status"`
	StatusMessage string      `json:"status_message"`
	Data          interface{} `json:"data,omitempty"`
}
