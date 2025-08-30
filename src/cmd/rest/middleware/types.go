package middleware

import "net/http"

type HandlerFunc func(w http.ResponseWriter, r *http.Request) error

type ErrorResponse struct {
	Error string `json:"error"`
}
