package models

type UserProfile struct {
	Email string `json:"email"`
	Username string `json:"display_name"`
	Country string `json:"country"`
	ExternalURLs map[string]string `json:"external_urls"`
	Service string
}
