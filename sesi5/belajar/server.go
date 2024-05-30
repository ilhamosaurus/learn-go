package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var PORT = ":8080"

type User struct {
	ID       int
	Username string
}

var users = []*User{
	{ID: 1, Username: "codeium"},
}

func main() {
	defer func() {
		r := recover()
		if r != nil {
			log.Println(r)
		}
	}()

	http.HandleFunc("/", greet)

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			w.Header().Set("Content-Type", "application/json")

			err := json.NewEncoder(w).Encode(users)
			if err != nil {
				log.Println(err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}

		if r.Method == http.MethodPost {
			username := r.FormValue("username")
			log.Println(username)
			if username == "" {
				http.Error(w, "username must be provided", http.StatusBadRequest)
				return
			}

			user := &User{Username: username, ID: len(users) + 1}
			users = append(users, user)
			err := json.NewEncoder(w).Encode(users)
			if err != nil {
				log.Println(err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
	})

	log.Printf("Server running on port %s", PORT)
	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}

func greet(w http.ResponseWriter, r *http.Request) {
	msg := "Hello world"
	fmt.Fprint(w, msg)
}
