package entity

type BillDetail struct {
	Id        string `json:"id"`
	BillId    string `json:"bill_id"`
	ProductId string `json:"product_id"`
	Quantity  int    `json:"quantity"`
	Subtotal  int    `json:"subtotal"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
