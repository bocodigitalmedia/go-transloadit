package transloadit_api

import "strings"

type ErrorResponse struct {
	Code    string `json:"error"`
	Message string `json:"message"`
}

func (e *ErrorResponse) Error() string {
	parts := []string{}

	if e.Code != "" {
		parts = append(parts, e.Code)
	}

	if e.Message != "" {
		parts = append(parts, e.Message)
	}

	return strings.Join(parts, ": ")
}

func (e *ErrorResponse) IsEmpty() bool {
	return e.Code == "" && e.Message == ""
}
