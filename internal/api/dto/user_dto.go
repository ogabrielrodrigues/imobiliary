package dto

type UserDTO struct {
	ID        string `json:"id"`
	CreciID   string `json:"creci_id"`
	Fullname  string `json:"fullname"`
	Cellphone string `json:"cellphone"`
	Email     string `json:"email"`
	Avatar    string `json:"avatar"`
}

type UserCreateDTO struct {
	ID        string `json:"id"`
	CreciID   string `json:"creci_id"`
	Fullname  string `json:"fullname"`
	Cellphone string `json:"cellphone"`
	Email     string `json:"email"`
}

type UserUpdateDTO struct {
	CreciID   string `json:"creci_id,omitempty"`
	Fullname  string `json:"fullname,omitempty"`
	Cellphone string `json:"cellphone,omitempty"`
	Email     string `json:"email,omitempty"`
}

type UserAuthDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
