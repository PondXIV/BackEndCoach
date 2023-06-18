package models

type Gbprimpay struct {
	Amount         float64 `json:"amount"`
	RetryFlag      string  `json:"retryFlag"`
	ReferenceNo    string  `json:"referenceNo"`
	GbpReferenceNo string  `json:"gbpReferenceNo"`
	CurrencyCode   string  `json:"currencyCode"`
	ResultCode     string  `json:"resultCode"`
	TotalAmount    float64 `json:"totalAmount"`
	Fee            float64 `json:"fee"`
	Vat            float64 `json:"vat"`
	ThbAmount      int64   `json:"thbAmount"`
	CustomerName   string  `json:"customerName"`
	Date           string  `json:"date"`
	Time           string  `json:"time"`
	PaymentType    string  `json:"paymentType"`
}
