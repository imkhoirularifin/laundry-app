package entity

type Employee struct {
	Id        string `json:"id"`
	Name      string `json:"name" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Username  string `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
