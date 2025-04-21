package owner

import (
	"net/http"
	"regexp"

	"github.com/ogabrielrodrigues/imobiliary/internal/types/response"
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
