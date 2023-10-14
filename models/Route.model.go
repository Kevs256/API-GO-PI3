package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Route struct {
	gorm.Model
	ID               uuid.UUID `gorm:"type:uuid;primaryKey"`
	UserID           string
	TransportMethod  string
	TypeRoute        string
	RouteState       string
	NameRoute        string
	DescriptionRoute string
	DurationRoute    time.Time
	DistanceRoute    int
	DateRoute        time.Time
	LocationRoute    string
	PriceRoute       float64
	MainImage        string
	Likes            int
	TraceRoute       []Point `gorm:"type:point[]"`
}

type Point struct {
	X float64
	Y float64
}
