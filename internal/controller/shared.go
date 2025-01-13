package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func RespondJSON[T any](w http.ResponseWriter, r *http.Request, status int, value T) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(value); err != nil {
		return fmt.Errorf("encode json: %v", err)
	}

	return nil
}

func DecodeJSON[T any](r *http.Request) (T, error) {
	var value T
	if err := json.NewDecoder(r.Body).Decode(&value); err != nil {
		return value, fmt.Errorf("decode json: %v", err)
	}

	return value, nil
}
