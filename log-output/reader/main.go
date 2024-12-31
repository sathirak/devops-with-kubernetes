
package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Status struct {
	Timestamp string `json:"timestamp"`
	Message   string `json:"message"`
	Hash      string `json:"hash"`
}

func main() {
	statusFile := "/usr/src/app/files/status.txt"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data, err := os.ReadFile(statusFile)
		if (err != nil) {
			http.Error(w, "File not found", 404)
			return
		}

		var status Status
		json.Unmarshal(data, &status)
		
		hash := sha256.Sum256(data)
		status.Hash = fmt.Sprintf("%x", hash)
		
		json.NewEncoder(w).Encode(status)
	})

	http.ListenAndServe(":3000", nil)
}