package services

import (
	"WealthNoteBackend/internal/database"
	"WealthNoteBackend/internal/models"
	"WealthNoteBackend/internal/repositories"
)

// GetTransactionByUserID - ดึง transactions ของ user ที่ login
func GetTransactionByUserID(userID int) ([]models.Transaction, error) {
	transactionRepo := repositories.NewTransactionRepository(database.GetPostgresDB())
	return transactionRepo.FindByUserID(userID)
}

// CreateTransaction - สร้าง transaction ใหม่
func CreateTransaction(input models.CreateTransactionInput) (*models.Transaction, error) {
	transactionRepo := repositories.NewTransactionRepository(database.GetPostgresDB())
	return transactionRepo.CreateTransaction(input)
}
