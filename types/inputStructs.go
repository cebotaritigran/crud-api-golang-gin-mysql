package types

// binding is needed to be able to validated if the client entered fields(which are binding: required)
type LoginInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterInput struct {
	First_name string `json:"firstname"`
	Last_name  string `json:"lastname"`
	Email      string `json:"email" binding:"required"`
	Password   string `json:"password" binding:"required"`
}

// post test
type ManagerInput struct {
	First_name   string `json:"first_name" binding:"required"`
	Last_name    string `json:"last_name" binding:"required"`
	Phone_number string `json:"phone_number" binding:"required"`
	Email        string `json:"email" binding:"required"`
}
