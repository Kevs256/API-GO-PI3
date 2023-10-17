package CheckPointServices

import (
	"api/routes/db"
	CheckPointDTO "api/routes/modules/checkpoints/dtos"
	CheckpoinSchema "api/routes/modules/checkpoints/schema"
	"fmt"
	"mime/multipart"
	"strings"
)

func Generate() (string, error) {
	fmt.Println("Hola, mundo!")
	return "elefate", nil
}

func CreateCheckpoint(checkpoint CheckPointDTO.ReqCompleteCheckPointDTO) (*CheckpoinSchema.CheckPoint, error) {
	result := &CheckpoinSchema.CheckPoint{
		Name:           checkpoint.Name,
		Description:    checkpoint.Description,
		MultimediaPath: checkpoint.MultimediaPath,
		Coordinates:    checkpoint.Coordinates,
		RouteID:        checkpoint.RouteID,
	}

	dbResult := db.DB.Create(result)
	if dbResult.Error != nil {
		// Ocurrió un error al crear la ruta
		return nil, dbResult.Error
	}
	rowsAffected := dbResult.RowsAffected
	print(rowsAffected)
	return result, nil
}

func CheckFileType(handler *multipart.FileHeader) bool {
	validExtensions := []string{"jpeg", "jpg", "png", "mp4"}
	fileType := strings.Split(handler.Filename, ".")[1]
	for _, extension := range validExtensions {
		if fileType == extension {
			return true
		}
	}
	return false
}

//funcion para por el nombre del archivo crear la estructura de carpetas
