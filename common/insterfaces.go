package common

import "go-template-api/internal/model"

type SqlQueryStatus struct {
	Message string
	Code    int
	Err     error
}

type LoginResponse struct {
	Message string `json:"message"`
	Token   string `json:"Token"`
}

type ReisterResponse struct {
	Message string           `json:"message"`
	User    model.User_Table `json:"user"`
}
