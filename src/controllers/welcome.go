package controllers

import (
	"net/http"
	"encoding/json"
)

type Response struct {
    Message string
}

func WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	data := Response{"Why Hello There. Go rocks!"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
