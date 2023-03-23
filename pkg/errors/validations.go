package errors

const (
	NICK_FIELD_REQUIRED = "campo de nick é obrigatório"
	NICK_FIELD_LENGTH   = "campo de nick deve conter entre 3 e 50 caracteres"

	INCORRECT_SIGNIN_DATA = "credenciais informadas estão incorretas"

	PASSWORD_FIELD_REQUIRED = "campo de senha é obrigatório"
	PASSWORD_FIELD_LENGTH   = "campo de senha deve conter entre 6 a 200 caracteres"

	NAME_FIELD_REQUIRED = "campo de nome é obrigatório"
	NAME_FIELD_LENGTH   = "campo de e-mail deve conter entre 3 e 50 caracteres"

	EMAIL_FIELD_REQUIRED = "campo de e-mail é obrigatório"
	EMAIL_FIELD_LENGTH   = "campo de e-mail deve conter entre 5 a 200 caracteres"
	EMAIL_FORMAT_INVALID = "campo de e-mail possui formato inválido"

	EMAIL_IN_USE = "email já em uso"
	NICK_IN_USE  = "nome de usuário já em uso"

	LOGOUT_FROM_ALL_DEVICES_REQUIRED = "campo de logout de todos os dispositivos é obrigatório"
	LOGOUT_FROM_ALL_DEVICES_INVALID  = "campo de logout de todos os dispositivos contém valor inesperado"

	CONTENT_REQUIRED = "campo de conteúdo é obrigatório"
	CONTENT_LENGTH   = "campo de conteúdo deve conter entre 1 a 250 caracteres"

	SEND_TO_REQUIRED = "campo de destinatário é obrigatório"
	SEND_TO_LENGTH   = "campo de destinatário deve conter entre 3 a 50 caracteres"

	MAX_FILE_SIZE     = "por favor, enviar imagem com tamanho menor que 1MB"
	FILE_TYPE_INVALID = "tipo de arquivo informado não é permitido"

	LOCALE_FIELD_REQUIRED = "campo de localidade é obrigatório"
	LOCALE_FIELD_INVALID  = "campo de localidade contém valor inesperado"

	ENABLE_APP_EMAILS_FIELD_REQUIRED        = "campo para habilitar/desabilitar e-mails é obrigatório"
	ENABLE_APP_NOTIFICATIONS_FIELD_REQUIRED = "campo para habilitar/desabilitar notificações é obrigatório"

	REASON_FIELD_REQUIRED = "campo de motivo é obrigatório"
	REASON_FIELD_INVALID  = "campo de motivo contém valor inesperado"

	TYPE_FIELD_REQUIRED = "campo de tipo é obrigatório"
	TYPE_FIELD_INVALID  = "campo de tipo contém valor inesperado"
)

const (
	USER_NOT_FOUND           = "usuário não encontrado"
	USER_TO_BLOCK_REQUIRED   = "campo de quem deve ser bloqueado é obrigatório"
	USER_TO_UNBLOCK_REQUIRED = "campo de quem deve ser bloqueado é obrigatório"
)

const (
	QUESTION_NOT_FOUND       = "a pergunta com o id informado não existe"
	QUESTION_NOT_SENT_FOR_ME = "esta pergunta não foi enviada para você"
	QUESTION_NOT_AUTHORIZED  = "você não possui acesso à essa pergunta"
	QUESTION_ALREADY_REPLIED = "esta pergunta já foi respondida anteriormente"

	REACHED_QUESTIONS_LIMIT = "você não pode enviar esta pergunta, pois já atingiu o limite de envios"

	CANT_DELETE_QUESTION_NOT_SENT_BY_YOU = "você não pode deletar esta pergunta, pois ela não foi enviada por você"
	SENDING_QUESTION_TO_YOURSELF         = "você não pode enviar uma pergunta para si mesmo"

	CANT_HIDE_ALREADY_HIDDEN = "você não pode ocultar esta pergunta, pois ela já está oculta"
	CANT_SEND_INVALID_ID     = "id do destinátário é inválido"
)

const (
	CANT_EDIT_REPLY_NOT_REPLIED_YET = "não foi é possível editar resposta desta pergunta, pois ela ainda não foi respondida"
	CANT_EDIT_REPLY_REACHED_LIMIT   = "não foi é possível editar resposta desta pergunta, pois a mesma já foi editada cinco vezes"
)

const (
	DID_BLOCKED_RECEIVER     = "você não pode enviar esta pergunta porque você bloqueou o usuário"
	BLOCKED_BY_RECEIVER      = "o usuário destinatário bloqueou você"
	ALREADY_BLOCKED          = "você não pode bloquear este usuário porque você já o bloqueou"
	CANT_BLOCK_YOURSELF      = "você não pode se bloquear"
	CANT_UNBLOCK_NOT_BLOCKED = "você não pode remover o bloqueio deste usuário porque este usuário não está bloqueado"
)

const (
	REPORT_NOT_FOUND                   = "a denúncia com o id informado não existe"
	REPORT_NOT_AUTHORIZED              = "você não possui acesso à essa denúncia"
	CANT_DELETE_REPORT_NOT_SENT_BY_YOU = "você não pode deletar esta denúncia, pois ela não foi enviada por você"

	CANT_REPORT_ALREADY_SENT = "você não pode enviar essa denúncia, pois já enviou uma semelhante anteriormente"
	CANT_REPORT_YOURSELF     = "você não pode enviar uma denúncia para si mesmo"
)

const (
	TOKEN_NOT_FOUND = "o token fornecido não existe"
	TOKEN_EXPIRED   = "o token fornecido expirou"
)

const (
	CODE_NOT_FOUND = "o código fornecido não existe"
	CODE_REQUIRED  = "campo de código é obrigatório"
	CODE_EXPIRED   = "o código fornecido expirou"
)
