package entities

// ResponseWithUser is a model to use with Response model.
// It can be returned like: { ..., message: null, data: { user: { ... } }}
type ResponseWithUser struct {
	User         *User  `json:"user"`
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}
