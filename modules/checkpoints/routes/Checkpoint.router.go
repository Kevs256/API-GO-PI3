package CheckpointRoutes

import (
	CheckPointDTO "api/routes/modules/checkpoints/dtos"
	CheckPointServices "api/routes/modules/checkpoints/services"
	RoutesServices "api/routes/modules/routes/services"
	"encoding/json"
	"fmt"
	"net/http"
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
	fmt.Println(handler.Filename, err)
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
	fmt.Fprintf(response, "Nombre del archivo: %v\n", handler.Filename)
}
