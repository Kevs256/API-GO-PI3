package CheckpointRoutes

import (
	CheckPointDTO "api/routes/modules/checkpoints/dtos"
	CheckPointServices "api/routes/modules/checkpoints/services"
	RoutesServices "api/routes/modules/routes/services"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// la funcion es como una funcion flecha y recibe 2 parametros
// request y response, del modulo http, el de toda la vida
// ahora mandamos por el response un .write
func Test(reponse http.ResponseWriter, request *http.Request) {
	CheckPointServices.Generate()
	reponse.Write([]byte("HOLA MUNDO,checkpoints"))
}

func CreateCheckpoint(response http.ResponseWriter, request *http.Request) {
	var createCheckPointDTO CheckPointDTO.ReqCompleteCheckPointDTO
	json.NewDecoder(request.Body).Decode(&createCheckPointDTO)
	fmt.Println(createCheckPointDTO)
	var resCheckPointDTO CheckPointDTO.ResCompleteCheckPointDTO

	var ExistRoute = RoutesServices.GetRouteByIdBool(createCheckPointDTO.RouteID)

	if createCheckPointDTO.RouteID == 0 || !ExistRoute {
		response.WriteHeader(http.StatusBadRequest)
		resCheckPointDTO.StatusCode = int(http.StatusBadRequest)
		resCheckPointDTO.Message = "El id de la ruta es requerido o no existe"
		json.NewEncoder(response).Encode(&resCheckPointDTO)
		return
	}

	createdCheckPoint, error := CheckPointServices.CreateCheckpoint(createCheckPointDTO)
	if error != nil {
		response.WriteHeader(http.StatusBadRequest)
		resCheckPointDTO.StatusCode = int(http.StatusBadRequest)
		resCheckPointDTO.Message = "Ha ocurrido un error al crear el objeto" + error.Error()
		json.NewEncoder(response).Encode(&resCheckPointDTO)
		return
	}

	response.WriteHeader(http.StatusOK)
	resCheckPointDTO.StatusCode = int(http.StatusOK)
	resCheckPointDTO.Message = "Ruta creada correctamente"
	resCheckPointDTO.RouteID = createdCheckPoint.RouteID
	resCheckPointDTO.ID = createdCheckPoint.ID
	json.NewEncoder(response).Encode(&resCheckPointDTO)
	return
}

func UploadFileCheckPoint(response http.ResponseWriter, request *http.Request) {
	request.ParseMultipartForm(40 << 20)
	file, handler, err := request.FormFile("myFile")
	var resCheckPointDTO CheckPointDTO.ResUploadFileDTO
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		resCheckPointDTO.StatusCode = int(http.StatusBadRequest)
		resCheckPointDTO.Message = "Ha fallado la carga del archivo"
		json.NewEncoder(response).Encode(&resCheckPointDTO)
		return
	}
	if !CheckPointServices.CheckFileType(handler) {
		response.WriteHeader(http.StatusBadRequest)
		resCheckPointDTO.StatusCode = int(http.StatusBadRequest)
		resCheckPointDTO.Message = "El archivo no es de formato correcto"
		json.NewEncoder(response).Encode(&resCheckPointDTO)
		return
	}
	defer file.Close()
	//hacer servicios de creacion de carpetas recibiendo por nombre de archivo
	partsWithExtension := strings.Split(handler.Filename, ".")
	parts := strings.Split(partsWithExtension[0], "_")
	if len(parts) != 4 {
		response.WriteHeader(http.StatusBadRequest)
		resCheckPointDTO.StatusCode = int(http.StatusBadRequest)
		resCheckPointDTO.Message = "El archivo no contiene el formato correcto"
		json.NewEncoder(response).Encode(&resCheckPointDTO)
		return
	}
	// Obtener las partes de la ruta
	userId := parts[1]
	routeId := parts[2]
	checkpointId := parts[3]

	fmt.Println(userId, routeId, checkpointId)

	routeIdint, err := strconv.ParseUint(routeId, 10, 32)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		resCheckPointDTO.StatusCode = int(http.StatusBadRequest)
		resCheckPointDTO.Message = "El archivo no contiene el formato correcto, falta el id de la ruta"
		json.NewEncoder(response).Encode(&resCheckPointDTO)
		return
	}

	checkpointIdint, err := strconv.ParseUint(checkpointId, 10, 32)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		resCheckPointDTO.StatusCode = int(http.StatusBadRequest)
		resCheckPointDTO.Message = "El archivo no contiene el formato correcto, falta el id del checkpoint"
		json.NewEncoder(response).Encode(&resCheckPointDTO)
		return
	}

	UserIdObtain, err := RoutesServices.GetUserIdByRouteIdBool(uint(routeIdint))

	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		resCheckPointDTO.StatusCode = int(http.StatusBadRequest)
		resCheckPointDTO.Message = "El archivo no contiene el formato correcto, falta el id de la ruta"
		json.NewEncoder(response).Encode(&resCheckPointDTO)
		return
	}

	RouteIdObtain, err := CheckPointServices.GetIdRouteByIdCheckpointBool(uint(checkpointIdint))
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		resCheckPointDTO.StatusCode = int(http.StatusBadRequest)
		resCheckPointDTO.Message = "El archivo no contiene un id de checkpoint correcto"
		json.NewEncoder(response).Encode(&resCheckPointDTO)
		return
	}

	fmt.Println(RoutesServices.GetRouteByIdBool(uint(routeIdint)), UserIdObtain, userId, CheckPointServices.GetCheckpointByIdBool(uint(checkpointIdint)), uint(routeIdint), RouteIdObtain)
	if !RoutesServices.GetRouteByIdBool(uint(routeIdint)) || UserIdObtain != userId || RouteIdObtain != uint(routeIdint) || !CheckPointServices.GetCheckpointByIdBool(uint(checkpointIdint)) {
		response.WriteHeader(http.StatusBadRequest)
		resCheckPointDTO.StatusCode = int(http.StatusBadRequest)
		resCheckPointDTO.Message = "El archivo no contiene datos de la base de datos"
		json.NewEncoder(response).Encode(&resCheckPointDTO)
		return
	}

	pathTotal, err := CheckPointServices.CreateDirectoryPath(handler.Filename)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		resCheckPointDTO.StatusCode = int(http.StatusBadRequest)
		resCheckPointDTO.Message = "Error al crear la ruta o en el formato"
		json.NewEncoder(response).Encode(&resCheckPointDTO)
		return
	}
	pathFinal := pathTotal + "\\" + handler.Filename
	err = CheckPointServices.SaveFile(file, pathFinal)

	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		resCheckPointDTO.StatusCode = int(http.StatusBadRequest)
		resCheckPointDTO.Message = "error al guardar el archivo"
		json.NewEncoder(response).Encode(&resCheckPointDTO)
		return
	}

	err = CheckPointServices.UpdateCheckPointMultimediaPath(uint(checkpointIdint), pathFinal)

	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		resCheckPointDTO.StatusCode = int(http.StatusBadRequest)
		resCheckPointDTO.Message = "ocurrio un error al actualizar el archivo"
		json.NewEncoder(response).Encode(&resCheckPointDTO)
		return
	}

	var resUploadFileDTO CheckPointDTO.ResUploadFileDTO
	response.WriteHeader(http.StatusOK)
	resUploadFileDTO.StatusCode = int(http.StatusOK)
	resUploadFileDTO.Message = "se ha subido el archivo correctamente y guardado el path"
	resUploadFileDTO.Path = pathFinal
	resUploadFileDTO.NameFile = handler.Filename
	json.NewEncoder(response).Encode(&resUploadFileDTO)
	return
}

//obtain file for id checkpoint probar

func StreamFileCheckpoint(response http.ResponseWriter, request *http.Request) {

	checkpoint_idJson := mux.Vars(request)
	checkpoint_id := checkpoint_idJson["checkpoint_id"]
	fmt.Println(checkpoint_id)

	checkpointIDUint, err2 := strconv.ParseUint(checkpoint_id, 10, 64)

	if err2 != nil {
		var resStreamFileDTO CheckPointDTO.ResStreamFileDTO
		resStreamFileDTO.StatusCode = http.StatusBadRequest
		resStreamFileDTO.Message = "no se ha encontrado el path"
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode(&resStreamFileDTO)
		return
	}

	var path, err = CheckPointServices.GetPathImageByCheckPoinId(uint(checkpointIDUint))
	if err != nil {
		var resStreamFileDTO CheckPointDTO.ResStreamFileDTO
		resStreamFileDTO.StatusCode = http.StatusBadRequest
		resStreamFileDTO.Message = "no se ha encontrado el path"
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode(&resStreamFileDTO)
		return
	}

	file, err := os.Open(path)
	if err != nil {
		var resStreamFileDTO CheckPointDTO.ResStreamFileDTO
		resStreamFileDTO.StatusCode = http.StatusBadRequest
		resStreamFileDTO.Message = "no se ha encontrado el archivo en al ruta"
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode(&resStreamFileDTO)
		return
	}

	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		http.Error(response, "File error", http.StatusInternalServerError)
		return
	}

	http.ServeContent(response, request, file.Name(), fileInfo.ModTime(), file)
}

func GetCheckpointsByRouterId(response http.ResponseWriter, request *http.Request) {
	checkpoint_idJson := mux.Vars(request)
	checkpoint_id := checkpoint_idJson["checkpoint_id"]
	fmt.Println(checkpoint_id)

	checkpointIDUint, err2 := strconv.ParseUint(checkpoint_id, 10, 64)
	if err2 != nil {
		var resStreamFileDTO CheckPointDTO.ResStreamFileDTO
		resStreamFileDTO.StatusCode = http.StatusBadRequest
		resStreamFileDTO.Message = "no se ha encontrado el id"
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode(&resStreamFileDTO)
		return
	}
	var path, err = CheckPointServices.GetAllTotalCheckpointsByRouteId(uint(checkpointIDUint))

	if len(path) == 0 {
		var resStreamFileDTO CheckPointDTO.ResStreamFileDTO
		resStreamFileDTO.StatusCode = http.StatusBadRequest
		resStreamFileDTO.Message = "no se ha encontrado los checkpoint"
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode(&resStreamFileDTO)
		return
	}

	if err != nil {
		var resStreamFileDTO CheckPointDTO.ResStreamFileDTO
		resStreamFileDTO.StatusCode = http.StatusBadRequest
		resStreamFileDTO.Message = "no se ha encontrado el id"
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode(&resStreamFileDTO)
		return
	}

}
