package models

import (
	CheckPointDTO "api/routes/modules/checkpoints/dtos"
	"time"

	"github.com/jinzhu/gorm/dialects/postgres"
)

// dto de la ruta completa, req
type CompleteRouteDTO struct {
	UserID           string         `json:"user_id"`
	TransportMethod  string         `json:"transport_method"`
	TypeRoute        string         `json:"type_route"`
	NameRoute        string         `json:"name_route"`
	DescriptionRoute string         `json:"description_route"`
	DurationRoute    int            `json:"duration_route"`
	DistanceRoute    int            `json:"distance_route"`
	DateRoute        time.Time      `json:"date_route"`
	LocationRoute    string         `json:"location_route"`
	PriceRoute       int            `json:"price_route"`
	TraceRoute       postgres.Jsonb `json:"trace_route"`
}

// dto de la ruta respuesta
type ResRouteCreateDTO struct {
	ID         int    `json:"id"`
	UserID     string `json:"user_id"`
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

// voy acá, falta hacer lo de checkpoint by id e imagenes para devolver la ruta
// con los checkpointss
// dto de la ruta parcial reqire
type ResParcialRouteDTO struct {
	ID              int                                     `json:"id"`
	UserID          string                                  `json:"user_id"`
	MainImage       string                                  `json:"main_image"`
	DurationRoute   int                                     `json:"duration_route"`
	DistanceRoute   int                                     `json:"distance_route"`
	PriceRoute      int                                     `json:"price_route"`
	Likes           int                                     `json:"likes"`
	ListCheckPoints []CheckPointDTO.ResParcialCheckPointDTO `json:"list_checkpoints"`
}

//dto de la ruta parcial respuesta

type ResParcialRouteErrDTO struct {
	StatusCode      int                 `json:"status_code"`
	Message         string              `json:"message"`
	ResParcialRoute *ResParcialRouteDTO `json:"res_parcial_route"`
}
