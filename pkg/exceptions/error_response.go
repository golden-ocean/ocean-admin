package exceptions

//	export interface response {
//		success: boolean; // if request is success
//		data?: any; // response data
//		errorCode?: string; // code for errorType
//		errorMessage?: string; // message display to user
//		showType?: number; // error display type： 0 silent; 1 message.warn; 2 message.error; 4 notification; 9 page
//		traceId?: string; // Convenient for back-end Troubleshooting: unique request ID
//		host?: string; // onvenient for backend Troubleshooting: host of current access server
//	  }
type Error struct {
	Success bool   `json:"success"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	TraceId string `json:"traceId"`
	Host    string `json:"host"`
}

func (e *Error) Error() string {
	return e.Message
}

func DefaultError(c int, m string) *Error {
	return &Error{
		Success: false,
		Code:    c,
		Message: m,
		TraceId: "",
		Host:    "",
	}
}

// 404
func NotFound() *Error {
	return &Error{
		Success: false,
		Code:    404,
		Message: "404 NOT FOUND",
		TraceId: "",
		Host:    "",
	}
}

// 400
func BadRequest() *Error {
	return &Error{
		Success: false,
		Code:    400,
		Message: "400 BAD REQUEST",
		TraceId: "",
		Host:    "",
	}
}

// 403
func Forbidden() *Error {
	return &Error{
		Success: false,
		Code:    403,
		Message: "403 Forbidden",
		TraceId: "",
		Host:    "",
	}
}

// 500
func InternalServerError() *Error {
	return &Error{
		Success: false,
		Code:    500,
		Message: "500 Internal Server Error",
		TraceId: "",
		Host:    "",
	}

}
