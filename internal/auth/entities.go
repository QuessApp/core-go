package auth

import (
	"github.com/coreos/go-oidc/v3/oidc"
)

type Tokens struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	IDToken      string `json:"idToken"`
	TokenType    string `json:"tokenType"`
	Expiry       string `json:"expiry"`
}

type ResponseWithAuthenticatedUserData struct {
	Tokens Tokens         `json:"tokens"`
	User   *oidc.UserInfo `json:"user"`
}
