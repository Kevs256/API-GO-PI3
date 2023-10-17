package RoutesRoutes

import (
	RouteDTO "api/routes/modules/routes/dtos"
	RouteServices "api/routes/modules/routes/services"
	"encoding/json"
	"net/http"
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

}

//func RestorePassword() {}
