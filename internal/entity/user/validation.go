package user

import (
	"net/http"
	"regexp"

	"github.com/ogabrielrodrigues/imobiliary/internal/response"
)

const (
	ERR_AVATAR_SIZE_INVALID          string = "o tamanho do arquivo de avatar deve ser menor que 3MB"
	ERR_MUST_BE_PROVIDE_AVATAR       string = "deve enviar um arquivo de avatar"
	ERR_AVATAR_FORMAT_INVALID        string = "o formato do avatar deve ser .jpeg, .png, .jpg ou .webp"
	ERR_USER_ALREADY_EXISTS          string = "usuário já existe"
	ERR_INTERNAL_SERVER_ERROR        string = "erro interno do servidor"
	ERR_FULLNAME_EMPTY               string = "nome completo não pode ser vazio"
	ERR_FULLNAME_INVALID             string = "nome completo deve ter entre 10 e 100 caracteres"
	ERR_EMAIL_EMPTY                  string = "email do usuário não pode ser vazio"
	ERR_EMAIL_INVALID                string = "email do usuário deve ser válido"
	ERR_CRECIID_EMPTY                string = "registro CRECI não pode ser vazio"
	ERR_CRECIID_INVALID              string = "registro CRECI deve ter 7 caracteres"
	ERR_CELLPHONE_EMPTY              string = "telefone celular não pode ser vazio"
	ERR_CELLPHONE_INVALID            string = "telefone celular deve ter o formato (xx) xxxxx-xxxx"
	ERR_PASSWORD_EMPTY               string = "senha não pode ser vazia"
	ERR_PASSWORD_INVALID             string = "senha deve ter no mínimo 8 caracteres"
	ERR_USER_NOT_FOUND_OR_NOT_EXISTS string = "usuário não encontrado ou não existente"
	ERR_PASSWORD_DONT_MATCH          string = "a senha do usuário não coincide com a cadastrada"
	ERR_FAILED_GENERATE_TOKEN        string = "erro ao gerar o token de autenticação"
	ERR_FAILED_TO_PROCESS_USER       string = "erro ao processar o usuário"
	ERR_UUID_INVALID                 string = "informe um id de usuário válido"
	ERR_INVALID_USER_REQUEST_BODY    string = "informe um corpo de requisição válido"
	ERR_AVATAR_NOT_FOUND             string = "avatar não encontrado ou não existe"
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
