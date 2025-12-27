package repositories

import (
	"WealthNoteBackend/internal/models"
	"database/sql"
	"errors"
	"time"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

// FindByUserCode - ค้นหา user ด้วย usercode
func (r *UserRepository) FindByUserCode(usercode string) (*models.User, string, error) {
	var user models.User
	var hashedPassword string

	query := `SELECT id_user, usercode, email, password FROM users WHERE usercode = $1`
	err := r.db.QueryRow(query, usercode).Scan(&user.IDUser, &user.UserCode, &user.Email, &hashedPassword)

	if err == sql.ErrNoRows {
		return nil, "", errors.New("user not found")
	}
	if err != nil {
		return nil, "", err
	}

	return &user, hashedPassword, nil
}

// FindByID - ค้นหา user ด้วย ID
func (r *UserRepository) FindByID(id int) (*models.User, error) {
	query := `SELECT id_user, usercode, email, fnamet, lnamet, fnamee, lnamee, 
              provider, provider_id, created_at, updated_at FROM users WHERE id_user = $1`

	var user models.User
	err := r.db.QueryRow(query, id).Scan(
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

// FindAll - ดึง users ทั้งหมด
func (r *UserRepository) FindAll() ([]models.User, error) {
	query := `SELECT id_user, usercode, email, fnamet, lnamet, fnamee, lnamee, 
              provider, provider_id, created_at, updated_at FROM users ORDER BY id_user`

	rows, err := r.db.Query(query)
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

// ExistsByUserCode - เช็คว่า usercode ซ้ำหรือไม่
func (r *UserRepository) ExistsByUserCode(usercode string) (bool, error) {
	var exists bool
	err := r.db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE usercode = $1)", usercode).Scan(&exists)
	return exists, err
}

// ExistsByEmail - เช็คว่า email ซ้ำหรือไม่
func (r *UserRepository) ExistsByEmail(email string) (bool, error) {
	var exists bool
	err := r.db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)", email).Scan(&exists)
	return exists, err
}

// Create - สร้าง user ใหม่
func (r *UserRepository) Create(input models.RegisterInput, hashedPassword string) (*models.User, error) {
	query := `INSERT INTO users (usercode, email, fnamet, lnamet, fnamee, lnamee, password)
              VALUES ($1, $2, $3, $4, $5, $6, $7)
              RETURNING id_user, usercode, email, fnamet, lnamet, fnamee, lnamee, created_at, updated_at`

	var user models.User
	err := r.db.QueryRow(
		query,
		input.UserCode, input.Email, input.FNameT, input.LNameT,
		input.FNameE, input.LNameE, hashedPassword,
	).Scan(
		&user.IDUser, &user.UserCode, &user.Email,
		&user.FNameT, &user.LNameT, &user.FNameE, &user.LNameE,
		&user.CreatedAt, &user.UpdatedAt,
	)

	return &user, err
}

// Update - แก้ไขข้อมูล user
func (r *UserRepository) Update(id int, input models.UpdateUserInput) (*models.User, error) {
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
	err := r.db.QueryRow(
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

	return &user, err
}

// Delete - ลบ user
func (r *UserRepository) Delete(id int) error {
	result, err := r.db.Exec("DELETE FROM users WHERE id_user = $1", id)
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
