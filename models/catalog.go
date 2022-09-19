package models

import (
	"github.com/satori/go.uuid"
	"gorm.io/gorm"
	"time"
)

type Catalog struct {
	Id            uuid.UUID      `json:"id" form:"id" gorm:"primary_key;type:uuid;"`
	Code          string         `json:"code" form:"code" validate:"required,havent_script,max=255" gorm:"size:255"`
  ShortDescription string      `json:"short_description" form:"short_description" gorm:"size:255"`
	LongDescription string       `json:"long_description" form:"long_description" gorm:"size:4096"`
  StatusId      uuid.UUID      `json:"status_id" form:"status_id" gorm:"omitempty;column:status_id;type:uuid;"`
	Status        string         `json:"status" form:"status" validate:"max=64" gorm:"size:64"`
  Notes         string         `json:"notes" form:"notes" gorm:"size:4096"`
  Uom           string         `json:"uom" form:"uom" gorm:"size:50"`
  Price         int64          `json:"price" form:"price"`
	CreatedAt     time.Time      `json:"created_at" form:"created_at"`
	Created       string         `json:"created" form:"created" validate:"max=255" gorm:"size:255"`
	CreatedUserId uuid.UUID      `json:"created_user_id" form:"created_user_id" gorm:"type:uuid;"`
	UpdatedAt     time.Time      `json:"updated_at" form:"updated_at"`
	Updated       string         `json:"updated" form:"updated" validate:"max=255" gorm:"size:255"`
	UpdatedUserId uuid.UUID      `json:"updated_user_id" form:"updated_user_id" gorm:"type:uuid;"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at" form:"deleted_at" gorm:"index"`
	Deleted       string         `json:"deleted" form:"deleted" validate:"max=255" gorm:"size:255"`
	DeletedUserId uuid.UUID      `json:"deleted_user_id" form:"deleted_user_id" gorm:"type:uuid;"`
}
