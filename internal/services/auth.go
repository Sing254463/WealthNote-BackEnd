package services

import (
	"WealthNoteBackend/internal/database"
	"WealthNoteBackend/internal/models"
	"WealthNoteBackend/pkg/jwt"
	"database/sql"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (s *AuthService) GenerateTokens(userID string) (string, string, error) {
	accessToken, err := jwt.GenerateToken(userID, false)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := jwt.GenerateToken(userID, true)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (s *AuthService) ValidateToken(tokenString string) error {
	_, err := jwt.ValidateToken(tokenString)
	return err
}

// Login - เข้าสู่ระบบด้วย usercode
func Login(usercode, password string) (string, string, error) {
	db := database.GetPostgresDB()

	// ดึง user และ hashed password จาก database โดยใช้ usercode
	var user models.User
	var hashedPassword string

	query := `SELECT id_user, usercode, email, password FROM users WHERE usercode = $1`
	err := db.QueryRow(query, usercode).Scan(&user.IDUser, &user.UserCode, &user.Email, &hashedPassword)

	if err == sql.ErrNoRows {
		return "", "", errors.New("invalid usercode or password")
	}
	if err != nil {
		return "", "", err
	}

	// เปรียบเทียบ password
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return "", "", errors.New("invalid usercode or password")
	}

	// สร้าง JWT tokens
	accessToken, err := jwt.GenerateToken(string(user.IDUser), false)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := jwt.GenerateToken(string(user.IDUser), true)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

// Register - สมัครสมาชิกใหม่
func Register(input models.RegisterInput) (*models.User, error) {
	db := database.GetPostgresDB()

	// ตรวจสอบว่า usercode ซ้ำหรือไม่
	var exists bool
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE usercode = $1)", input.UserCode).Scan(&exists)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.New("usercode already exists")
	}

	// ตรวจสอบว่า email ซ้ำหรือไม่
	err = db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)", input.Email).Scan(&exists)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.New("email already exists")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// บันทึกผู้ใช้ใหม่
	query := `INSERT INTO users (usercode, email, fnamet, lnamet, fnamee, lnamee, password)
              VALUES ($1, $2, $3, $4, $5, $6, $7)
              RETURNING id_user, usercode, email, fnamet, lnamet, fnamee, lnamee, created_at, updated_at`

	var user models.User
	err = db.QueryRow(
		query,
		input.UserCode, input.Email, input.FNameT, input.LNameT,
		input.FNameE, input.LNameE, string(hashedPassword),
	).Scan(
		&user.IDUser, &user.UserCode, &user.Email,
		&user.FNameT, &user.LNameT, &user.FNameE, &user.LNameE,
		&user.CreatedAt, &user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
