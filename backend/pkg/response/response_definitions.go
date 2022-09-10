package response

//
// An ErrorMessage represents an error message format and list of error.
//
type ErrorMessage struct {
	Code  int          `json:"code"`
	Error *ErrorFormat `json:"error"`
}

//
// An ErrorFormat represents an error message format and code that we used.
//
type ErrorFormat struct {
	ErrorName        string `json:"error_name"`
	ErrorDescription string `json:"error_description"`
}

type SuccessCreatedMessage struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
