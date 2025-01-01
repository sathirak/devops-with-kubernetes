package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
	"net/http"
	"encoding/json"
)

type Status struct {
	Timestamp string `json:"timestamp"`
	Message   string `json:"message"`
}

var currentStatus Status

func generateRandomLetters(length int) string {
	const letters = "abcdefghijklmnopqrstuvwxyz"
	var sb strings.Builder
	for i := 0; i < length; i++ {
		sb.WriteByte(letters[rand.Intn(len(letters))])
	}
	return sb.String()
}

func main() {
	words := []string{"pooh", "tigger", "piglet", "eeyore", "rabbit", "owl", "roo", "kanga"}
	
	// Add HTTP handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(currentStatus)
	})
	
	// Start HTTP server in goroutine
	go http.ListenAndServe(":3000", nil)
	
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		timestamp := time.Now().Format("2006-01-02 15:04:05")
		randomWord := words[rand.Intn(len(words))]
		randomLetters := generateRandomLetters(5)
		message := fmt.Sprintf("%s_%s", randomWord, randomLetters)
		
		currentStatus = Status{
			Timestamp: timestamp,
			Message:   message,
		}
		
		fmt.Printf("[%s] %s\n", timestamp, message)
	}
}
