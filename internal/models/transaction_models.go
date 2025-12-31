package models

import "time"

// TypeTransaction represents transaction types (income/expense)
type TypeTransaction struct {
	ID        int    `json:"id" db:"id"`
	NameTypeT string `json:"name_type_t" db:"nametypet"`
	NameTypeE string `json:"name_type_e" db:"nametypee"`
}

// Category represents transaction categories
type Category struct {
	ID            int    `json:"id" db:"id"`
	NameCategoryT string `json:"name_category_t" db:"name_categoryt"`
	NameCategoryE string `json:"name_category_e" db:"name_categorye"`
}

// Transaction represents a financial transaction
type Transaction struct {
	ID                int       `json:"id" db:"id"`
	IDUser            int       `json:"id_user" db:"id_user"`
	IDType            int       `json:"id_type" db:"id_type"`
	Title             string    `json:"title" db:"title"`
	Description       *string   `json:"description,omitempty" db:"description"`
	Amount            float64   `json:"amount" db:"amount"`
	IDCategory        int       `json:"id_category" db:"id_category"`
	OtherCategoryName *string   `json:"other_category_name,omitempty" db:"other_category_name"`
	TransactionDate   time.Time `json:"transaction_date" db:"transaction_date"`
	CreatedAt         time.Time `json:"created_at" db:"created_at"`
	UpdatedAt         time.Time `json:"updated_at" db:"updated_at"`
}

// CreateTransactionInput for creating new transaction
type CreateTransactionInput struct {
	IDUser            int       `json:"id_user" validate:"required"`
	IDType            int       `json:"id_type" validate:"required"`
	Title             string    `json:"title" validate:"required"`
	Description       *string   `json:"description,omitempty"`
	Amount            float64   `json:"amount" validate:"required,gt=0"`
	IDCategory        int       `json:"id_category" validate:"required"`
	OtherCategoryName *string   `json:"other_category_name,omitempty"`
	TransactionDate   time.Time `json:"transaction_date" validate:"required"`
}

// UpdateTransactionInput for updating transaction
type UpdateTransactionInput struct {
	IDType            *int       `json:"id_type,omitempty"`
	Title             *string    `json:"title,omitempty"`
	Description       *string    `json:"description,omitempty"`
	Amount            *float64   `json:"amount,omitempty" validate:"omitempty,gt=0"`
	IDCategory        *int       `json:"id_category,omitempty"`
	OtherCategoryName *string    `json:"other_category_name,omitempty"`
	TransactionDate   *time.Time `json:"transaction_date,omitempty"`
}
