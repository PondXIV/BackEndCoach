package dto

// type RegisterDTO []RegisterDTOElement

type RegisterDTO struct {
	AliasName string `json:"AliasName"`
	Password  string `json:"Password"`
	Email     string `json:"Email"`
	FullName  string `json:"FullName"`
	Birthday  string `json:"Birthday "`
	Gender    string `json:"Gender"`
	Phone     string `json:"Phone"`
	Image     string `json:"Image"`
	Weight    int64  `json:"Weight"`
	Height    int64  `json:"Height"`
	Price     int64  `json:"Price "`
}
