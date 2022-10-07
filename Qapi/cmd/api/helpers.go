//Filename: qapi/cmd/api/helpers.go

package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
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

// this functions allows how to read what was requested of the api
func (app *application) readIDParam(r *http.Request) (int64, error) {
	//getting the data from a slice
	params := httprouter.ParamsFromContext(r.Context())

	//getting the int supplied
	userInt, err := strconv.ParseInt(params.ByName("seed:"), 10, 64)
	if err != nil {
		return 0, errors.New("invalid seed parameter, ensure that your seed is a number")
	}
	return userInt, nil
}
