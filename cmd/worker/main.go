package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"time"

	_ "github.com/lib/pq"

	"github.com/hibiken/asynq"
)

func main() {

	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: "localhost:6379"},
		asynq.Config{
			Concurrency: 5,
			ErrorHandler: asynq.ErrorHandlerFunc(func(ctx context.Context, task *asynq.Task, err error) {
				fmt.Println("❌ Task failed permanently:", err)
			}),
		},
	)

	mux := asynq.NewServeMux()
	mux.HandleFunc("process:transaction", processTransaction)

	if err := srv.Run(mux); err != nil {
		log.Fatal(err)
	}
}

func processTransaction(ctx context.Context, t *asynq.Task) error {
	var payload map[string]int

	if err := json.Unmarshal(t.Payload(), &payload); err != nil {
		return err
	}

	txID := payload["transaction_id"]

	db := newDB()

	fmt.Println("Processing transaction ID:", txID)

	time.Sleep(2 * time.Second)

	// simulasi success / failed
	if rand.Intn(2) == 0 {
		fmt.Println("SUCCESS:", txID)

		_, err := db.Exec(`
			UPDATE transactions
			SET status = 'success', updated_at = NOW()
			WHERE id = $1
		`, txID)

		return err
	}

	// ❌ gagal → return error supaya retry
	fmt.Println("FAILED → RETRY:", txID)

	return fmt.Errorf("payment failed, will retry")
}

func newDB() *sql.DB {
	connStr := "host=localhost port=5432 user=postgres password=password dbname=payment_db sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
