//Filename: qapi/cmd/api/helpers.go

package main

import (
	"encoding/json"
	"net/http"
)

// Defining envelope
type envelope map[string]interface {
}

// this function converts a map -> JSON object
func (app *application) writeJSON(w http.ResponseWriter, status int, data envelope, headers http.Header) error {
	//the actual conversion
	js, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	js = append(js, '\n')
	for key, value := range headers {
		w.Header()[key] = value
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)
	return nil
}
