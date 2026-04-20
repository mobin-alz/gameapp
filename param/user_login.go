package param

type LoginRequest struct {
	PhoneNumber string `json:"phone_number" binding:"required"`
	Password    string `json:"password" binding:"required"`
}

type LoginResponse struct {
	User   UserInfo `json:"mysqluser"`
	Tokens Tokens   `json:"tokens"`
}
