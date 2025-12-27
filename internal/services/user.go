package services

import (
	"WealthNoteBackend/internal/database"
	"WealthNoteBackend/internal/models"
	"WealthNoteBackend/internal/repositories"
	"context"

	"github.com/go-playground/validator/v10"
)

type UserService struct {
	validator *validator.Validate
	userRepo  *repositories.UserRepository
}

func NewUserService() *UserService {
	return &UserService{
		validator: validator.New(),
		userRepo:  repositories.NewUserRepository(database.GetPostgresDB()),
	}
}

func (s *UserService) CreateUser(ctx context.Context, user *models.User) error {
	if err := s.validator.Struct(user); err != nil {
		return err
	}
	// TODO: Implement database operations
	return nil
}

func (s *UserService) UpdateUser(ctx context.Context, user *models.User) error {
	if err := s.validator.Struct(user); err != nil {
		return err
	}
	// TODO: Implement database operations
	return nil
}

func (s *UserService) DeleteUser(ctx context.Context, id string) error {
	// TODO: Implement database operations
	return nil
}

// ✅ ใช้ Repository แทน Query ตรงๆ
func GetAllUsers() ([]models.User, error) {
	userRepo := repositories.NewUserRepository(database.GetPostgresDB())
	return userRepo.FindAll()
}

func GetUserByID(id int) (*models.User, error) {
	userRepo := repositories.NewUserRepository(database.GetPostgresDB())
	return userRepo.FindByID(id)
}

func UpdateUser(id int, input models.UpdateUserInput) (*models.User, error) {
	userRepo := repositories.NewUserRepository(database.GetPostgresDB())
	return userRepo.Update(id, input)
}

func DeleteUser(id int) error {
	userRepo := repositories.NewUserRepository(database.GetPostgresDB())
	return userRepo.Delete(id)
}
