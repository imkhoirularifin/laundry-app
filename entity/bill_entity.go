package entity

type Bill struct {
	Id         string `json:"id"`
	BillDate   string `json:"bill_date"`
	CustomerId string `json:"customer_id"`
	EmployeeId string `json:"employee_id"`
	TotalPrice int    `json:"total_price"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}
