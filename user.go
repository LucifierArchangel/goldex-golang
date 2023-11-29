package GoldEx

type UserBalance struct {
	Balance float64 `json:"balance"`
}

type User struct {
	Id        int    `json:"id"`
	Role      string `json:"role"`
	Name      string `json:"name"`
	ApiKey    string `json:"apiKey"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	IsShowing bool   `json:"isShowing"`
}
