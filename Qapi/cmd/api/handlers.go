//Filename: qapi/cmd/api/handlers.go

package main

import (
	"net/http"
)

// This handler facilitates the generation of a random string
func (app *application) makeRandomStringHandler(w http.ResponseWriter, r *http.Request) {
	//getting the int supplied using a helper function
	userInt, err := app.readIntParam(r)
	if err != nil {
		app.logError(r, err)
		return
	}

	//stopping the server from crash due to large number

	//The int valid
	randomString := app.generateRandomString(int(userInt))

	//Returning the randomstring
	err = app.writeJSON(w, http.StatusOK, envelope{"Random String": randomString}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
