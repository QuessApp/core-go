package errors

const (
	NICK_FIELD_REQUIRED = "nick_field_required"
	NICK_FIELD_LENGTH   = "nick_field_length"

	INCORRECT_SIGNIN_DATA = "incorrect_signin_data"

	PASSWORD_FIELD_REQUIRED = "password_field_required"
	PASSWORD_FIELD_LENGTH   = "password_field_length"

	NAME_FIELD_REQUIRED = "name_field_required"
	NAME_FIELD_LENGTH   = "name_field_length"

	EMAIL_FIELD_REQUIRED = "email_field_required"
	EMAIL_FIELD_LENGTH   = "email_field_length"
	EMAIL_FORMAT_INVALID = "email_format_invalid"
	EMAIL_IN_USE         = "email_in_use"

	NICK_IN_USE = "nick_in_use"

	LOGOUT_FROM_ALL_DEVICES_REQUIRED = "logout_from_all_devices_required"
	LOGOUT_FROM_ALL_DEVICES_INVALID  = "logout_from_all_devices_invalid"

	CONTENT_REQUIRED = "content_field_required"
	CONTENT_LENGTH   = "content_field_length"

	SEND_TO_REQUIRED = "send_to_field_required"
	SEND_TO_LENGTH   = "send_to_field_length"

	MAX_FILE_SIZE     = "max_file_size"
	FILE_TYPE_INVALID = "file_type_invalid"

	LOCALE_FIELD_REQUIRED = "locale_field_required"
	LOCALE_FIELD_INVALID  = "locale_field_invalid"

	ENABLE_APP_EMAILS_FIELD_REQUIRED        = "enable_app_emails_field_required"
	ENABLE_APP_NOTIFICATIONS_FIELD_REQUIRED = "enable_app_notifications_field_required"

	REASON_FIELD_REQUIRED = "reason_field_required"
	REASON_FIELD_INVALID  = "reason_field_invalid"

	TYPE_FIELD_REQUIRED = "type_field_required"
	TYPE_FIELD_INVALID  = "type_field_invalid"
)

const (
	USER_NOT_FOUND           = "user_not_found"
	USER_TO_BLOCK_REQUIRED   = "user_to_block_required"
	USER_TO_UNBLOCK_REQUIRED = "user_to_unblock_required"
)

const (
	QUESTION_NOT_FOUND       = "question_not_found"
	QUESTION_NOT_SENT_FOR_ME = "question_not_sent_for_me"
	QUESTION_NOT_AUTHORIZED  = "question_not_authorized"
	QUESTION_ALREADY_REPLIED = "question_already_replied"

	REACHED_QUESTIONS_LIMIT = "reached_questions_limit"

	CANT_DELETE_QUESTION_NOT_SENT_BY_YOU = "cant_delete_question_not_sent_by_you"
	SENDING_QUESTION_TO_YOURSELF         = "sending_question_to_yourself"

	CANT_HIDE_ALREADY_HIDDEN = "cant_hide_already_hidden"
	CANT_SEND_INVALID_ID     = "cant_send_invalid_id"
)

const (
	CANT_EDIT_REPLY_NOT_REPLIED_YET = "cant_edit_reply_not_replied_yet"
	CANT_EDIT_REPLY_REACHED_LIMIT   = "cant_edit_reply_reached_limit"
)

const (
	DID_BLOCKED_RECEIVER     = "did_blocked_receiver"
	BLOCKED_BY_RECEIVER      = "blocked_by_receiver"
	ALREADY_BLOCKED          = "already_blocked"
	CANT_BLOCK_YOURSELF      = "cant_block_yourself"
	CANT_UNBLOCK_NOT_BLOCKED = "cant_unblock_not_blocked"
)

const (
	REPORT_NOT_FOUND                   = "report_not_found"
	REPORT_NOT_AUTHORIZED              = "report_not_authorized"
	CANT_DELETE_REPORT_NOT_SENT_BY_YOU = "cant_delete_report_not_sent_by_you"
	CANT_REPORT_ALREADY_SENT           = "cant_report_already_sent"
	CANT_REPORT_YOURSELF               = "cant_report_yourself"
)

const (
	TOKEN_NOT_FOUND = "token_not_found"
	TOKEN_EXPIRED   = "token_expired"
)

const (
	CODE_NOT_FOUND = "code_not_found"
	CODE_REQUIRED  = "code_required"
	CODE_EXPIRED   = "code_expired"
)

const (
	TRUST_IP_FIELD_REQUIRED = "trust_ip_field_required"
)
