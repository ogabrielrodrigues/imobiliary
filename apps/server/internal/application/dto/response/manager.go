package response

type ManagerDTO struct {
	ID       string `json:"id"`
	Fullname string `json:"fullname"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
}
