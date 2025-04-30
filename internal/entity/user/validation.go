package user

import (
	"net/http"
	"regexp"

	"github.com/ogabrielrodrigues/imobiliary/internal/types/response"
)

const (
	ERR_INTERNAL_SERVER_ERROR                    string = "erro interno do servidor"
	ERR_FULLNAME_EMPTY                           string = "user fullname is empty"
	ERR_FULLNAME_INVALID                         string = "user fullname must be between 10 and 100 characters"
	ERR_EMAIL_EMPTY                              string = "user email is empty"
	ERR_EMAIL_INVALID                            string = "user email is invalid format"
	ERR_CRECIID_EMPTY                            string = "user creci_id is empty"
	ERR_CRECIID_INVALID                          string = "user creci_id is invalid format"
	ERR_CELLPHONE_EMPTY                          string = "user cellphone is empty"
	ERR_CELLPHONE_INVALID                        string = "user cellphone is invalid format"
	ERR_PASSWORD_EMPTY                           string = "user password is empty"
	ERR_PASSWORD_INVALID                         string = "user password is invalid must be at least 8 characters"
	ERR_USER_NOT_FOUND_OR_NOT_EXISTS             string = "usuário não encontrado ou não existente"
	ERR_PASSWORD_DONT_MATCH                      string = "user password don't match"
	ERR_ONLY_ONE_MUST_PARAMETER_MUST_BE_PROVIDED string = "only one of the parameters must be provided"
	ERR_FAILED_GENERATE_TOKEN                    string = "failed to generate token"
	ERR_FAILED_TO_PROCESS_USER                   string = "failed to process user"
	ERR_UUID_INVALID                             string = "informe um id de usuário válido"
	ERR_INVALID_USER_REQUEST_BODY                string = "user request body is invalid"
	ERR_AVATAR_NOT_FOUND                         string = "avatar not found"
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
