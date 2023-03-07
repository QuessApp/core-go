package questions

const (
	CONTENT_REQUIRED = "campo de conteúdo é obrigatório"
	CONTENT_LENGTH   = "campo de conteúdo deve conter entre 1 a 250 caracteres"

	SEND_TO_REQUIRED = "campo de destinatário é obrigatório"
	SEND_TO_LENGTH   = "campo de destinatário deve conter entre 3 a 50 caracteres"

	QUESTION_NOT_FOUND       = "a pergunta com o id informado não existe"
	QUESTION_NOT_SENT_FOR_ME = "esta pergunta não foi enviada para você"

	SENDING_QUESTION_TO_YOURSELF = "você não pode enviar uma pergunta para si mesmo"

	QUESTION_NOT_AUTHORIZED = "você não possui acesso à essa pergunta"

	REACHED_QUESTIONS_LIMIT = "você não pode enviar esta pergunta, pois já atingiu o limite de envios"

	CANT_DELETE_QUESTION_NOT_SENT_BY_YOU = "você não pode deletar esta pergunta, pois ela não enviada por você"

	CANT_HIDE_ALREADY_HIDDEN = "você não pode ocultar esta pergunta, pois ela já está oculta"
	CANT_SEND_INVALID_ID     = "id do destinátário é invalido"

	QUESTION_ALREADY_REPLIED        = "esta pergunta já foi respondida anteriormente"
	CANT_EDIT_REPLY_NOT_REPLIED_YET = "não foi é possível editar resposta desta pergunta, pois ela ainda não foi respondida"
	CANT_EDIT_REPLY_REACHED_LIMIT   = "não foi é possível editar resposta desta pergunta, pois a mesma já foi editada cinco vezes"
)
