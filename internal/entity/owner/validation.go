package owner

import (
	"net/http"
	"regexp"

	"github.com/ogabrielrodrigues/imobiliary/internal/response"
)

const (
	ERR_FULLNAME_EMPTY                string = "nome completo do proprietário não pode ser vazio"
	ERR_FULLNAME_INVALID              string = "nome completo do proprietário deve ter entre 10 e 100 caracteres"
	ERR_CPF_EMPTY                     string = "cpf do proprietário não pode ser vazio"
	ERR_CPF_INVALID                   string = "cpf do proprietário deve ter o formato xxx.xxx.xxx-xx"
	ERR_RG_EMPTY                      string = "rg do proprietário não pode ser vazio"
	ERR_RG_INVALID                    string = "rg do proprietário deve ter entre 5 e 15 caracteres"
	ERR_EMAIL_EMPTY                   string = "email do proprietário não pode ser vazio"
	ERR_EMAIL_INVALID                 string = "email do proprietário deve ser válido"
	ERR_CELLPHONE_EMPTY               string = "email do proprietário deve ser vazio"
	ERR_CELLPHONE_INVALID             string = "telefone celular do proprietário deve ter o formato (xx) xxxxx-xxxx"
	ERR_OWNER_NOT_FOUND_OR_NOT_EXISTS string = "proprietário não encontrado ou não existente"
	ERR_OWNER_ALREADY_EXISTS          string = "proprietário já existe"
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
