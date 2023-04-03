package locales

// GetBrazilianPortugueseTranslation returns a map with all the translations for the Brazilian Portuguese locale.
func GetBrazilianPortugueseTranslation() *map[string]string {
	return &map[string]string{
		"nick_field_required": "campo de nick é obrigatório",
		"nick_field_length":   "campo de nick deve conter entre 3 e 50 caracteres",

		"incorrect_signin_data": "credenciais informadas estão incorretas",

		"password_field_required": "campo de senha é obrigatório",
		"password_field_length":   "campo de senha deve conter entre 6 a 200 caracteres",

		"name_field_required": "campo de nome é obrigatório",
		"name_field_length":   "campo de e-mail deve conter entre 3 e 50 caracteres",

		"email_field_required": "campo de e-mail é obrigatório",
		"email_field_length":   "campo de e-mail deve conter entre 5 a 200 caracteres",
		"email_format_invalid": "campo de e-mail possui formato inválido",
		"email_in_use":         "email já em uso",

		"nick_in_use": "nome de usuário já em uso",

		"logout_from_all_devices_required": "campo de logout de todos os dispositivos é obrigatório",
		"logout_from_all_devices_invalid":  "campo de logout de todos os dispositivos possui valor inválido",

		"content_required": "campo de conteúdo é obrigatório",
		"content_length":   "campo de conteúdo deve conter entre 1 e 250 caracteres",

		"send_to_required": "campo de enviar para é obrigatório",
		"send_to_length":   "campo de enviar para deve conter entre 3 e 50 caracteres",

		"max_file_size":     "por favor, envie uma imagem com tamanho menor que 1MB",
		"file_type_invalid": "tipo de arquivo informado não é permitido",

		"locale_field_required": "campo de localidade é obrigatório",
		"locale_field_invalid":  "campo de localidade possui valor inválido",

		"enable_app_emails_field_required":        "campo de ativar emails do app é obrigatório",
		"enable_app_notifications_field_required": "campo de ativar notificações do app é obrigatório",

		"reason_field_required": "campo de motivo é obrigatório",

		"type_field_required": "campo de tipo é obrigatório",
		"type_field_invalid":  "campo de tipo possui valor inválido",

		"user_not_found":           "usuário não encontrado",
		"user_to_block_required":   "campo de quem deve ser bloqueado é obrigatório",
		"user_to_unblock_required": "campo de quem deve ser bloqueado é obrigatório",

		"question_not_found":       "pergunta não encontrada",
		"question_not_sent_for_me": "esta pergunta não foi enviada para você",

		"question_not_authorized":  "você não possui acesso à essa pergunta",
		"question_already_replied": "esta pergunta já foi respondida anteriormente",

		"reached_questions_limit": "você não pode enviar esta pergunta, pois já atingiu o limite de envios",

		"cant_delete_question_not_sent_by_you": "você não pode deletar esta pergunta, pois ela não foi enviada por você",
		"sending_question_to_yourself":         "você não pode enviar uma pergunta para você mesmo",

		"can_hide_already_hidden": "você não pode ocultar esta pergunta, pois ela já está oculta",
		"cant_send_invalid_Id":    "id do destinátário é inválido",

		"cant_edit_reply_not_replied_yet": "não foi é possível editar resposta desta pergunta, pois ela ainda não foi respondida",
		"cant_edit_reply_reached_limit":   "não foi é possível editar resposta desta pergunta, pois a mesma já foi editada cinco vezes",

		"did_blocked_receiver": "você não pode enviar esta pergunta porque você bloqueou o usuário",

		"blocked_by_receiver":      "o usuário destinatário bloqueou você",
		"already_blocked":          "você não pode bloquear este usuário porque você já o bloqueou",
		"cant_block_yourself":      "você não pode se bloquear",
		"cant_unblock_not_blocked": "você não pode remover o bloqueio deste usuário porque este usuário não está bloqueado",

		"report_not_found":                   "a denúncia com o id informado não existe",
		"report_not_authorized":              "você não possui acesso à essa denúncia",
		"cant_delete_report_not_sent_by_you": "você não pode deletar esta denúncia, pois ela não foi enviada por você",
		"cant_report_already_sent":           "você não pode enviar essa denúncia, pois já enviou uma semelhante anteriormente",
		"cant_report_yourself":               "você não pode enviar uma denúncia para si mesmo",

		"token_not_found": "token not found",
		"token_expired":   "token expired",

		"code_not_found": "o código fornecido não existe",
		"code_required":  "campo de código é obrigatório",
		"code_expired":   "o código fornecido expirou",

		"trust_ip_field_required": "campo de confiança sessão é obrigatório",

		"mongo: no documents in result": "nenhum registro encontrado",

		"emails_new_question_subject": "Você recebeu uma nova pergunta",
		"emails_new_question_body":    "Você recebeu uma pergunta de @",

		"emails_forgot_password_subject": "Recuperação de senha",
		"emails_forgot_password_body":    "Você solicitou a recuperação de senha, para alterar sua senha, use o código abaixo: ",

		"emails_password_changed_subject": "Senha alterada",
		"emails_password_changed_body":    "Sua senha foi alterada com sucesso. Se você não solicitou a alteração de senha, entre em contato conosco imediatamente",

		"emails_report_sent_subject": "Denúncia enviada",
		"emails_report_sent_body":    "A denúncia que você enviou foi recebida com sucesso. Obrigado por nos ajudar a manter a comunidade segura",
	}
}
