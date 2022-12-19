package dto

import "time"

type RegisterCusDTO struct {
	AliasName string    `json:"AliasName"`
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
