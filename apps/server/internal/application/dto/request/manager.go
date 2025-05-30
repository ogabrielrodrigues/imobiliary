package request

type CreateManagerDTO struct {
	Fullname string `json:"fullname"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
