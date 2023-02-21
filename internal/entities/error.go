package entities

// RequestError is a model for each request error in app.
type RequestError struct {
	Message any `json:"message"`
	Status  int `json:"status"`
}
