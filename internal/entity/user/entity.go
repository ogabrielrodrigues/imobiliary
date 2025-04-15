package user

import (
	"fmt"
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/ogabrielrodrigues/imobiliary/internal/entity/plan"
	"github.com/ogabrielrodrigues/imobiliary/internal/types/response"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        uuid.UUID
	Fullname  string
	CreciID   string
	Cellphone string
	Email     string
	password  string
	Avatar    string
	Plan      plan.Plan
}

func (u *User) hashPwd() *response.Err {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.password), 14)
	if err != nil {
		return response.NewErr(http.StatusInternalServerError, ERR_FAILED_TO_PROCESS_USER)
	}

	u.password = string(hash)
	return nil
}

func (u *User) ComparePwd(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.password), []byte(password))

	return err == nil
}

func New(creci_id, fullname, cellphone, email, password string) (*User, *response.Err) {
	u := &User{
		ID:        uuid.New(),
		CreciID:   creci_id,
		Fullname:  fullname,
		Cellphone: cellphone,
		Email:     email,
		password:  password,
		Plan:      *plan.New(plan.PlanKindFree, 0, 30, 0, 30),
	}

	// TODO: fix on deployment
	u.Avatar = fmt.Sprintf("http://%s/users/%s/avatar", os.Getenv("SERVER_ADDR"), u.ID.String())

	if err := u.validate(); err != nil {
		return nil, err
	}

	if err := u.hashPwd(); err != nil {
		return nil, err
	}

	return u, nil
}

func (u *User) SetPassword(password string) {
	u.password = password
}

func (u *User) ChangePassword(password string) {
	u.password = password

	u.hashPwd()
}

func (u *User) ToDTO() *DTO {
	return &DTO{
		ID:        u.ID.String(),
		CreciID:   u.CreciID,
		Fullname:  u.Fullname,
		Cellphone: u.Cellphone,
		Email:     u.Email,
		Avatar:    u.Avatar,
		Plan:      *u.Plan.ToDTO(),
	}
}

func UserFromDTO(dto *DTO) *User {
	return &User{
		ID:        uuid.MustParse(dto.ID),
		CreciID:   dto.CreciID,
		Fullname:  dto.Fullname,
		Cellphone: dto.Cellphone,
		Email:     dto.Email,
		Avatar:    dto.Avatar,
		Plan: *plan.New(
			plan.PlanKind(dto.Plan.Kind),
			dto.Plan.Price,
			dto.Plan.PropertiesTotalQuota,
			dto.Plan.PropertiesUsedQuota,
			dto.Plan.PropertiesRemainingQuota,
		),
	}
}
