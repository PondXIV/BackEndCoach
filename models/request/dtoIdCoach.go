package request

type IDCoachDTO struct {
	Cid int `json:"cid"`
}
type UpdateStatusCoachDTO struct {
	CoID   int    `json:"coID"`
	Status string `json:"status"`
}

type UpdateCoachDTO struct {
	CoID    int    `json:"coID"`
	Name    string `json:"name"`
	Details string `json:"details"`
	Level   string `json:"level"`
	Amount  int    `json:"amount"`
	Image   string `json:"image"`
	Days    string `json:"days"`
	Price   int    `json:"price"`
	Status  string `json:"status"`
}
