package GoldEx

type Currency struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Currencies []Currency
