package user

const (
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
	ERR_USER_NOT_FOUND_OR_NOT_EXISTS             string = "user not found or not exists"
	ERR_ONLY_ONE_MUST_PARAMETER_MUST_BE_PROVIDED string = "only one of the parameters must be provided"
	ERR_FAILED_GENERATE_TOKEN                    string = "failed to generate token"
	ERR_FAILED_TO_PROCESS_USER                   string = "failed to process user"
	ERR_UUID_INVALID                             string = "user uuid is invalid"
	ERR_INVALID_USER_REQUEST_BODY                string = "user request body is invalid"
	ERR_AVATAR_NOT_FOUND                         string = "avatar not found"
)
