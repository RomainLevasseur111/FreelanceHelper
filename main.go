package main

import (
	"log"

	api "freelancehelper/backend"
)

func main() {
	server, err := api.NewAPI(":8080")
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Server up and running on port %s\n", server.Addr)
	err = server.ListenAndServe()
	if err != nil {
		log.Println("Error", err)
	}
}
