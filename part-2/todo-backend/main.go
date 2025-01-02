package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const port = ":3000"

type Todo struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}

var todos []Todo
var nextID = 1

func main() {
	http.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		switch r.Method {
		case http.MethodGet:
			json.NewEncoder(w).Encode(todos)
		case http.MethodPost:
			var todo Todo
			if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			todo.ID = nextID
			nextID++
			todos = append(todos, todo)
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(todo)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	fmt.Printf("Server started on port %s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
