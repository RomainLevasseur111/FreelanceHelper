package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func register(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost{
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	var data User

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Println(data)


}