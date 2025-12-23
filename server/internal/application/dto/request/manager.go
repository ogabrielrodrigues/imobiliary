package request

type CreateManagerDto struct {
	Fullname string `json:"fullname"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ResetPasswordDto struct {
	Email string `json:"email"`
}
