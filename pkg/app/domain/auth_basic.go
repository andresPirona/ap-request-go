package domain

// Basic Authentication Entity
type BasicAuth struct {
	User string `json:"user"`
	Pass string `json:"pass"`
}
