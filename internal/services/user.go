package services

import (
	"WealthNoteBackend/internal/database"
	"WealthNoteBackend/internal/models"
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/go-playground/validator/v10"
)

type UserService struct {
	validator *validator.Validate
}

func NewUserService() *UserService {
	return &UserService{
		validator: validator.New(),
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

func GetAllUsers() ([]models.User, error) {
	db := database.GetPostgresDB()

	query := `SELECT id_user, usercode, email, fnamet, lnamet, fnamee, lnamee, 
              provider, provider_id, created_at, updated_at FROM users ORDER BY id_user`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(
			&user.IDUser, &user.UserCode, &user.Email,
			&user.FNameT, &user.LNameT, &user.FNameE, &user.LNameE,
			&user.Provider, &user.ProviderID,
			&user.CreatedAt, &user.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func GetUserByID(id int) (*models.User, error) {
	db := database.GetPostgresDB()

	query := `SELECT id_user, usercode, email, fnamet, lnamet, fnamee, lnamee, 
              provider, provider_id, created_at, updated_at FROM users WHERE id_user = $1`

	var user models.User
	err := db.QueryRow(query, id).Scan(
		&user.IDUser, &user.UserCode, &user.Email,
		&user.FNameT, &user.LNameT, &user.FNameE, &user.LNameE,
		&user.Provider, &user.ProviderID,
		&user.CreatedAt, &user.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, errors.New("user not found")
	}
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func CreateUser(usercode, email string) (*models.User, error) {
	db := database.GetPostgresDB()

	query := `INSERT INTO users (usercode, email)
              VALUES ($1, $2)
              RETURNING id_user, usercode, email, created_at, updated_at`

	var user models.User
	err := db.QueryRow(query, usercode, email).Scan(
		&user.IDUser, &user.UserCode, &user.Email,
		&user.CreatedAt, &user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func UpdateUser(id int, input models.UpdateUserInput) (*models.User, error) {
	db := database.GetPostgresDB()

	query := `UPDATE users SET 
              usercode = COALESCE($1, usercode),
              email = COALESCE($2, email),
              fnamet = COALESCE($3, fnamet),
              lnamet = COALESCE($4, lnamet),
              fnamee = COALESCE($5, fnamee),
              lnamee = COALESCE($6, lnamee),
              updated_at = $7
              WHERE id_user = $8
              RETURNING id_user, usercode, email, fnamet, lnamet, fnamee, lnamee, provider, provider_id, created_at, updated_at`

	var user models.User
	err := db.QueryRow(
		query,
		input.UserCode, input.Email, input.FNameT, input.LNameT,
		input.FNameE, input.LNameE, time.Now(), id,
	).Scan(
		&user.IDUser, &user.UserCode, &user.Email,
		&user.FNameT, &user.LNameT, &user.FNameE, &user.LNameE,
		&user.Provider, &user.ProviderID,
		&user.CreatedAt, &user.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, errors.New("user not found")
	}
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func DeleteUser(id int) error {
	db := database.GetPostgresDB()

	query := `DELETE FROM users WHERE id_user = $1`
	result, err := db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("user not found")
	}

	return nil
}
