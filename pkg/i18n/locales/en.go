package locales

// GetAmericanEnglishTranslations returns a map with all the american english translations.
func GetAmericanEnglishTranslations() *map[string]string {
	return &map[string]string{
		"nick_field_required": "nick field is required",
		"nick_field_length":   "nick field must contain between 3 and 50 characters",

		"incorrect_signin_data": "credentials provided are incorrect",

		"password_field_required": "password field is required",
		"password_field_length":   "password field must contain between 6 and 200 characters",

		"name_field_required": "name field is required",
		"name_field_length":   "name field must contain between 3 and 50 characters",

		"email_field_required": "email field is required",
		"email_field_length":   "email field must contain between 5 and 200 characters",
		"email_format_invalid": "email field has invalid format",
		"email_in_use":         "email already in use",

		"nick_in_use": "nick already in use",

		"logout_from_all_devices_required": "logout from all devices field is required",
		"logout_from_all_devices_invalid":  "logout from all devices field has invalid value",

		"content_required": "content field is required",
		"content_length":   "content field must contain between 1 and 250 characters",

		"send_to_required": "send to field is required",
		"send_to_length":   "send to field must contain between 3 and 50 characters",

		"max_file_size":     "please, send image with size smaller than 1MB",
		"file_type_invalid": "file type informed is not allowed",

		"locale_field_required": "locale field is required",
		"locale_field_invalid":  "locale field has invalid value",

		"enable_app_emails_field_required":        "enable app emails field is required",
		"enable_app_notifications_field_required": "enable app notifications field is required",

		"reason_field_required": "reason field is required",
		"reason_field_invalid":  "reason field has invalid value",

		"type_field_required": "type field is required",
		"type_field_invalid":  "type field has invalid value",

		"user_not_found":           "user not found",
		"user_to_block_required":   "user to block field is required",
		"user_to_unblock_required": "user to unblock field is required",

		"question_not_found":       "question not found",
		"question_not_sent_for_me": "you don't have access to this question",

		"question_not_authorized":  "you don't have access to this question",
		"question_already_replied": "you can't reply this question because it was already replied",

		"reached_questions_limit": "you can't send more questions because you reached the limit of questions sent",

		"cant_delete_question_not_sent_by_you": "you can't delete this question because it was not sent by you",
		"sending_question_to_yourself":         "you can't send a question to yourself",

		"can_hide_already_hidden": "you can't hide this question because it was already hidden",
		"cant_send_invalid_Id":    "you can't send a question to an invalid user",

		"cant_edit_reply_not_replied_yet": "you can't edit this reply because it was not replied yet",
		"cant_edit_reply_reached_limit":   "you can't edit this reply because you reached the limit of edits, that is five",

		"did_blocked_receiver":     "you blocked the user that you are trying to send a question to",
		"blocked_by_receiver":      "you can't send a question to this user because he blocked you",
		"already_blocked":          "you can't block this user because he is already blocked",
		"cant_block_yourself":      "you can't block yourself",
		"cant_unblock_not_blocked": "you can't unblock this user because he is not blocked",

		"report_not_found":                   "report not found",
		"report_not_authorized":              "you don't have access to this report",
		"cant_delete_report_not_sent_by_you": "you can't delete this report because it was not sent by you",
		"cant_report_already_sent":           "you can't send a report because you already sent one",
		"cant_report_yourself":               "you can't send a report to yourself",

		"code_not_found": "the provided code does not exist",
		"code_required":  "code field is required",
		"code_expired":   "the provided code has expired",

		"trust_ip_field_required": "trust ip field is required",

		"mongo: no documents in result": "no record found with the provided data",

		"emails_new_question_subject": "You just received a new question",
		"emails_new_question_body":    "You just received a new question from @",

		"emails_forgot_password_subject": "Recovery password",
		"emails_forgot_password_body":    "You requested a password recovery, to do this, use the following code: ",

		"emails_password_changed_subject": "Password changed",
		"emails_password_changed_body":    "Your password was changed, if you did not request this change, please contact us immediately",

		"emails_report_sent_subject": "Report sent",
		"emails_report_sent_body":    "Thank you for sending a report, we will analyze it and take the necessary actions",
	}
}
