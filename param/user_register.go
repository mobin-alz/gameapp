package param

// data transfer object

type RegisterRequest struct {
	Name        string `json:"name" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	Password    string `json:"password" binding:"required"`
}

type RegisterResponse struct {
	User UserInfo `json:"user"`
}
