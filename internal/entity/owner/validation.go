package owner

import (
	"net/http"
	"regexp"

	"github.com/ogabrielrodrigues/imobiliary/internal/types/response"
)

const (
	ERR_FULLNAME_EMPTY                string = "owner fullname is empty"
	ERR_FULLNAME_INVALID              string = "owner fullname must be between 10 and 100 characters"
	ERR_CPF_EMPTY                     string = "owner cpf is empty"
	ERR_CPF_INVALID                   string = "owner cpf is invalid format"
	ERR_RG_EMPTY                      string = "owner rg is empty"
	ERR_RG_INVALID                    string = "owner rg must is invalid"
	ERR_EMAIL_EMPTY                   string = "owner email is empty"
	ERR_CELLPHONE_EMPTY               string = "owner cellphone is empty"
	ERR_CELLPHONE_INVALID             string = "owner cellphone is invalid format"
	ERR_EMAIL_INVALID                 string = "owner email is invalid format"
	ERR_OWNER_NOT_FOUND_OR_NOT_EXISTS string = "owner not found or not exists"
	ERR_OWNER_BODY_INVALID            string = "owner body is invalid"
)

func (o *Owner) validate() *response.Err {
	if o.Fullname == "" {
		return response.NewErr(http.StatusBadRequest, ERR_FULLNAME_EMPTY)
	}

	if len(o.Fullname) < 10 || len(o.Fullname) > 100 {
		return response.NewErr(http.StatusUnprocessableEntity, ERR_FULLNAME_INVALID)
	}

	if o.CPF == "" {
		return response.NewErr(http.StatusBadRequest, ERR_CPF_EMPTY)
	}

	if match, _ := regexp.MatchString(`^\d{3}\.\d{3}\.\d{3}-\d{2}$`, o.CPF); !match {
		return response.NewErr(http.StatusUnprocessableEntity, ERR_CPF_INVALID)
	}

	if o.RG == "" {
		return response.NewErr(http.StatusBadRequest, ERR_RG_EMPTY)
	}

	if len(o.RG) < 5 || len(o.RG) > 15 {
		return response.NewErr(http.StatusUnprocessableEntity, ERR_RG_INVALID)
	}

	if o.Email == "" {
		return response.NewErr(http.StatusBadRequest, ERR_EMAIL_EMPTY)
	}

	if match, _ := regexp.MatchString(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`, o.Email); !match {
		return response.NewErr(http.StatusUnprocessableEntity, ERR_EMAIL_INVALID)
	}

	if o.Cellphone == "" {
		return response.NewErr(http.StatusBadRequest, ERR_CELLPHONE_EMPTY)
	}

	if match, _ := regexp.MatchString(`^\(\d{2}\)\s\d{4,5}-\d{4}$`, o.Cellphone); !match {
		return response.NewErr(http.StatusUnprocessableEntity, ERR_CELLPHONE_INVALID)
	}

	return nil
}
