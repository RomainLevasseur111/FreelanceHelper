package api

import "net/http"

type API struct {
	http.Server
}

func NewAPI(addr string) (*API, error) {
	server := new(API)
	server.Server.Addr = addr

	router := http.NewServeMux()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "dist/index.html")
	})
	router.Handle("/assets/", http.FileServer(http.Dir("dist")))

	server.Server.Handler = router

	return server, nil
}
