package errors

import "net/http"

type ErrorResponse struct {
	statusCode int    `json:"-"`
	Code       string `json:"code"`
	Message    string `json:"message"`
	Resource   string `json:"resource,omitempty"`
	RequestID  string `json:"request_id"`
	RawError   string `json:"error"`
}

func NewErrorResponse(requestID, resource string, err Error, errs ...error) *ErrorResponse {
	er := &ErrorResponse{
		statusCode: err.Code,
		Code:       err.Name,
		Message:    err.Message,
		Resource:   resource,
		RequestID:  requestID,
	}
	if len(errs) > 0 {
		er.RawError = errs[0].Error()
	}

	return er
}

// implements gogo.StatusCoder interface
func (er *ErrorResponse) StatusCode() int {
	return er.statusCode
}

func (er *ErrorResponse) Error() string {
	if er.RawError != "" {
		return "[" + er.Code + "] " + er.RawError
	}

	return "[" + er.Code + "] " + er.Message
}

type Response struct {
	statusCode int         `json:"-"`
	Code       string      `json:"code"`
	Message    string      `json:"message"`
	Resource   string      `json:"resource,omitempty"`
	RequestId  string      `json:"request_id,omitempty"`
	Data       interface{} `json:"data"`
}

func NewResponse(requestID, resource string, data interface{}, statusCode ...int) *Response {
	resp := &Response{
		statusCode: http.StatusOK,
		Resource:   resource,
		RequestId:  requestID,
		Data:       data,
	}

	if len(statusCode) > 0 {
		resp.statusCode = statusCode[0]
	}

	return resp
}
