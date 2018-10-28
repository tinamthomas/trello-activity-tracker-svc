package controllers

import (
    "testing"
    "net/http"
    "net/http/httptest"
    "github.com/stretchr/testify/assert"

    "encoding/json"
)

func TestWelcomeHandler(t *testing.T) {
    req, _ := http.NewRequest("GET", "/health-check", nil)
    expectedResponse := Response{"Why Hello There. Go rocks!"}

    w := httptest.NewRecorder()
    handler := http.HandlerFunc(WelcomeHandler)

    handler.ServeHTTP(w, req)

    actualResponse := Response{}
    err := json.Unmarshal([]byte(w.Body.String()), &actualResponse)

    assert.Nil(t, err)
    assert.Equal(t, 200, w.Code)
    assert.Equal(t, expectedResponse, actualResponse)
}