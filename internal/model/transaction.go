package model

import "time"

type Transaction struct {
	ID          int       `json:"id"`
	UserID      int       `json:"user_id"`
	Amount      float64   `json:"amount"`
	Status      string    `json:"status"`
	ReferenceID string    `json:"reference_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
