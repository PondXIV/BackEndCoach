package coachdto

type IDCoachDTO struct {
	Cid int `json:"cid"`
}
type UpdateStatusCoachDTO struct {
	CoID   int    `json:"coID"`
	Status string `json:"status"`
}
