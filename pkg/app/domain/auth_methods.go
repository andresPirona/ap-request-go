package domain

import (
	"errors"
)

type AuthMethod struct {
	Name AMethod     `json:"name"`
	Data interface{} `json:"data"`
}

type A string

// Types of auth supported

const(
	Basic A = "Basic"
)

type AMethod struct {
	A
}

// Check the method
func (a *AuthMethod) Check() error {

	if a != nil && a.Data == nil {
		return errors.New("auth_method_without_properties")
	}

	switch a.GetName() {
	case string(Basic):
		return nil
	default:
		return errors.New("auth_method_not_found")
	}


}

// Getter for the authentication method
func (a *AuthMethod) GetName() string {
	return string(a.Name.A)
}
