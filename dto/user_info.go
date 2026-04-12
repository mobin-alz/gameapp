package dto

type UserInfo struct {
	ID          uint
	PhoneNumber string
	Name        string
}
type ProfileRequest struct {
	UserID uint
}
type ProfileResponse struct {
	Name string `json:"name"`
}
