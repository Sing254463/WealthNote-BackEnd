package services

import (
	"WealthNoteBackend/internal/database"
	"WealthNoteBackend/internal/models"
	"WealthNoteBackend/internal/repositories"
)

// GetTransactionByUserID - ดึง transactions ของ user ที่ login
func GetTransactionByUserID(userID int, lang string) ([]models.Transaction, error) {
	transactionRepo := repositories.NewTransactionRepository(database.GetPostgresDB())
	return transactionRepo.FindByUserID(userID, lang)
}

// CreateTransaction - สร้าง transaction ใหม่
func CreateTransaction(input models.CreateTransactionInput) (*models.Transaction, error) {
	transactionRepo := repositories.NewTransactionRepository(database.GetPostgresDB())
	return transactionRepo.CreateTransaction(input)
}

func UpdateTransaction(transactionID int, input models.UpdateTransactionInput) (*models.Transaction, error) {
	transactionRepo := repositories.NewTransactionRepository(database.GetPostgresDB())
	return transactionRepo.UpdateTransaction(transactionID, input)
}
