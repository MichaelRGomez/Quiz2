//Filename: qapi/cmd/api/handlers.go

package main

import (
	"net/http"
)

// This handler facilitates the generation of a random string
func (app *application) makeRandomStringHandler(w http.ResponseWriter, r *http.Request) {
	//getting the int supplied using a helper function
	userInt, err := app.readIDParam(r)
	if err != nil {
		app.failedToParseInt(w, r, err)
		return
	}

	//The int valid
	randomString := Tools.generateRandomString(userInt)
}
