package models

import (
	"time"
)

// Transaction represents the structure of a transaction in the system.
type Transaction struct {
	IDTransaction int       `json:"id_transaction" db:"id_transaction"`
	Amount        float64   `json:"amount" db:"amount"`
	Description   *string   `json:"description,omitempty" db:"description"`
	Date          time.Time `json:"date" db:"date"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" db:"updated_at"`
}

// CreateTransactionInput represents the input data for creating a new transaction.
type CreateTransactionInput struct {
	Amount      float64   `json:"amount" validate:"required"`
	Description *string   `json:"description,omitempty"`
	Date        time.Time `json:"date" validate:"required"`
}

// UpdateTransactionInput represents the input data for updating an existing transaction.
type UpdateTransactionInput struct {
	Amount      *float64   `json:"amount,omitempty"`
	Description *string    `json:"description,omitempty"`
	Date        *time.Time `json:"date,omitempty"`
}
