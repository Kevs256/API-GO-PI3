package RoutesSchemas

import (
	CheckpoinSchema "api/routes/modules/checkpoints/schema"
	ReportSchema "api/routes/modules/reports/schema"
	"time"

	"gorm.io/gorm"
)

type Route struct {
	gorm.Model
	ID               uint      `gorm:"primaryKey;autoIncrement"`
	User_id          string    `gorm:"type:varchar;not null"`
	TransportMethod  string    `gorm:"type:varchar;not null"`
	TypeRoute        string    `gorm:"type:varchar;not null"`
	RouteState       string    `gorm:"type:varchar;not null;default:active"`
	NameRoute        string    `gorm:"type:varchar"`
	DescriptionRoute string    `gorm:"type:varchar"`
	DurationRoute    int       `gorm:"type:integer"`
	DistanceRoute    int       `gorm:"type:integer;default:0"`
	DateRoute        time.Time `gorm:"type:date"`
	LocationRoute    string    `gorm:"type:varchar"`
	PriceRoute       int       `gorm:"type:integer;default:0"`
	MainImage        string    `gorm:"type:varchar"`
	Likes            int       `gorm:"type:integer;default:0"`

	//foreignKey reports
	Report []ReportSchema.Report

	//foreignKey checkpoints
	CheckPoint []CheckpoinSchema.CheckPoint

	//tipo de dato para almecenar diferentes puntos
	TraceRoute string `gorm:"type:varchar"`
}
