package TagSchema

import (
	TagxCheckSchema "api/routes/modules/tagxcheck/schema"

	"gorm.io/gorm"
)

type Tag struct {
	gorm.Model
	ID   uint   `gorm:"primaryKey;autoIncrement"`
	Name string `gorm:"not null; type: varchar; unique"`

	//foreingkey tagxcheck
	TagxCheck []TagxCheckSchema.TagxCheck
}
