package entities

// Location is a model for each location in app.
type Location struct {
	CountryCode string  `json:"country_code"`
	CountryName string  `json:"country_name"`
	City        string  `json:"city"`
	Postal      string  `json:"postal"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	IPv4        string  `json:"IPv4"`
	State       string  `json:"state"`
}
