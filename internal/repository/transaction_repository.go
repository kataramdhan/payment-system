package repository

import (
	"database/sql"
	"payment-system/internal/model"
)

type TransactionRepository struct {
	DB *sql.DB
}

func NewTransactionRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{DB: db}
}

func (r *TransactionRepository) Create(userID int, amount float64) (int, error) {
	var id int

	query := `
		INSERT INTO transactions (user_id, amount, status)
		VALUES ($1, $2, 'pending')
		RETURNING id
	`

	err := r.DB.QueryRow(query, userID, amount).Scan(&id)
	return id, err
}

func (r *TransactionRepository) UpdateStatus(id int, status string) error {
	query := `
		UPDATE transactions
		SET status = $1, updated_at = NOW()
		WHERE id = $2
	`
	_, err := r.DB.Exec(query, status, id)
	return err
}

func (r *TransactionRepository) FindAll() ([]model.Transaction, error) {
	rows, err := r.DB.Query(`
		SELECT id, user_id, amount, status, created_at
		FROM transactions
		ORDER BY id DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []model.Transaction

	for rows.Next() {
		var tx model.Transaction
		err := rows.Scan(&tx.ID, &tx.UserID, &tx.Amount, &tx.Status, &tx.CreatedAt)
		if err != nil {
			return nil, err
		}
		result = append(result, tx)
	}

	return result, nil
}
