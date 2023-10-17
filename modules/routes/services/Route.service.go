package RouteServices

import (
	"api/routes/db"
	RouterDTO "api/routes/modules/routes/dtos"
	RouterSchema "api/routes/modules/routes/schema"
	"fmt"
)

func Generate() (string, error) {
	fmt.Println("Hola, mundo!")
	return "elefate", nil
}

func CreateRoute(route RouterDTO.CompleteRouteDTO) (*RouterSchema.Route, error) {
	result := &RouterSchema.Route{
		User_id:          route.UserID,
		TransportMethod:  route.TransportMethod,
		NameRoute:        route.NameRoute,
		TypeRoute:        route.TypeRoute,
		DescriptionRoute: route.DescriptionRoute,
		DurationRoute:    route.DurationRoute,
		DistanceRoute:    route.DistanceRoute,
		DateRoute:        route.DateRoute,
		LocationRoute:    route.LocationRoute,
		PriceRoute:       route.PriceRoute,
		MainImage:        route.MainImage,
		TraceRoute:       route.TraceRoute}

	dbResult := db.DB.Create(result)
	if dbResult.Error != nil {
		// Ocurrió un error al crear la ruta
		return nil, dbResult.Error
	}
	rowsAffected := dbResult.RowsAffected
	print(rowsAffected)
	return result, nil
}

func GetRouteByIdBool(id uint) bool {
	var route RouterSchema.Route
	dbResult := db.DB.First(&route, id)
	if dbResult.Error != nil {
		// Ocurrió un error al obtener la ruta
		return false
	}
	return true
}

// obtener el id del usuario por el id de la ruta
func GetUserIdByRouteIdBool(id uint) (string, error) {
	var route RouterSchema.Route
	dbResult := db.DB.First(&route, id)
	if dbResult.Error != nil {
		// Ocurrió un error al obtener la ruta
		return "null", dbResult.Error
	}
	fmt.Println(route.User_id)
	if route.User_id == "" {
		return "null", nil
	}
	return route.User_id, nil
}

//primero terminamos checkpoints para devolver los checkpoints

//func GetParcialRouteById() ([]RouterSchema.Route, error) {

//}

//func GetTotalRouteById() ([]RouterSchema.Route, error) {

//}

//get parcial por id
//get total por id
//get all parciales segun like, distancia , duracion
//todo lo de filtros
