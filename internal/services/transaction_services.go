package services

import (
	"WealthNoteBackend/internal/database"
	"WealthNoteBackend/internal/models"
	"WealthNoteBackend/internal/repositories"
)

func GetTransactionAll() ([]models.Transaction, error) {
	transactionRepo := repositories.NewTransactionRepository(database.GetPostgresDB())
	return transactionRepo.FindAll()
}
