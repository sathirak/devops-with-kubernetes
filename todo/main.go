package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = ":3000"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprintf(w, `
			<!DOCTYPE html>
			<html>
				<head>
					<title>Welcome</title>
				</head>
				<body>
					<h1>Welcome to the Todo Server</h1>
					<p>This is a simple web server built with Go.</p>
				</body>
			</html>
		`)
	})

	fmt.Printf("Server started on port %s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
