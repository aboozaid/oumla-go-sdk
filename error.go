package oumlagosdk

type Error struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type OumlaError struct {
	ErrorType string
	Status    int
	Message   string `json:"message"`
	Path      string
	Errors    *[]Error `json:"errors"`
	// Details   *any
}

func (e *OumlaError) SetErrorType(status int) {
	switch status {
	case 400, 402:
		e.ErrorType = "ValidationError"
	case 401, 403:
		e.ErrorType = "AuthenticationError"
	case 404:
		e.ErrorType = "NotFoundError"
	case 500:
		e.ErrorType = "ServerError"
	default:
		e.ErrorType = "UnknownError"
	}
}

// func (e *OumlaError) SetDetails(details *any) {
// 	e.Details = details
// }

func (e *OumlaError) Error() string {
	return e.Message
}