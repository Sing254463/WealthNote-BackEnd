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

func (r *TransactionRepository) FindByUserID(userID int) ([]models.Transaction, error) {
	query := `
        SELECT 
            transactions.id, 
            transactions.id_user, 
            transactions.id_type,
            type_transactions.nametypet, 
            transactions.title, 
            transactions.description, 
            transactions.amount, 
            transactions.id_category,
            category.name_categoryt, 
            transactions.other_category_name, 
            transactions.transaction_date, 
            transactions.created_at, 
            transactions.updated_at
        FROM public.transactions
        LEFT JOIN type_transactions ON type_transactions.id = transactions.id_type
        LEFT JOIN category ON category.id = transactions.id_category
        WHERE transactions.id_user = $1
        ORDER BY transactions.transaction_date DESC
    `

	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []models.Transaction
	for rows.Next() {
		var transaction models.Transaction
		var typeNameT, categoryNameT sql.NullString

		err := rows.Scan(
			&transaction.ID,
			&transaction.IDUser,
			&transaction.IDType,
			&typeNameT, // ชื่อประเภท
			&transaction.Title,
			&transaction.Description,
			&transaction.Amount,
			&transaction.IDCategory,
			&categoryNameT, // ชื่อหมวดหมู่
			&transaction.OtherCategoryName,
			&transaction.TransactionDate,
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
