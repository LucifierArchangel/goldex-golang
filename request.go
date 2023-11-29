package GoldEx

type Request struct {
	Id                 int     `json:"id"`
	RequestId          string  `json:"requestId"`
	Status             string  `json:"status"`
	CurrencyId         int     `json:"currencyId"`
	AdditionalAmount   float32 `json:"additionalAmount"`
	Amount             float64 `json:"amount"`
	CreatedAt          string  `json:"createdAt"`
	UpdatedAt          string  `json:"updatedAt"`
	CardHolder         string  `json:"cardHolder"`
	CardNumber         string  `json:"cardNumber"`
	Comment            string  `json:"comment"`
	RequestUSDTAmount  float64 `json:"requestUSDTAmount"`
	RequestTotalAmount float64 `json:"RequestTotalAmount"`
	CurrencyName       string  `json:"currencyName"`
	UserName           string  `json:"userName"`
	UserRole           string  `json:"userRole"`
}

type CreateRequest struct {
	RequestId  string  `json:"requestId"`
	Currency   int     `json:"currency"`
	Amount     float32 `json:"amount"`
	CardNumber string  `json:"cardNumber"`
	CardHolder string  `json:"cardHolder"`
}

type UpdateRequest struct {
	Status  string `json:"status"`
	Comment string `json:"comment,omitempty"`
}

type Requests []Request
