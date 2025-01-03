package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type Status struct {
	Timestamp     string `json:"timestamp"`
	Message       string `json:"message"`
	PingPongs     int    `json:"pingpongs"`
	ConfigMessage string `json:"config_message"`
	FileContent   string `json:"file_content"`
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

func readInfoFile() string {
	content, err := os.ReadFile("/usr/src/app/files/information.txt")
	if err != nil {
		log.Printf("Error reading information.txt: %v", err)
		return "Error reading file"
	}
	return string(content)
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		var status Status

		status.Timestamp = time.Now().UTC().Format(time.RFC3339Nano)
		status.PingPongs = getPingPongs()
		status.ConfigMessage = os.Getenv("MESSAGE")
		status.FileContent = readInfoFile()

		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintf(w, "%s\nPing / Pongs: %d\nmessage: %s\nfile contents:\n%s",
			status.Timestamp,
			status.PingPongs,
			status.ConfigMessage,
			status.FileContent)
	})

	http.ListenAndServe(":3000", nil)
}
