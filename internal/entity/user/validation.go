package user

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/ogabrielrodrigues/imobiliary/internal/types/response"
)

func (u *User) validate() *response.Err {
	if u.Fullname == "" {
		return response.NewErr(http.StatusBadRequest, ERR_FULLNAME_EMPTY)
	}

	if len(u.Fullname) < 10 || len(u.Fullname) > 100 {
		return response.NewErr(http.StatusUnprocessableEntity, ERR_FULLNAME_INVALID)
	}

	if u.Email == "" {
		return response.NewErr(http.StatusBadRequest, ERR_EMAIL_EMPTY)
	}

	if match, _ := regexp.MatchString(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`, u.Email); !match {
		return response.NewErr(http.StatusUnprocessableEntity, ERR_EMAIL_INVALID)
	}

	if u.CreciID == "" {
		return response.NewErr(http.StatusBadRequest, ERR_CRECIID_EMPTY)
	}

	if len(u.CreciID) != 7 {
		return response.NewErr(http.StatusUnprocessableEntity, ERR_CRECIID_INVALID)
	}

	if u.Cellphone == "" {
		return response.NewErr(http.StatusBadRequest, ERR_CELLPHONE_EMPTY)
	}

	if match, _ := regexp.MatchString(`^\(\d{2}\)\s\d{4,5}-\d{4}$`, u.Cellphone); !match {
		fmt.Println(u.Cellphone)
		fmt.Println(match)
		return response.NewErr(http.StatusUnprocessableEntity, ERR_CELLPHONE_INVALID)
	}

	if u.password == "" {
		return response.NewErr(http.StatusBadRequest, ERR_PASSWORD_EMPTY)
	}

	if len(u.password) < 8 {
		return response.NewErr(http.StatusUnprocessableEntity, ERR_PASSWORD_INVALID)
	}

	return nil
}
