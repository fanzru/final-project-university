package response

const (
	InternalServerName        = "internal_server_error"
	InternalServerDescription = "The server is unable to complete your request"
)

const (
	UnauthorizedErrorName        = "access_denied"
	UnauthorizedErrorDescription = "Authorization failed by filter."
)

const (
	BadRequestErrorName        = "bad_request"
	BadRequestErrorDescription = "Your request resulted in error"
)

const (
	SuccessCreatedName = "success_created"
	SuccessOK          = "success_ok"
	SuccessLogin       = "success_login"
)
