package models

// Response is a model for each response of app.
type Response struct {
	Ok      bool        `json:"ok"`
	Message interface{} `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Status  int         `json:"status,omitempty"`
}
