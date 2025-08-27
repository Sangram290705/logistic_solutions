package main

import "net/http"

func (app *application) testHandler(w http.ResponseWriter, r *http.Request) {
	user := app.contextGetUser(r)

	err := app.writeJSON(w, http.StatusOK, envelope{"current_user": user}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
