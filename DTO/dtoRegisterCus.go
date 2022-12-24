package dto

import "time"

type RegisterCusDTO struct {
	Username string    `json:"Username"`
	Password  string    `json:"Password"`
	Email     string    `json:"Email"`
	FullName  string    `json:"FullName"`
	Birthday  time.Time `json:"Birthday"`
	Gender    string    `json:"Gender"`
	Phone     string    `json:"Phone"`
	Image     string    `json:"Image"`
	Weight    int       `json:"Weight"`
	Height    int       `json:"Height"`
	Price     int       `json:"Price"`
}
