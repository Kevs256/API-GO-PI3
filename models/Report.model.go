package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Report struct {
	gorm.Model
	report_id      uuid.UUID `gorm:"type:uuid;primary_key"`
	report_subject string    `gorm:"not null;type:varchar(100)"`

	//foreignkey
	route_id string `gorm:"not null"`
}
