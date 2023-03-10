package auth

const (
	NICK_FIELD_REQUIRED = "campo de nick é obrigatório"
	NICK_FIELD_LENGTH   = "campo de nick deve conter entre 3 e 50 caracteres"

	PASSWORD_FIELD_REQUIRED = "campo de senha é obrigatório"
	PASSWORD_FIELD_LENGTH   = "campo de senha deve conter entre 6 a 200 caracteres"

	NAME_FIELD_REQUIRED = "campo de nome é obrigatório"
	NAME_FIELD_LENGTH   = "campo de e-mail deve conter entre 3 e 50 caracteres"

	EMAIL_FIELD_REQUIRED = "campo de e-mail é obrigatório"
	EMAIL_FIELD_LENGTH   = "campo de e-mail deve conter entre 5 a 200 caracteres"
	EMAIL_FORMAT_INVALID = "campo de e-mail possui formato inválido"

	USER_TO_BLOCK_REQUIRED   = "campo de quem deve ser bloqueado é obrigatório"
	USER_TO_UNBLOCK_REQUIRED = "campo de quem deve ser bloqueado é obrigatório"

	EMAIL_IN_USE = "email já em uso"
	NICK_IN_USE  = "nome de usuário já em uso"

	INCORRECT_SIGNIN_DATA = "credenciais informadas estão incorretas"
	LOCALE_FIELD_REQUIRED = "campo de localidade é obrigatório"
	LOCALE_FIELD_INVALID  = "campo de localidade contém valor inesperado"
)
