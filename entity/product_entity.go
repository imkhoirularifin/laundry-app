package entity

type Product struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Price     int    `json:"price"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
