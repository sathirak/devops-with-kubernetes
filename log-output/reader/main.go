package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type Status struct {
	Timestamp string `json:"timestamp"`
	Message   string `json:"message"`
	Hash      string `json:"hash"`
	PingPongs int    `json:"pingpongs"`
}

func readPingPongs() int {
	data, err := os.ReadFile("/usr/src/app/shared/counter.txt")
	if err != nil {
		log.Printf("Error reading ping-pong counter: %v", err)
		return 0
	}
	count := 0
	if _, err := fmt.Sscanf(string(data), "%d", &count); err != nil {
		log.Printf("Error parsing counter value: %v", err)
		return 0
	}
	return count
}

func main() {
	statusFile := "/usr/src/app/files/status.txt"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data, err := os.ReadFile(statusFile)
		if err != nil {
			http.Error(w, "File not found", 404)
			return
		}

		var status Status
		json.Unmarshal(data, &status)

		status.Timestamp = time.Now().UTC().Format(time.RFC3339Nano)
		hash := sha256.Sum256(data)
		status.Hash = fmt.Sprintf("%x", hash)
		status.PingPongs = readPingPongs()

		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintf(w, "%s: %s\nPing / Pongs: %d",
			status.Timestamp, status.Hash, status.PingPongs)
	})

	http.ListenAndServe(":3000", nil)
}
