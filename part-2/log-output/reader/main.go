package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type Status struct {
	Timestamp string `json:"timestamp"`
	Message   string `json:"message"`
	PingPongs int    `json:"pingpongs"`
}

func getPingPongs() int {
	resp, err := http.Get("http://ping-pong-svc:2345/count")
	if err != nil {
		log.Printf("Error getting ping-pong count: %v", err)
		return 0
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response: %v", err)
		return 0
	}

	var count int
	if _, err := fmt.Sscanf(string(body), "%d", &count); err != nil {
		log.Printf("Error parsing counter value: %v", err)
		return 0
	}
	return count
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		var status Status

		status.Timestamp = time.Now().UTC().Format(time.RFC3339Nano)
		status.PingPongs = getPingPongs()

		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintf(w, "%s\nPing / Pongs: %d",
			status.Timestamp, status.PingPongs)
	})

	http.ListenAndServe(":3000", nil)
}
