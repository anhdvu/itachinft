package main

import (
	"encoding/json"
	"net/http"
)

type capsule map[string]any

func SendJSON(w http.ResponseWriter, statuscode int, headers http.Header, content capsule) error {
	payload, err := json.Marshal(content)
	if err != nil {
		return err
	}
	payload = append(payload, '\n')

	for key, value := range headers {
		w.Header()[key] = value
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(statuscode)
	w.Write(payload)
	return nil
}
