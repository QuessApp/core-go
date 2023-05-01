package locales

// GetSpanishTranslations returns a map with all the translations for the Spanish language.
func GetSpanishTranslations() *map[string]string {
	return &map[string]string{
		"nick_field_required": "campo de nick es obligatorio",
		"nick_field_length":   "campo de nick debe contener entre 3 y 50 caracteres",

		"incorrect_signin_data": "credenciales informadas están incorrectas",

		"password_field_required": "campo de contraseña es obligatorio",
		"password_field_length":   "campo de contraseña debe contener entre 6 a 200 caracteres",

		"name_field_required": "campo de nombre es obligatorio",
		"name_field_length":   "campo de e-mail debe contener entre 3 y 50 caracteres",

		"email_field_required": "campo de e-mail es obligatorio",
		"email_field_length":   "campo de e-mail debe contener entre 5 a 200 caracteres",
		"email_format_invalid": "campo de e-mail tiene formato inválido",
		"email_in_use":         "email ya en uso",

		"nick_in_use": "nombre de usuario ya en uso",

		"logout_from_all_devices_required": "campo de logout de todos los dispositivos es obligatorio",
		"logout_from_all_devices_invalid":  "campo de logout de todos los dispositivos tiene valor inválido",

		"content_required": "campo de contenido es obligatorio",
		"content_length":   "campo de contenido debe contener entre 1 y 250 caracteres",

		"send_to_required": "campo de enviar para es obligatorio",
		"send_to_length":   "campo de enviar para debe contener entre 3 y 50 caracteres",

		"max_file_size":     "por favor, envie una imagen con tamaño menor que 1MB",
		"file_type_invalid": "tipo de archivo informado no es permitido",

		"locale_field_required": "campo de localidad es obligatorio",
		"locale_field_length":   "campo de localidad debe contener entre 5 y 10 caracteres",

		"enable_app_emails_field_required":        "campo de habilitar e-mails de la app es obligatorio",
		"enable_app_notifications_field_required": "campo de habilitar notificaciones de la app es obligatorio",

		"reason_field_required": "campo de razón es obligatorio",
		"reason_field_invalid":  "campo de razón tiene valor inválido",

		"type_field_required": "campo de tipo es obligatorio",
		"type_field_invalid":  "campo de tipo tiene valor inválido",

		"user_not_found":           "no se encontró el usuario con los datos dados",
		"user_to_block_required":   "campo de usuario a bloquear es obligatorio",
		"user_to_unblock_required": "campo de usuario a desbloquear es obligatorio",

		"question_not_found":       "pregunta no encontrada",
		"question_not_sent_for_me": "esta pregunta no fue enviada para usted",

		"question_not_authorized":  "esta pregunta no fue enviada para usted",
		"question_already_replied": "esta pregunta ya fue respondida",

		"reached_questions_limit": "llegó al límite de preguntas enviadas",

		"cant_delete_question_not_sent_by_you": "no puede eliminar una pregunta que no fue enviada por usted",
		"sending_question_to_yourself":         "no puede enviar una pregunta a usted mismo",

		"cant_hide_already_hidden": "ya ocultó esta pregunta",
		"cant_send_invalid_id":     "no puede enviar una pregunta a un usuario con id inválido",

		"cant_edit_reply_not_replied_yet": "no puedes editar esta respuesta porque aún no ha sido respondida",
		"cant_edit_reply_reached_limit":   "no puedes editar esta respuesta porque alcanzaste el límite de ediciones, es decir, cinco",

		"did_blocked_receiver":     "bloqueaste al usuario al que intentas enviar una pregunta",
		"blocked_by_receiver":      "no puedes enviar una pregunta a este usuario porque te bloqueó",
		"already_blocked":          "no puedes bloquear a este usuario porque ya está bloqueado",
		"cant_block_yourself":      "no puedes bloquear a ti mismo",
		"cant_unblock_not_blocked": "no puedes desbloquear a un usuario que no está bloqueado",

		"report_not_found":                   "la denuncia con el id informado no existe",
		"report_not_authorized":              "no tienes acceso a este informe",
		"cant_delete_report_not_sent_by_you": "no puede eliminar este informe porque no lo envió usted",
		"cant_report_already_sent":           "no puede enviar este informe porque ya ha enviado uno similar antes",
		"cant_report_yourself":               "no puede enviar un informe a sí mismo",

		"code_not_found": "el código proporcionado no existe",
		"code_required":  "el campo de código es obligatorio",
		"code_expired":   "el código proporcionado ha caducado",

		"trust_ip_field_required": "el campo de confianza de la sesión es obligatorio",

		"mongo: no documents in result": "no se encontraron registros",

		"emails_new_question_subject": "Tienes una nueva pregunta",
		"emails_new_question_body":    "Recibiste una pregunta de @",

		"emails_forgot_password_subject": "Recuperación de contraseña",
		"emails_forgot_password_body":    "Solicitó la recuperación de contraseña, para cambiar su contraseña, use el siguiente código:",

		"emails_password_changed_subject": "contraseña cambiada",
		"emails_password_changed_body":    "Tu contraseña ha sido cambiada exitosamente. Si no ha solicitado un cambio de contraseña, contáctenos de inmediato.",

		"emails_report_sent_subject": "Reclamación enviada",
		"emails_report_sent_body":    "El informe que envió se recibió con éxito. Gracias por ayudarnos a mantener segura a la comunidad.",

		"emails_unkown_login_attempt_subject": "Intento de inicio de sesión desconocido",
		"emails_unkown_login_attempt_body":    "Notamos que alguien intentó iniciar sesión en su cuenta, si no solicitó este cambio, contáctenos de inmediato. La ubicación del intento de inicio de sesión fue: ",

		"max_rate_limit": "no puede realizar esta acción porque alcanzó el límite de solicitudes, intente nuevamente más tarde",
	}
}
