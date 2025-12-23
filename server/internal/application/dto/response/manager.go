package response

type ManagerDto struct {
	ID       string `json:"id"`
	Fullname string `json:"fullname"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Active   bool   `json:"active"`
}
