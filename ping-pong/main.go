package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"os"
	"path/filepath"
)

const port = ":3000"
const counterFile = "/usr/src/app/shared/counter.txt"
var counter = 0

func readCounter() int {
	data, err := os.ReadFile(counterFile)
	if err != nil {
		return 0
	}
	count, _ := strconv.Atoi(string(data))
	return count
}

func writeCounter(count int) error {
	dir := filepath.Dir(counterFile)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}
	return os.WriteFile(counterFile, []byte(strconv.Itoa(count)), 0644)
}

func main() {
	counter = readCounter()

	http.HandleFunc("/pingpong", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		counter++
		if err := writeCounter(counter); err != nil {
			log.Printf("Error writing counter: %v", err)
		}
		fmt.Fprintf(w, "pong %d", counter)
	})

	fmt.Printf("Server started on port %s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
