
package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Status struct {
	Timestamp string `json:"timestamp"`
	Message   string `json:"message"`
}

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
	outputFile := "/usr/src/app/files/status.txt"
	
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		timestamp := time.Now().Format("2006-01-02 15:04:05")
		randomWord := words[rand.Intn(len(words))]
		randomLetters := generateRandomLetters(5)
		message := fmt.Sprintf("%s_%s", randomWord, randomLetters)
		
		status := Status{
			Timestamp: timestamp,
			Message:   message,
		}
		
		jsonData, _ := json.Marshal(status)
		os.WriteFile(outputFile, jsonData, 0644)
		fmt.Printf("Written: [%s] %s\n", timestamp, message)
	}
}