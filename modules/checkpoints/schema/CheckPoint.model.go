package CheckpoinSchema

import (
	TagxCheckSchema "api/routes/modules/tagxcheck/schema"

	"gorm.io/gorm"
)

type CheckPoint struct {
	gorm.Model
	ID             uint   `gorm:"primaryKey;autoIncrement"`
	Name           string `gorm:"type:varchar;not null"`
	Description    string `gorm:"type:varchar"`
	MultimediaPath string `gorm:"type:varchar"`
	Coordinates    string `gorm:"type:varchar;not null"`

	//foreingkey id for router
	RouteID uint

	//foreigntagxcheck
	TagxCheck []TagxCheckSchema.TagxCheck
}
