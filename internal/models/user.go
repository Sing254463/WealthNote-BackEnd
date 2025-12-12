package models

import (
	"time"
)

// User represents the structure of a user in the system.
type User struct {
	IDUser     int       `json:"id_user" db:"id_user"`
	UserCode   *string   `json:"usercode,omitempty" db:"usercode"`
	Email      string    `json:"email" db:"email"`
	FNameT     *string   `json:"fnamet,omitempty" db:"fnamet"`
	LNameT     *string   `json:"lnamet,omitempty" db:"lnamet"`
	FNameE     *string   `json:"fnamee,omitempty" db:"fnamee"`
	LNameE     *string   `json:"lnamee,omitempty" db:"lnamee"`
	Password   string    `json:"-" db:"password"` // ไม่ส่งออก JSON
	Provider   *string   `json:"provider,omitempty" db:"provider"`
	ProviderID *string   `json:"provider_id,omitempty" db:"provider_id"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
}

// LoginInput represents the input data for user login.
type LoginInput struct {
	UserCode string `json:"usercode" validate:"required" example:"user001"`
	Password string `json:"password" validate:"required" example:"password123"`
}

// RegisterInput represents the input data for registering a new user.
type RegisterInput struct {
	UserCode string  `json:"usercode" validate:"required" example:"user001"`
	Email    string  `json:"email" validate:"required,email" example:"user@example.com"`
	FNameT   *string `json:"fnamet,omitempty" example:"สมชาย"`
	LNameT   *string `json:"lnamet,omitempty" example:"ใจดี"`
	FNameE   *string `json:"fnamee,omitempty" example:"Somchai"`
	LNameE   *string `json:"lnamee,omitempty" example:"Jaidee"`
	Password string  `json:"password" validate:"required,min=6" example:"password123"`
}

// UpdateUserInput represents the input data for updating an existing user.
type UpdateUserInput struct {
	UserCode *string `json:"usercode"`
	Email    *string `json:"email" validate:"omitempty,email"`
	FNameT   *string `json:"fnamet"`
	LNameT   *string `json:"lnamet"`
	FNameE   *string `json:"fnamee"`
	LNameE   *string `json:"lnamee"`
}
