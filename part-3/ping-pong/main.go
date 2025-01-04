package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

const port = ":3000"

var db *sql.DB

func initDB() error {
	psqlInfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"))

	var err error
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		return err
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS ping_pong_counter (
			id SERIAL PRIMARY KEY,
			counter INTEGER NOT NULL
		)
	`)
	if err != nil {
		return err
	}

	_, err = db.Exec(`
		INSERT INTO ping_pong_counter (id, counter)
		SELECT 1, 0
		WHERE NOT EXISTS (SELECT 1 FROM ping_pong_counter WHERE id = 1)
	`)
	return err
}

func getCounter() (int, error) {
	var count int
	err := db.QueryRow("SELECT counter FROM ping_pong_counter WHERE id = 1").Scan(&count)
	return count, err
}

func incrementCounter() (int, error) {
	var count int
	err := db.QueryRow(`
		UPDATE ping_pong_counter 
		SET counter = counter + 1 
		WHERE id = 1 
		RETURNING counter
	`).Scan(&count)
	return count, err
}

func main() {
	if err := initDB(); err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	http.HandleFunc("/pingpong", func(w http.ResponseWriter, r *http.Request) {
		count, err := incrementCounter()
		if (err != nil) {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintf(w, "pong %d", count)
	})

	http.HandleFunc("/count", func(w http.ResponseWriter, r *http.Request) {
		count, err := getCounter()
		if (err != nil) {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintf(w, "%d", count)
	})

	fmt.Printf("Server started on port %s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
