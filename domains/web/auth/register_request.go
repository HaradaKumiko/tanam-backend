package auth

type RegisterRequest struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	Fullname    string `json:"full_name"`
	PhoneNumber string `json:"phone_number"`
}
