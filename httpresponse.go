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

// WriteJSON serializes v to JSON before writing any response.
// On marshal failure it returns the error without setting status or body.
// It sets Content-Type only when the caller has not already set it.
// A trailing newline is written to match encoding/json.Encoder behavior.
func WriteJSON(w http.ResponseWriter, status int, v any) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}
	b = append(b, '\n')

	if w.Header().Get("Content-Type") == "" {
		w.Header().Set("Content-Type", contentTypeJSON)
	}
	w.WriteHeader(status)
	_, err = w.Write(b)
	return err
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
