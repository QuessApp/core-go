package errors

const (
	CONTENT_REQUIRED = "campo de conteúdo é obrigatório"
	CONTENT_LENGTH   = "campo de conteúdo deve conter entre 3 a 250 caracteres"

	SEND_TO_REQUIRED = "campo de destinatário é obrigatório"
	SEND_TO_LENGTH   = "campo de destinatário deve conter entre 3 a 50 caracteres"

	QUESTION_NOT_FOUND           = "a pergunta com o id informado não existe"
	QUESTION_NOT_SENT_FOR_ME     = "esta pergunta não foi enviada para você"
	SENDING_QUESTION_TO_YOURSELF = "você não pode enviar uma pergunta para si mesmo"

	QUESTION_NOT_AUTHORIZED = "você não possui acesso à essa pergunta"

	REACHED_QUESTIONS_LIMIT = "você não pode enviar esta pergunta porque já atingiu o limite de envios"

	CANT_DELETE_QUESTION_NOT_SENT_BY_YOU = "você não pode deletar esta pergunta porque ela não enviada por você"

	CANT_HIDE_ALREADY_HIDDEN = "você não pode ocultar esta pergunta porque ela já está oculta"
)
