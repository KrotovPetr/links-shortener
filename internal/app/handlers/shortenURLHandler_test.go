package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

type URLData struct {
	URL string `json:"url"`
}

func TestShortenURLHandler(t *testing.T) {
	// Подготовка данных
	urlData := URLData{
		URL: "https://example.com",
	}

	jsonData, err := json.Marshal(urlData)
	if err != nil {
		t.Fatal(err)
	}

	reqBody := bytes.NewReader(jsonData)

	req, err := http.NewRequest("POST", "/", reqBody)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	// Вызов обработчика
	ShortenURLHandler(recorder, req)

	// Проверка ответа
	if recorder.Code != http.StatusCreated {
		t.Errorf("Expected status code %d, but got %d", http.StatusCreated, recorder.Code)
	}

	//expectedURL := "http://localhost:8080/abcd1234" // Подставьте сгенерированный короткий URL
	//responseBody := strings.TrimSpace(recorder.Body.String())
	//
	//if responseBody != expectedURL {
	//	t.Errorf("Expected response body %q, but got %q", expectedURL, responseBody)
	//}
}
