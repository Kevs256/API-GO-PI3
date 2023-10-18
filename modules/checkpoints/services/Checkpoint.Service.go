package CheckPointServices

import (
	"api/routes/db"
	CheckPointDTO "api/routes/modules/checkpoints/dtos"
	CheckpoinSchema "api/routes/modules/checkpoints/schema"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
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

func CreateDirectoryPath(fileName string) (string, error) {
	partsWithExtension := strings.Split(fileName, ".")
	parts := strings.Split(partsWithExtension[0], "_")
	if len(parts) != 4 {
		return "", fmt.Errorf("El nombre del archivo no sigue el formato esperado")
	}
	// Obtener las partes de la ruta
	userId := parts[1]
	routeId := parts[2]
	checkpointId := parts[3]

	uploadDir := "./uploads"
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		os.Mkdir(uploadDir, 0777)
	}

	userDir := filepath.Join(uploadDir, userId)
	if _, err := os.Stat(userDir); os.IsNotExist(err) {
		os.Mkdir(userDir, 0777)
	}

	// Crear directorio de ruta si no existe
	routeDir := filepath.Join(userDir, routeId)
	if _, err := os.Stat(routeDir); os.IsNotExist(err) {
		os.Mkdir(routeDir, 0777)
	}

	checkpointDir := filepath.Join(routeDir, checkpointId)
	if _, err := os.Stat(checkpointDir); os.IsNotExist(err) {
		os.Mkdir(checkpointDir, 0777)
	}

	finalPath := checkpointDir
	return finalPath, nil
}

func GetCheckpointByIdBool(id uint) bool {
	var checkpoint CheckpoinSchema.CheckPoint
	dbResult := db.DB.First(&checkpoint, id)
	if dbResult.Error != nil {
		// Ocurrió un error al obtener la ruta
		return false
	}
	return true
}

func GetIdRouteByIdCheckpointBool(checkpointID uint) (uint, error) {
	var checkpoint CheckpoinSchema.CheckPoint
	dbResult := db.DB.First(&checkpoint, checkpointID)
	if dbResult.Error != nil {
		// Ocurrió un error al obtener el checkpoint
		return 0, dbResult.Error
	}
	return checkpoint.RouteID, nil
}

func SaveFile(file multipart.File, path string) error {
	defer file.Close()

	// Abrir un nuevo archivo en el path especificado
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0777)
	fmt.Println(path)
	if err != nil {
		return err
	}
	defer f.Close()

	// Copiar el contenido del archivo subido al archivo de destino
	_, err = io.Copy(f, file)
	if err != nil {
		return err
	}

	return nil
}

func UpdateCheckPointMultimediaPath(checkpointID uint, multimediaPath string) error {
	var checkpoint CheckpoinSchema.CheckPoint
	result := db.DB.First(&checkpoint, checkpointID)
	if result.Error != nil {
		return result.Error
	}

	checkpoint.MultimediaPath = multimediaPath
	result = db.DB.Save(&checkpoint)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

//get all checkpoints by id route

func GetAllParcialCheckpointsByRouteId(routeID uint) ([]CheckPointDTO.ResParcialCheckPointDTO, error) {
	var checkpoints []CheckpoinSchema.CheckPoint
	result := db.DB.Find(&checkpoints, "route_id = ?", routeID)
	if result.Error != nil {
		return nil, result.Error
	}

	var resParcialCheckpoints []CheckPointDTO.ResParcialCheckPointDTO
	for _, checkpoint := range checkpoints {
		resParcialCheckpoint := CheckPointDTO.ResParcialCheckPointDTO{
			ID:             checkpoint.ID,
			MultimediaPath: checkpoint.MultimediaPath,
		}
		resParcialCheckpoints = append(resParcialCheckpoints, resParcialCheckpoint)
	}
	return resParcialCheckpoints, nil
}

func GetParcialCheckPointByCheckPointId(checkpointID uint) (*CheckPointDTO.ResParcialCheckPointDTO, error) {
	var checkpoint CheckpoinSchema.CheckPoint
	result := db.DB.First(&checkpoint, checkpointID)
	if result.Error != nil {
		return nil, result.Error
	}

	parcialCheckpoint := &CheckPointDTO.ResParcialCheckPointDTO{
		ID:             checkpoint.ID,
		MultimediaPath: checkpoint.MultimediaPath,
	}

	return parcialCheckpoint, nil
}

//getImageBypath
