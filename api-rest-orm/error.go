package server

import "fmt"

type ErrorApi struct {
	Code    string `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

var (
	ErrEntityAlreadyExists = fmt.Errorf("entity already exists")
	ErrEntityIDNotProvided = fmt.Errorf("entity id not provided")
	ErrEntityNotFound      = fmt.Errorf("entity not found")
	ErrEntityOutOfRange    = fmt.Errorf("entity out of range")
)
