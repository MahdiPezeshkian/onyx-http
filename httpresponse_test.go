package httpresponse_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/MahdiPezeshkian/onyx-http"
	"github.com/MahdiPezeshkian/onyx-response"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func TestWriteJSON_SetsHeadersAndBody(t *testing.T) {
	rr := httptest.NewRecorder()
	payload := map[string]string{"hello": "world"}
	if err := httpresponse.WriteJSON(rr, http.StatusCreated, payload); err != nil {
		t.Fatalf("WriteJSON: %v", err)
	}
	if rr.Code != http.StatusCreated {
		t.Fatalf("status: want %d got %d", http.StatusCreated, rr.Code)
	}
	if ct := rr.Header().Get("Content-Type"); ct != "application/json; charset=utf-8" {
		t.Fatalf("content-type: got %q", ct)
	}
	var got map[string]string
	if err := json.NewDecoder(rr.Body).Decode(&got); err != nil {
		t.Fatalf("decode: %v", err)
	}
	if got["hello"] != "world" {
		t.Fatalf("body: %#v", got)
	}
}

func TestWriteSuccess(t *testing.T) {
	rr := httptest.NewRecorder()
	user := &User{ID: 1, Name: "Ada"}
	if err := httpresponse.WriteSuccess(rr, http.StatusOK, user, "ok"); err != nil {
		t.Fatalf("WriteSuccess: %v", err)
	}
	var got response.Response[User]
	if err := json.NewDecoder(rr.Body).Decode(&got); err != nil {
		t.Fatalf("decode: %v", err)
	}
	if got.HasError || got.Message != "ok" || got.Value == nil || got.Value.Name != "Ada" {
		t.Fatalf("unexpected envelope: %#v", got)
	}
}

func TestWriteFailure(t *testing.T) {
	rr := httptest.NewRecorder()
	if err := httpresponse.WriteFailure(rr, http.StatusBadRequest, "bad input"); err != nil {
		t.Fatalf("WriteFailure: %v", err)
	}
	if rr.Code != http.StatusBadRequest {
		t.Fatalf("status: want %d got %d", http.StatusBadRequest, rr.Code)
	}
	var got response.Response[any]
	if err := json.NewDecoder(rr.Body).Decode(&got); err != nil {
		t.Fatalf("decode: %v", err)
	}
	if !got.HasError || got.ErrorMessage != "bad input" || got.Value != nil {
		t.Fatalf("unexpected envelope: %#v", got)
	}
}

func TestWritePagedSuccess(t *testing.T) {
	rr := httptest.NewRecorder()
	items := []User{{ID: 1, Name: "Ada"}}
	if err := httpresponse.WritePagedSuccess(rr, http.StatusOK, items, 1, 10, 25, 200, "ok"); err != nil {
		t.Fatalf("WritePagedSuccess: %v", err)
	}
	var got response.PagedResponse[User]
	if err := json.NewDecoder(rr.Body).Decode(&got); err != nil {
		t.Fatalf("decode: %v", err)
	}
	if got.HasError || !got.HasNextPage || got.StatusCode != 200 || len(got.Items) != 1 {
		t.Fatalf("unexpected envelope: %#v", got)
	}
}

func TestWritePagedFailure(t *testing.T) {
	rr := httptest.NewRecorder()
	if err := httpresponse.WritePagedFailure[User](rr, http.StatusNotFound, 404, "missing"); err != nil {
		t.Fatalf("WritePagedFailure: %v", err)
	}
	if rr.Code != http.StatusNotFound {
		t.Fatalf("status: want %d got %d", http.StatusNotFound, rr.Code)
	}
	var got response.PagedResponse[User]
	if err := json.NewDecoder(rr.Body).Decode(&got); err != nil {
		t.Fatalf("decode: %v", err)
	}
	if !got.HasError || got.ErrorMessage != "missing" || got.StatusCode != 404 {
		t.Fatalf("unexpected envelope: %#v", got)
	}
}
