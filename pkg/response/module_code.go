package response

// Format: NewError(httpStatusCode, errorType, errorDescription)
var (
	DisplayNameInvalidError  = NewError(400, "DISPLAY_NAME_INVALID", "Display name is invalid")
	UsernameInvalidError     = NewError(400, "USERNAME_INVALID", "Username is invalid")
	BioInvalidError          = NewError(400, "BIO_INVALID", "Bio is invalid")
	UsernameUnavailableError = NewError(400, "USERNAME_UNAVAILABLE", "Username not available")
	UserNotExistsError       = NewError(400, "USER_NOT_EXIST", "User does not exists")
	AccountExistsError       = NewError(400, "ACCOUNT_EXIST", "An account for this user already exists")
	InvalidRequestError      = NewError(400, "INVALID_REQUEST", "Invalid request")

	ForbiddenError = NewError(403, "FORBIDDEN")

	NotFoundError = NewError(404, "NOT_FOUND", "Resource not found")

	InternalError = NewError(500, "INTERNAL", "Please try again")

	// Media Related Errors
	// MediaUploadFailed    = NewError(400, "90000", "Media upload failed")
	// MediaInvalidSize     = NewError(400, "90001", "Invalid File Size")
	// MediaUpdateFailed    = NewError(400, "90002", "Failed to update the media details")
	// MediaNotFound        = NewError(400, "90003", "Media not found")
	// MediaTypeUnsupported = NewError(400, "90004", "Media type is not supported")
	// MediaDeleteFailed    = NewError(400, "90005", "Media delete failed")

	// PasswordLengthLimit     = NewError(20004, "Password length 6~16")
	// UserHasBeenBanned       = NewError(20006, "This account has been blocked")
	// NoPermission            = NewError(20007, "No permission to execute this request")
	// UserHasBindOTP          = NewError(20008, "The current user has bound the secondary authentication")
	// UserOTPInvalid          = NewError(20009, "Second Captcha failed")
	// UserNoBindOTP           = NewError(20010, "The current user is not bound to secondary authentication")
	// ErrorOldPassword        = NewError(20011, "Current user password authentication failed")
	// ErrorCaptchaPassword    = NewError(20012, "Captcha failed to verify")
	// AccountNoPhoneBind      = NewError(20013, "Rejected Operation: Account not bound to mobile phone number")
	// TooManyLoginError       = NewError(20014, "Too many login failures, please try again later")
	// GetPhoneCaptchaError    = NewError(20015, "SMS Captcha failed to get")
	// TooManyPhoneCaptchaSend = NewError(20016, "The number of SMS verification codes obtained has reached today's limit")
	// ExistedUserPhone        = NewError(20017, "The phone number has been bound")
	// ErrorPhoneCaptcha       = NewError(20018, "Incorrect phone captcha")
	// MaxPhoneCaptchaUseTimes = NewError(20019, "The mobile phone verification code has reached the maximum number of times")
	// NicknameLengthLimit     = NewError(20020, "nickname length 2~12")
	// NoAdminPermission       = NewError(20022, "No administrative privileges")
	// NoExistUsername         = NewError(20021, "User does not exist")

	// GetPostsFailed          = NewError(30001, "Failed  to get dynamic list")
	// CreatePostFailed        = NewError(30002, "Dynamic publish failed")
	// GetPostFailed           = NewError(30003, "Failed  to get dynamic details")
	// DeletePostFailed        = NewError(30004, "Dynamic delete failed")
	// LockPostFailed          = NewError(30005, "Dynamic lock failed")
	// GetPostTagsFailed       = NewError(30006, "Failed  to get list of topics")
	// InvalidDownloadReq      = NewError(30007, "Attachment download request is not legitimate")
	// DownloadReqError        = NewError(30008, "Attachment download request failed")
	// InsuffientDownloadMoney = NewError(30009, "Attachment download failed: Insufficient account funds")
	// DownloadExecFail        = NewError(30010, "Attachment download failed: Charge failed")
	// StickPostFailed         = NewError(30011, "Dynamic pin failed")
	// VisblePostFailed        = NewError(30012, "Failed to update visibility")

	// GetCommentsFailed   = NewError(40001, "Failed to get list of comments")
	// CreateCommentFailed = NewError(40002, "Comment post failed")
	// GetCommentFailed    = NewError(40003, "Failed to get comment details")
	// DeleteCommentFailed = NewError(40004, "Comment deletion failed")
	// CreateReplyFailed   = NewError(40005, "Comment reply failed")
	// GetReplyFailed      = NewError(40006, "Failed to get comment details")
	// MaxCommentCount     = NewError(40007, "Maximum number of comments reached")

	// GetMessagesFailed = NewError(50001, "Failed to get message list")
	// ReadMessageFailed = NewError(50002, "Failed to mark message read")
	// SendWhisperFailed = NewError(50003, "Private message failed")
	// NoWhisperToSelf   = NewError(50004, "Not allowed to send private messages to yourself")
	// TooManyWhisperNum = NewError(50005, "The number of private messages today has reached its upper limit")

	// GetCollectionsFailed = NewError(60001, "Failed to get favorites list")
	// GetStarsFailed       = NewError(60002, "Failed to get like list")

)
