package errors

var (
	NICK_FIELD_REQUIRED = "Campo de nick é obrigatório"
	NICK_FIELD_LENGTH   = "Campo de nick deve conter entre 3 e 50 caracteres"

	PASSWORD_FIELD_REQUIRED = "Campo de senha é obrigatório"
	PASSWORD_FIELD_LENGTH   = "Campo de senha deve conter entre 6 a 200 caracteres"

	NAME_FIELD_REQUIRED = "Campo de nome é obrigatório"
	NAME_FIELD_LENGTH   = "Campo de e-mail deve conter entre 3 e 50 caracteres"

	EMAIL_FIELD_REQUIRED = "Campo de e-mail é obrigatório"
	EMAIL_FIELD_LENGTH   = "Campo de e-mail deve conter entre 5 a 200 caracteres"
	EMAIL_FORMAT_INVALID = "Campo de e-mail possui formato inválido"
)
