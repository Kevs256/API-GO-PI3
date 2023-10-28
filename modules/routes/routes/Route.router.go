package RoutesRoutes

import (
	CheckPointServices "api/routes/modules/checkpoints/services"
	RouteDTO "api/routes/modules/routes/dtos"
	RouteServices "api/routes/modules/routes/services"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// la funcion es como una funcion flecha y recibe 2 parametros
// request y response, del modulo http, el de toda la vida
// ahora mandamos por el response un .write
func Test(reponse http.ResponseWriter, request *http.Request) {
	RouteServices.Generate()
	reponse.Write([]byte("HOLA MUNDO, routes"))
}

func CreateRoute(response http.ResponseWriter, request *http.Request) {
	var createRouteDTO RouteDTO.CompleteRouteDTO
	json.NewDecoder(request.Body).Decode(&createRouteDTO)
	var resRouteDTO RouteDTO.ResRouteCreateDTO

	if createRouteDTO.UserID == "" {
		response.WriteHeader(http.StatusBadRequest)
		resRouteDTO.StatusCode = int(http.StatusBadRequest)
		resRouteDTO.Message = "El id del usuarioes requerido"
		json.NewEncoder(response).Encode(&resRouteDTO)
		return
	}

	createdRoute, error := RouteServices.CreateRoute(createRouteDTO)
	if error != nil {
		response.WriteHeader(http.StatusBadRequest)
		resRouteDTO.StatusCode = int(http.StatusBadRequest)
		resRouteDTO.Message = "Ha ocurrido un error al crear el objeto" + error.Error()
		json.NewEncoder(response).Encode(&resRouteDTO)
		return
	}
	response.WriteHeader(http.StatusOK)
	resRouteDTO.StatusCode = int(http.StatusOK)
	resRouteDTO.Message = "Ruta creada correctamente"
	resRouteDTO.UserID = createdRoute.User_id
	resRouteDTO.ID = int(createdRoute.ID)
	json.NewEncoder(response).Encode(&resRouteDTO)
	return
}

func GetRoutesByUserId(response http.ResponseWriter, request *http.Request) {

}

func GetRouteByRouteId(response http.ResponseWriter, request *http.Request) {
	route_idJson := mux.Vars(request)
	route_id := route_idJson["route_id"]
	fmt.Println(route_id)

	var resRouteDTO RouteDTO.ResParcialRouteErrDTO

	if route_id == "" {
		response.WriteHeader(http.StatusBadRequest)
		resRouteDTO.StatusCode = int(http.StatusBadRequest)
		resRouteDTO.Message = "El id del usuarioes requerido"
		json.NewEncoder(response).Encode(&resRouteDTO)
		return
	}
	checkpointIdint, err := strconv.ParseUint(route_id, 10, 32)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		resRouteDTO.StatusCode = int(http.StatusBadRequest)
		resRouteDTO.Message = "formato incorrecto de id"
		json.NewEncoder(response).Encode(&resRouteDTO)
		return
	}

	parcialRouteDTO, err := RouteServices.GetParcialRoutesByRouteId(uint(checkpointIdint))

	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		resRouteDTO.StatusCode = int(http.StatusBadRequest)
		resRouteDTO.Message = "la ruta no existe"
		json.NewEncoder(response).Encode(&resRouteDTO)
		return
	}

	parcialRoute, err := RouteServices.GetParcialRoutesByRouteId(uint(checkpointIdint))

	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		resRouteDTO.StatusCode = int(http.StatusBadRequest)
		resRouteDTO.Message = "no se encontró la ruta"
		json.NewEncoder(response).Encode(&resRouteDTO)
		return
	}

	CheckPointsOfRoute, err := CheckPointServices.GetAllParcialCheckpointsByRouteId(uint(parcialRoute.ID))

	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		resRouteDTO.StatusCode = int(http.StatusBadRequest)
		resRouteDTO.Message = "no se ecncontrarón checkpoints de la ruta, db"
		json.NewEncoder(response).Encode(&resRouteDTO)
		return
	}

	if len(CheckPointsOfRoute) == 0 {
		response.WriteHeader(http.StatusBadRequest)
		resRouteDTO.StatusCode = int(http.StatusBadRequest)
		resRouteDTO.Message = "no se ecncontrarón checkpoints de la ruta"
		json.NewEncoder(response).Encode(&resRouteDTO)
		return
	}

	parcialRouteDTO.ListCheckPoints = CheckPointsOfRoute
	resRouteDTO.ResParcialRoute = parcialRouteDTO
	response.WriteHeader(http.StatusOK)
	resRouteDTO.StatusCode = int(http.StatusOK)
	resRouteDTO.Message = "Ruta encontrada correctamente"
	json.NewEncoder(response).Encode(&resRouteDTO)
	return
}

//func RestorePassword() {}
