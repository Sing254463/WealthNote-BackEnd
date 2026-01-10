package repositories

import (
	"WealthNoteBackend/internal/models"
	"database/sql"
)

type TransactionRepository struct {
	db *sql.DB
}

// NewTransactionRepository - สร้าง instance ของ TransactionRepository
func NewTransactionRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

// FindByUserID - ดึง transactions ของ user
func (r *TransactionRepository) FindByUserID(userID int, lang string) ([]models.Transaction, error) {
	query := `
        SELECT 
            transactions.id, 
            transactions.id_user, 
            transactions.id_type,
            type_transactions.nametype` + lang + ` as nametypet, 
            transactions.title, 
            transactions.description, 
            transactions.amount, 
            transactions.id_category,
            category.name_category` + lang + ` as name_category, 
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

		err := rows.Scan(
			&transaction.ID,
			&transaction.IDUser,
			&transaction.IDType,
			&transaction.NameType,
			&transaction.Title,
			&transaction.Description,
			&transaction.Amount,
			&transaction.IDCategory,
			&transaction.NameCategory,
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

// CreateTransaction - สร้าง transaction ใหม่
func (r *TransactionRepository) CreateTransaction(input models.CreateTransactionInput) (*models.Transaction, error) {
	query := `
        INSERT INTO transactions(
            id_user, id_type, 
            title, description, amount, 
            id_category, other_category_name, 
            transaction_date, created_at, updated_at
        )
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, NOW(), NOW())
        RETURNING id, id_user, id_type, title, description, amount, 
                  id_category, other_category_name, transaction_date, 
                  created_at, updated_at
    `

	var transaction models.Transaction
	err := r.db.QueryRow(
		query,
		input.IDUser,
		input.IDType,
		input.Title,
		input.Description,
		input.Amount,
		input.IDCategory,
		input.OtherCategoryName,
		input.TransactionDate,
	).Scan(
		&transaction.ID,
		&transaction.IDUser,
		&transaction.IDType,
		&transaction.Title,
		&transaction.Description,
		&transaction.Amount,
		&transaction.IDCategory,
		&transaction.OtherCategoryName,
		&transaction.TransactionDate,
		&transaction.CreatedAt,
		&transaction.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &transaction, nil
}

func (r *TransactionRepository) UpdateTransaction(id int, input models.UpdateTransactionInput) (*models.Transaction, error) {

	query := `
		UPDATE transactions
		SET id_type = $1,
			title = $2,
			description = $3,
			amount = $4,
			id_category = $5,
			other_category_name = $6,
			transaction_date = $7,
			updated_at = NOW()
		WHERE id = $8
		RETURNING id, id_user, id_type, title, description, amount,
				  id_category, other_category_name, transaction_date,
				  created_at, updated_at
	`
	var transaction models.Transaction
	err := r.db.QueryRow(
		query,
		input.IDType,
		input.Title,
		input.Description,
		input.Amount,
		input.IDCategory,
		input.OtherCategoryName,
		input.TransactionDate,
		id,
	).Scan(
		&transaction.ID,
		&transaction.IDUser,
		&transaction.IDType,
		&transaction.Title,
		&transaction.Description,
		&transaction.Amount,
		&transaction.IDCategory,
		&transaction.OtherCategoryName,
		&transaction.TransactionDate,
		&transaction.CreatedAt,
		&transaction.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return nil, nil
}
