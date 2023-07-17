package handlers

import (
	"github.com/KrotovPetr/links-shortener.git/internal/app/storage"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPositiveRedirectURLHandler(t *testing.T) {
	storage.URLMap = map[string]string{
		"abcd1234": "https://example.com",
	}

	req, err := http.NewRequest("GET", "/abcd1234", nil)
	if err != nil {
		t.Fatal(err)
	}

	newRecorder := httptest.NewRecorder()

	RedirectURLHandler(newRecorder, req)

	if newRecorder.Code != http.StatusTemporaryRedirect {
		t.Errorf("Expected status code %d, but got %d", http.StatusTemporaryRedirect, newRecorder.Code)
	}

	expectedLocation := "https://example.com"
	location := newRecorder.Header().Get("Location")

	if location != expectedLocation {
		t.Errorf("Expected Location header %q, but got %q", expectedLocation, location)
	}

}

func TestNegativeHandleGetRequestInvalidURL(t *testing.T) {
	storage.URLMap = map[string]string{
		"abcd1234": "https://example.com",
	}
	req, err := http.NewRequest("POST", "/test2", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(RedirectURLHandler)

	handler.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusBadRequest {
		t.Errorf("handleGetRequest returned wrong status code: got %v want %v", recorder.Code, http.StatusBadRequest)
	}
}
