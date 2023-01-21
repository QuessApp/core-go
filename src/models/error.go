package models

// RequestError is a model for each request error in app.
type RequestError struct {
	Message interface{} `json:"message"`
	Status  int         `json:"status"`
}
