package api

import (
	"net/http"
)

func (server *API) GetTest(w http.ResponseWriter, r *http.Request) error {
	return writeJSON(w, http.StatusOK, "Hello World")
}
