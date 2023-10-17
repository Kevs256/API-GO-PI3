package TagxCheckSchema

import (
	"gorm.io/gorm"
)

type TagxCheck struct {
	gorm.Model
	ID uint `gorm:"primaryKey;autoIncrement"`

	//foreingkey tag
	TagID uint
	//foreingkey checkpoint
	CheckPointID uint
}
