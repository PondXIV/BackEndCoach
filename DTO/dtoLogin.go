package dto

type LoginDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Type     int    `json:"Type"`
}
type LoginFBDTO struct {
	FacebookID string `json:"facebookID"`
}
