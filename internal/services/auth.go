package services

import (
	"WealthNoteBackend/internal/database"
	"WealthNoteBackend/internal/models"
	"WealthNoteBackend/internal/repositories"
	"WealthNoteBackend/pkg/jwt"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo *repositories.UserRepository
}

func NewAuthService() *AuthService {
	return &AuthService{
		userRepo: repositories.NewUserRepository(database.GetPostgresDB()),
	}
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
	userRepo := repositories.NewUserRepository(database.GetPostgresDB())

	// ✅ ใช้ Repository แทน Query ตรงๆ
	user, hashedPassword, err := userRepo.FindByUserCode(usercode)
	if err != nil {
		return "", "", errors.New("invalid usercode or password")
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
	userRepo := repositories.NewUserRepository(database.GetPostgresDB())

	// ✅ ใช้ Repository เช็คข้อมูลซ้ำ
	exists, err := userRepo.ExistsByUserCode(input.UserCode)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.New("usercode already exists")
	}

	exists, err = userRepo.ExistsByEmail(input.Email)
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

	// ✅ ใช้ Repository สร้าง user
	return userRepo.Create(input, string(hashedPassword))
}
