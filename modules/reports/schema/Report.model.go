package ReportSchema

import (
	"gorm.io/gorm"
)

type Report struct {
	gorm.Model
	ID      uint   `gorm:"primaryKey;autoIncrement"`
	Subject string `gorm:"not null;type:varchar"`

	//foreignkey
	RouteID uint
}
