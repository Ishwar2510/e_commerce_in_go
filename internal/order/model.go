package order

type Order struct {
	Id         int     `json:"id"`
	UserId     int     `json:"user_id"`
	ProductId  int     `json:"product_id"`
	Quantity   int     `json:"quantity"`
	TotalValue float64 `json:"total"`
	Status     string  `json:"status"`
}
