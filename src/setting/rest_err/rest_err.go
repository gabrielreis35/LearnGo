package rest_err

import "net/http"

type RestErr struct {
	Message string   `json:"message"`
	Err     string   `json:"error"`
	Code    int      `json:"code"`
	Causes  []Causes `json:"causes"`
}

type Causes struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (r *RestErr) Error() string {
	return r.Message
}

func NewRestErr(message string, err string, code int, causes Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err:     err,
		Code:    code,
		Causes:  []Causes{causes},
	}
}

func NewBadRequestError(message string, causes []Causes) *RestErr {
	var causesList []Causes
	if len(causes) > 0 {
		causesList = causes
	} else {
		causesList = nil // causes será nil por padrão se não houver argumentos
	}

	return &RestErr{
		Message: message,
		Err:     "bad request",
		Code:    http.StatusBadRequest,
		Causes:  causesList,
	}
}

func NewInternalServerError(message string, causes Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "internal server error",
		Code:    http.StatusInternalServerError,
	}
}

func NewNotFoundError(message string, causes Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "not found",
		Code:    http.StatusNotFound,
	}
}

func NewForbiddenError(message string, causes Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "forbidden",
		Code:    http.StatusForbidden,
	}
}
