package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = ":3000"

var counter = 0

func main() {

	http.HandleFunc("/pingpong", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		counter++
		fmt.Fprintf(w, "pong %d", counter)
	})

	http.HandleFunc("/count", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintf(w, "%d", counter)
	})

	fmt.Printf("Server started on port %s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
