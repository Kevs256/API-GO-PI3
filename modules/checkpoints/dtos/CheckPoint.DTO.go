package models

import "github.com/jinzhu/gorm/dialects/postgres"

//req dto create checkpoint
type ReqCompleteCheckPointDTO struct {
	Name           string         `json:"name"`
	Description    string         `json:"description"`
	MultimediaPath string         `json:"multimedia_path"`
	Coordinates    postgres.Jsonb `json:"coordinates"`
	RouteID        uint           `json:"route_id"`
}

//res dto create checkpoint
type ResCompleteCheckPointDTO struct {
	ID         uint   `json:"id"`
	RouteID    uint   `json:"route_id"`
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

//res dto upload file
type ResUploadFileDTO struct {
	Message    string `json:"message"`
	Path       string `json:"path"`
	StatusCode int    `json:"status_code"`
	NameFile   string `json:"name_file"`
}