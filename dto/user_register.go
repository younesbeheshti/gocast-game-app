package dto

type RegisterRequest struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type RegisterResponse struct {
	User UserInfo `json:"user"`
}

type UserInfo struct {
	ID          uint   `json:"id"`
	PhoneNumber string `json:"phone_number"`
	Name        string `json:"name"`
}
