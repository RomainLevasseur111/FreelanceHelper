package api

import (
	"encoding/json"
	"log"
	"net/http"

	"freelancehelper/api/database"

	_ "github.com/mattn/go-sqlite3"
)

type API struct {
	http.Server
	Storage *database.SQLite3Store
	Sessions *database.SessionStore


}

type APIerror struct {
	Status  int    `json:"status"`
	Error   string `json:"error"`
	Message string `json:"message"`
}

type handlerFunc func(http.ResponseWriter, *http.Request) error

func writeJSON(writer http.ResponseWriter, statusCode int, v any) error {
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(statusCode)
	return json.NewEncoder(writer).Encode(v)
}

func handleFunc(fn handlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := fn(w, r); err != nil {
			log.Println(err)
			writeJSON(w, http.StatusInternalServerError,
				APIerror{
					Status:  http.StatusInternalServerError,
					Error:   "Internal Server Error",
					Message: err.Error(),
				})
		}
	}
}

func NewAPI(addr, dbFilePath string) (*API, error) {
	server := new(API)
	server.Server.Addr = addr

	router := http.NewServeMux()

	// GET routes
	//router.HandleFunc("/test", handleFunc(server.GetTest))

	// POST routes
	router.HandleFunc("/register", handleFunc(server.Register))

	// Server handling routes
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "dist/index.html")
	})
	router.Handle("/assets/", http.FileServer(http.Dir("dist")))


	server.Server.Handler = router

	storage, err := database.NewSQLite3Store(dbFilePath)
	if err != nil {
		return nil, err
	}
	server.Storage = storage

	server.Sessions = database.NewSessionStore()
	return server, nil
}
