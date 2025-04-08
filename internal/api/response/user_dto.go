package response

type UserResponse struct {
	ID       string `json:"id"`
	CreciID  string `json:"creci_id"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
}

type UserCreateDTO struct {
	ID       string `json:"id"`
	CreciID  string `json:"creci_id"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
}

type UserUpdateDTO struct {
	CreciID  string `json:"creci_id,omitempty"`
	Fullname string `json:"fullname,omitempty"`
	Email    string `json:"email,omitempty"`
}
