package helper

import (
	"encoding/json"
	"errors"
	"net/http"
)

var (
	ErrBadInputRequest    = errors.New("provide required fields")
	ErrDBUpdateFailed     = errors.New("database update get failed")
	ErrCredentialNotMatch = errors.New("credential not matched")
	ErrTicketIDMissing    = errors.New("ticket id should be mention")
)

// error helper response function
func ErrResponse(w http.ResponseWriter, code int, msg error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	data := map[string]any{"error": msg.Error()}
	if err := json.NewEncoder(w).Encode(data); err != nil {
		w.Write([]byte(msg.Error()))
	}
}

// success helper response function
func SuccResponse(w http.ResponseWriter, code int, msg any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(msg); err != nil {
		w.Write([]byte("request update successfully"))
	}
}
