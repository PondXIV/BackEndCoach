package dto

type LoginDTO []LoginDTODTOElement

type LoginDTODTOElement struct {
	Login Login `json:"Login"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
