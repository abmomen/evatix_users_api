package errors

import "net/http"

//RestErr defines the structure of Error in out applicatons
type RestErr struct {
	Message string `json:"message"`
	Status  int    `josn:"status"`
	Error   string `json:"error"`
}

//NewBadRequestError create a bad request error with the provided message.
func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
	}
}

//NewNotFoundError error if user not found
func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "not_found",
	}
}

//NewNotFoundError error if user not found
func NewInternalServerError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusInternalServerError,
		Error:   "Internal Server Error",
	}
}