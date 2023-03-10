package users

// UpdateProfileDTO is DTO for payload for update user profile handler.
type UpdateProfileDTO struct {
	Nick   string `json:"nick,omitempty"`
	Name   string `json:"name,omitempty"`
	Locale string `json:"locale,omitempty" bson:"locale"`
	Email  string `json:"email,omitempty"`
}

// UpdatePreferencesDTO is DTO for payload for update preferences handler.
type UpdatePreferencesDTO struct {
	EnanbleAPPPushNotifications bool `json:"enableAppPushNotifications" bson:"enableAppPushNotifications"`
	EnableAPPEmails             bool `json:"enableAppEmails" bson:"enableAppEmails"`
}
