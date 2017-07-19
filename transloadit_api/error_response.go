package transloadit_api

import "strings"

type ErrorResponse struct {
	Code    string `json:"error,omitempty"`
	Message string `json:"message,omitempty"`
	Reason  string `json:"reason,omitempty"`
}

func (e ErrorResponse) Error() string {
	parts := []string{}

	if e.Code != "" {
		parts = append(parts, e.Code)
	}

	if e.Message != "" {
		parts = append(parts, e.Message)
	}

	if e.Reason != "" {
		parts = append(parts, e.Reason)
	}

	return strings.Join(parts, ": ")
}

func (e *ErrorResponse) IsEmpty() bool {
	return e.Code == ""
}
