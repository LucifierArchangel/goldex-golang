package GoldEx

type ExchangeRate struct {
	Id                    int      `json:"id"`
	Value                 float32  `json:"value"`
	MinValue              float32  `json:"minValue"`
	MaxValue              float32  `json:"maxValue"`
	AdditionalAmountTo    float32  `json:"additionalAmountTo"`
	AdditionalAmountValue float32  `json:"additionalAmountValue"`
	CreatedAt             string   `json:"createdAt"`
	UpdatedAt             string   `json:"updatedAt"`
	Currency              Currency `json:"currency"`
}

type ExchangeRates []ExchangeRate
