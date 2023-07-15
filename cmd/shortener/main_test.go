package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http/httptest"
	"testing"
)

func TestRedirectURLHandler(t *testing.T) {
	type want struct {
		code     int
		response string
		location string
		method   string
	}
	tests := []struct {
		name string
		want want
	}{
		{
			name: "positive GET test #1",
			want: want{
				code:     307,
				response: `{"url":"http://localhost:8080/fgdggd"}`,
				location: "/fgdggd",
				method:   "GET",
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			request := httptest.NewRequest(test.want.method, "/fgdggd", nil)

			w := httptest.NewRecorder()
			redirectURLHandler(w, request)

			res := w.Result()

			assert.Equal(t, res.StatusCode, test.want.code)
			defer res.Body.Close()

			resBody, err := io.ReadAll(res.Body)

			require.NoError(t, err)
			assert.JSONEq(t, string(resBody), test.want.response)
		})
	}
}

func TestShortenURLHandler(t *testing.T) {
	type want struct {
		code        int
		response    string
		contentType string
		method      string
	}
	tests := []struct {
		name string
		want want
	}{
		{
			name: "positive POST test",
			want: want{
				code:        201,
				response:    `{"url":"http://localhost:8080/fgdggd"}`,
				contentType: "text/plain",
				method:      "POST",
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			request := httptest.NewRequest(test.want.method, "/fgdggd", nil)

			w := httptest.NewRecorder()
			shortenURLHandler(w, request)

			res := w.Result()

			assert.Equal(t, res.StatusCode, test.want.code)
			defer res.Body.Close()

			resBody, err := io.ReadAll(res.Body)

			require.NoError(t, err)
			assert.JSONEq(t, string(resBody), test.want.response)
		})
	}
}
