package presenter

import (
	domain2 "github.com/micheltank/crypto-price-alert/price-alert-service/internal/domain"
)

type ApiError struct {
	Message string `json:"message"`
	Key     string `json:"key"`
	Detail  string `json:"detail"`
}

func NewApiError(errorDomain domain2.ErrorDomain) ApiError {
	return ApiError{
		Message: errorDomain.Error(),
		Key:     errorDomain.Key(),
		Detail:  errorDomain.Detail(),
	}
}