package models

import (
	"github.com/satori/go.uuid"
	"gorm.io/gorm"
	"time"
)

type LoginInput struct {
	Email    string `json:"email" validate:"required,email,havent_script"`
	Password string `json:"password" validate:"required,havent_script"`
}

type ForgotPasswordInput struct {
	Email string `json:"email" validate:"required,email,havent_script"`
}

type ResetPasswordInput struct {
	Email    string    `json:"email" validate:"required,email,havent_script"`
	Token    uuid.UUID `json:"token" validate:"required" gorm:"type:uuid;"`
	Password string    `json:"password" validate:"required,havent_script"`
}

type ResetPasswordByPass struct {
	Email    string    `json:"email" validate:"email,havent_script"`
	Token    uuid.UUID `json:"token"  gorm:"type:uuid;"`
	Password string    `json:"password" validate:"havent_script"`
}

type UserRole struct {
	UserId uuid.UUID `json:"user_id" gorm:"primary_key;type:uuid;"`
	Role   string    `json:"role" validate:"required,havent_script" gorm:"primary_key;size:64"`
}

type User struct {
	Id            uuid.UUID      `json:"id" bson:"id" form:"id" gorm:"primary_key;type:uuid;"`
	Name          string         `json:"name" bson:"name" form:"name" validate:"required,havent_script,max=255" gorm:"size:255"`
	Email         string         `json:"email" bson:"email" form:"email" validate:"required,email,havent_script,max=255" gorm:"size:255;unique"`
	Ip            string         `json:"ip" bson:"ip" form:"ip" validate:"omitempty,havent_script,max=32" gorm:"size:32"`
	Password      string         `json:"password" bson:"password" form:"password" validate:"required,havent_script,max=255" gorm:"size:255"`
	Token         string         `json:"token" bson:"token" form:"token"`
	ResetToken    uuid.UUID      `json:"reset_token" bson:"reset_token" form:"reset_token" gorm:"type:uuid;"`
	ResetExpired  time.Time      `json:"reset_expired" bson:"reset_expired" form:"reset_expired"`
	AccessAt      time.Time      `json:"access_at" bson:"access_at" form:"access_at"`
	StatusId      uuid.UUID      `json:"status_id" bson:"status_id" form:"status_id" gorm:"omitempty;column:status_id;type:uuid;"`
	Status        string         `json:"status" bson:"status" form:"status" validate:"max=64" gorm:"size:64"`
	LoginAt       time.Time      `json:"login_at" bson:"login_at" form:"login_at"`
	LogoutAt      time.Time      `json:"logout_at" bson:"logout_at" form:"logout_at"`
	CreatedAt     time.Time      `json:"created_at" bson:"created_at" form:"created_at"`
	Created       string         `json:"created" bson:"created" form:"created" validate:"max=255" gorm:"size:255"`
	CreatedUserId uuid.UUID      `json:"created_user_id" bson:"created_user_id" form:"created_user_id" gorm:"type:uuid;"`
	UpdatedAt     time.Time      `json:"updated_at" bson:"updated_at" form:"updated_at"`
	Updated       string         `json:"updated" bson:"updated" form:"updated" validate:"max=255" gorm:"size:255"`
	UpdatedUserId uuid.UUID      `json:"updated_user_id" bson:"updated_user_id" form:"updated_user_id" gorm:"type:uuid;"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at" bson:"deleted_at" form:"deleted_at" gorm:"index"`
	Deleted       string         `json:"deleted" bson:"deleted" form:"deleted" validate:"max=255" gorm:"size:255"`
	DeletedUserId uuid.UUID      `json:"deleted_user_id" bson:"deleted_user_id" form:"deleted_user_id" gorm:"type:uuid;"`
}

// type UserAPI struct {
// 	Id            uuid.UUID      `json:"id"`
// 	Name          string         `json:"name"`
// 	Email         string         `json:"email"`
// 	Ip            string         `json:"-"`
// 	Password      string         `json:"-"`
// 	Token         string         `json:"-"`
// 	ResetToken    uuid.UUID      `json:"-"`
// 	ResetExpired  time.Time      `json:"-"`
// 	AccessAt      time.Time      `json:"access_at"`
// 	StatusId      int16          `json:"status_id"`
// 	Status        string         `json:"status"`
// 	LoginAt       time.Time      `json:"login_at"`
// 	LogoutAt      time.Time      `json:"logout_at"`
// 	CreatedAt     time.Time      `json:"created_at"`
// 	Created       string         `json:"created" `
// 	CreatedUserId uuid.UUID      `json:"-"`
// 	UpdatedAt     time.Time      `json:"updated_at"`
// 	Updated       string         `json:"updated"`
// 	UpdatedUserId uuid.UUID      `json:"-"`
// 	DeletedAt     gorm.DeletedAt `json:"-"`
// 	Deleted       string         `json:"-"`
// 	DeletedUserId uuid.UUID      `json:"-"`
// }
