package entities

// Response is a model for each response in app.
type Response struct {
	Ok      bool `json:"ok"`
	Error   bool `json:"error"`
	Message any  `json:"message"`
	Data    any  `json:"data"`
}

// ResponseWithUser is a model to use with Response model.
// It can be returned like: { ..., message: null, data: { user: { ... } }}
type ResponseWithUser struct {
	User *User `json:"user"`
}
