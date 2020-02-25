package model

// Show ...
type Show struct {
	Title       string `json:"title"`
	AvatarURL   string `json:"imageURL"`
	Description string `json:"description"`
	About       string `json:"about"` // descrição mais longa
}
