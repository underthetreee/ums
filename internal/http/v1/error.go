package v1

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"error"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("[%d] %s", e.Code, e.Message)
}

func ErrInvalidInput() *Error {
	return &Error{Code: http.StatusBadRequest, Message: "invalid input"}
}

func ErrInternalServer() *Error {
	return &Error{Code: http.StatusInternalServerError, Message: "internal server error"}
}

func ErrUnauthorized() *Error {
	return &Error{Code: http.StatusUnauthorized, Message: "unauthorized"}
}

func JSONError(w http.ResponseWriter, resp *Error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.Code)

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "error encoding json", http.StatusInternalServerError)
		return

	}
}
