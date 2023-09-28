package routes

import (
	"api/auth/models"
	"encoding/json"
	"net/http"
)

func Docs(reponse http.ResponseWriter, request *http.Request) {
	documentacion := models.Documentacion{
		Titulo:         "Documentación de Api-GO-Auth - by keds256",
		Uso:            "Por favor poner una extension de visulizar JSON para visualizar correctamente",
		NotaImportante: "La maquina donde está el servicio tiene menos de una giga de memoria, por favor ser prudente al usarla, tampoco hacer registro innecesarios, en caso de necesitar añadir registro directos, o eliminar alguno comuniquese con el numero al final de la pagina",
		URL:            "https://github.com/Kevs256/Api-GO-Auth.git",
		Descripcion:    "Esta basura de json busca documentar/explicar el funcionamiento de mi api para el servicio de autenticación de usuarios.",
		Tecnologias:    []string{"go", "mux", "gorm", "postgres", "docker"},
		Cancion1:       "https://youtu.be/IoSBNZE7VSE?si=gNb8X_X_6Tzf0Ue3",
		Cancion2:       "https://youtu.be/xMTUxIBxoHk?si=NqH_IBNmZ2LjNY4j",
		Cancion3:       "https://youtu.be/VRPxao3e_jY?si=Zmf_JcaBla_R4phm",
		IP:             "Aun no definida",
		Puerto:         "Aun no definido",
		RutaInicial:    "ip:puerto/",
		Endpoints: []models.Endpoint{
			{
				URL:               "/",
				Metodo:            "GET",
				Body:              "No tiene",
				RespuestaCorrecta: "200, Hello World",
				Descripcion:       "Esta ruta es para probar la API.",
			},
			{
				URL:               "/Docs",
				Metodo:            "GET",
				Body:              "No tiene",
				RespuestaCorrecta: "200 , es este JSON",
				Descripcion:       "Una ruta hermosa que explica la API.",
			},
			{
				URL:               "/Register",
				Metodo:            "POST",
				Body:              "{'email':'ejemplo1@gmail.com', 'password':'123456789'}",
				RespuestaCorrecta: "code 201 , { 'code': 201, 'message': 'Creado correctamente', 'user_id': 13}",
				RespuestaMala1:    "code 400 , { 'code': 400, 'message': 'Email ya registrado', 'user_id': 0}",
				RespuestaMala2:    "code 400 , { 'code': 400, 'message': 'No se han ingresado las credenciales completas', 'user_id': 0}",
				Descripcion:       "Es el registro xd",
			},
			{
				URL:               "/Login",
				Metodo:            "GET",
				Body:              "{'email':'ejemplo1@gmail.com', 'password':'123456789'}",
				RespuestaCorrecta: "code 202 , { 'code': 202, 'message': 'Ingreso correcto', 'token_jwt': 'unpocodeletrasynumerosdetokenjwt'}",
				RespuestaMala1:    "code 400 , { 'code': 400, 'message': 'Usuario no registrado', 'token_jwt': ''}",
				RespuestaMala2:    "code 400 , { 'code': 400, 'message': 'Contaseña incorrecta', 'token_jwt': ''}",
				Descripcion:       "Es un login",
			},
			{
				URL:               "/Auth",
				Metodo:            "GET",
				Body:              " {'token_jwt':'unpocodeletrasynumerosdetokenjwt'}",
				RespuestaCorrecta: "code 202 , { 'code': 202, 'message': JWT valido, 'id_user' : 13}",
				RespuestaMala1:    "code 400 , { 'code':400, 'message': 'JWT invalido', 'id_user' : 0}",
				RespuestaMala2:    "code 400 , { 'code':400, 'message': 'no ha ingresado nada', 'id_user' : 0",
				Descripcion:       "valida el jwt, obtiene los claims con los cuales se ha creado y los retorna, solo he usado el id_user, se usa para verificar que la sesión está validada, uselo para obtener el id con el jwt",
			},
			{
				URL:               "/User/{id_user}",
				Metodo:            "GET",
				Body:              "no tiene",
				RespuestaCorrecta: "code 202, { 'code': 202, 'message': 'usuario encontrado', 'id_user' : 13, 'email': 'ejemplo1@gmail.com'}",
				RespuestaMala1:    "code 400, { 'code': 400, 'message': 'usuario no encontrado', 'id_user' : 0, 'email': ''}",
				Descripcion:       "Uselo para obtener el email del usuario por su id",
			},
		},
		Notas: []string{
			"Por favor, si encuentras algún error, toma una captura de pantalla y envíala a: https://api.whatsapp.com/send?phone=573177659903",
			"Que les vaya bien terminando el proyecto uwu",
			"TODOS LOS BODY'S CONSTRUIDOS CON COMILLA SENCILLA DEBEN CAMBIARSE CON COMILLA DOBLE, SOLO LOS PUSE ASI POR LIMITACIONES",
		},
	}

	json.NewEncoder(reponse).Encode(&documentacion)
}
