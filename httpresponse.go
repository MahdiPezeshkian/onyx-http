// Package httpresponse provides small helpers for writing JSON HTTP responses
// using onyx-response envelopes. Callers supply HTTP status codes explicitly;
// this package does not map domain errors to status codes.
package httpresponse

import (
	"encoding/json"
	"net/http"

	"github.com/MahdiPezeshkian/onyx-response"
)

const contentTypeJSON = "application/json; charset=utf-8"

// WriteJSON writes v as JSON with the given HTTP status code.
// It sets Content-Type to application/json; charset=utf-8.
func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", contentTypeJSON)
	w.WriteHeader(status)
	enc := json.NewEncoder(w)
	return enc.Encode(v)
}

// WriteSuccess writes a successful onyx-response envelope.
func WriteSuccess[T any](w http.ResponseWriter, status int, value *T, message string) error {
	return WriteJSON(w, status, response.Success(value, message))
}

// WriteFailure writes an error onyx-response envelope.
func WriteFailure(w http.ResponseWriter, status int, errorMessage string) error {
	return WriteJSON(w, status, response.Failure(errorMessage))
}

// WritePagedSuccess writes a successful paginated onyx-response envelope.
// statusCodeBody is stored in the JSON body StatusCode field; httpStatus is
// the HTTP response status line.
func WritePagedSuccess[T any](w http.ResponseWriter, httpStatus int, items []T, pageNumber, pageSize, totalCount, statusCodeBody int, message string) error {
	return WriteJSON(w, httpStatus, response.PagedSuccess(items, pageNumber, pageSize, totalCount, statusCodeBody, message))
}

// WritePagedFailure writes a paginated error onyx-response envelope.
func WritePagedFailure[T any](w http.ResponseWriter, httpStatus int, statusCodeBody int, errorMessage string) error {
	return WriteJSON(w, httpStatus, response.PagedFailure[T](statusCodeBody, errorMessage))
}
