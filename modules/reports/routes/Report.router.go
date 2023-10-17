package ReportRoutes

import (
	ReportServices "api/routes/modules/routes/services"
	"net/http"
)

// la funcion es como una funcion flecha y recibe 2 parametros
// request y response, del modulo http, el de toda la vida
// ahora mandamos por el response un .write
func Test(reponse http.ResponseWriter, request *http.Request) {
	ReportServices.Generate()
	reponse.Write([]byte("HOLA MUNDO, reports"))
}

//func RestorePassword() {}
