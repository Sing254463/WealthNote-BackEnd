package repositories

import (
	"WealthNoteBackend/internal/models"
	"database/sql"
)

type TransactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{db: db}

}

// FindAll - ดึง transactions ทั้งหมด
func (r *TransactionRepository) FindAll() ([]models.Transaction, error) {
	query := `SELECT id_transaction, amount, description, date, created_at, updated_at FROM transactions`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var transactions []models.Transaction
	for rows.Next() {
		var transaction models.Transaction
		err := rows.Scan(
			&transaction.IDTransaction,
			&transaction.Amount,
			&transaction.Description,
			&transaction.Date,
			&transaction.CreatedAt,
			&transaction.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}
	return transactions, nil
}
