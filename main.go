package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/guilhermescr/crud-golang/domain"
)

func main() {
	http.HandleFunc("/person/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			var person domain.Person
			err := json.NewDecoder(r.Body).Decode(&person)
			if err != nil {
				fmt.Printf("Error trying to decode body. Body should be a json. Error: %s", err.Error())
				http.Error(w, "Error trying to create person", http.StatusBadRequest)
				return
			}
			if person.ID <= 0 {
				http.Error(w, "Error trying to create person. ID should be a positive integer", http.StatusBadRequest)
				return
			}

			w.WriteHeader(http.StatusCreated)
			return
		}

		http.Error(w, "Not Implemented", http.StatusInternalServerError)
	})

	http.ListenAndServe(":8080", nil)
}