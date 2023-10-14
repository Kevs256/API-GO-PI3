package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Tag struct {
	gorm.Model
	tag_id   uuid.UUID `gorm:"type:uuid;primary_key"`
	tag_name string    `gorm:"not null"`
}
