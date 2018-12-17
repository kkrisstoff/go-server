package app_error

import (
	"encoding/json"
	"net/http"

	"go.uber.org/zap"
)

const status = "error"

var logger *zap.Logger

// AppError is an error with some additional context for http requests
type AppError struct {
	Status string `json:"status"`
	Cause  string `json:"errorMessage"`
	Code   int    `json:"errorCode"`
}

// New creates a new AppError
func New(cause error, code int) error {
	return &AppError{
		Status: status,
		Cause:  cause.Error(),
		Code:   code,
	}
}

func (ae *AppError) Error() string {
	return ae.Cause
}

func WriteError(w http.ResponseWriter, err error) {
	logger.Error("Got http error", zap.Error(err))
	switch myError := err.(type) {
	case *AppError:
		w.WriteHeader(myError.Code)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	default:
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
