package domain

import "errors"

// List of http method supported

const (
	POST = "POST"
	GET  = "GET"
)

// MethodCheck Check for the method selected
func MethodCheck(method string) error {
	switch method {
	case POST, GET:
		return nil
	default:
		return errors.New("method_not_found")
	}
}
