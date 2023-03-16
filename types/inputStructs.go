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
