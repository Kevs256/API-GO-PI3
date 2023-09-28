package routes

import (
	"api/auth/db"
	"api/auth/models"
	"api/auth/services"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// la funcion es como una funcion flecha y recibe 2 parametros
// request y response, del modulo http, el de toda la vida
// ahora mandamos por el response un .write
func Test(reponse http.ResponseWriter, request *http.Request) {
	reponse.Write([]byte("Hello World"))
}

func Register(reponse http.ResponseWriter, request *http.Request) {

	//toma datos del body y crea el objeto
	var user models.UserAuth
	json.NewDecoder(request.Body).Decode(&user)

	//estructura del mensaje
	var res struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		UserID  int    `json:"user_id"`
	}

	//verifica si el body viene completo
	if user.Email == "" || user.Password == "" {
		reponse.WriteHeader(http.StatusBadRequest)
		res.Code = int(http.StatusBadRequest)
		res.Message = "No se han ingresado las credenciales completas"
		json.NewEncoder(reponse).Encode(&res)
		return
	}

	//busca el registro con el mismo email
	findUser := db.DB.First(&user, "email = ?", user.Email).Error
	if errors.Is(findUser, gorm.ErrRecordNotFound) {
		fmt.Println("no fue encontrado")

		//crea el usuario en la base de datos, si falla la creacion avisa
		createdUser := db.DB.Create(&user)
		if createdUser.Error != nil {
			reponse.WriteHeader(http.StatusBadRequest)
			res.Code = int(http.StatusBadRequest)
			res.Message = "No se han podido crear el usuario"
			json.NewEncoder(reponse).Encode(&res)
			return
		}

		//responde con la correcta creacion
		reponse.WriteHeader(http.StatusCreated)
		res.Code = int(http.StatusCreated)
		res.UserID = int(user.ID)
		res.Message = "Creado correctamente"
		json.NewEncoder(reponse).Encode(&res)
		return

	} else {
		//responde si ya está registrado
		fmt.Println("si fue encontrado")
		reponse.WriteHeader(http.StatusBadRequest)
		res.Code = int(http.StatusBadRequest)
		res.Message = "Email ya registrado"
		json.NewEncoder(reponse).Encode(&res)
		return
	}

}

func Loggin(reponse http.ResponseWriter, request *http.Request) {

	//toma datos del body y crea el objeto
	var user models.UserAuth
	json.NewDecoder(request.Body).Decode(&user)

	//estructura del mensaje
	var res struct {
		Code      int    `json:"code"`
		Message   string `json:"message"`
		Token_jwt string `json:"token_jwt"`
	}

	//verifica si el body viene completo
	if user.Email == "" || user.Password == "" {
		reponse.WriteHeader(http.StatusBadRequest)
		res.Code = int(http.StatusBadRequest)
		res.Message = "No se han ingresado las credenciales completas"
		json.NewEncoder(reponse).Encode(&res)
		return
	}

	//buscamos si hay algun email registrado
	//guardamos el objeto en finduser
	var finduser models.UserAuth
	findUser := db.DB.First(&finduser, "email = ?", user.Email)
	if errors.Is(findUser.Error, gorm.ErrRecordNotFound) {
		fmt.Println("No fue encontrado")

		//responde si no está registrado
		reponse.WriteHeader(http.StatusBadRequest)
		res.Code = int(http.StatusBadRequest)
		res.Message = "Usuario no registrado"
		json.NewEncoder(reponse).Encode(&res)
		return

	} else {
		//reviso la contraseña
		fmt.Println("fue encontrado")
		if finduser.Password == user.Password {
			fmt.Println("contraseña correcta")

			//jwt
			fmt.Println(finduser.ID)
			tokenJwtString, err := services.GenerateJWT(finduser)
			if err != nil {
				fmt.Println("Error generando el jwt")
			}

			//responde con la correcta creacion
			reponse.WriteHeader(http.StatusAccepted)
			res.Code = int(http.StatusAccepted)
			res.Message = "Ingreso correcto"
			res.Token_jwt = tokenJwtString
			fmt.Println(res.Token_jwt)
			json.NewEncoder(reponse).Encode(&res)
			return
		} else {
			fmt.Println("Contraseña incorrecta")
			//responde si contraseña incorrecta
			reponse.WriteHeader(http.StatusBadRequest)
			res.Code = int(http.StatusBadRequest)
			res.Message = "Contaseña incorrecta"
			json.NewEncoder(reponse).Encode(&res)
			return
		}

	}

}

func Auth(reponse http.ResponseWriter, request *http.Request) {
	//toma datos del body y crea el objeto
	var token_jwt models.TokenJWT
	json.NewDecoder(request.Body).Decode(&token_jwt)
	fmt.Println("asda", token_jwt.Token_jwt)

	//estructura de la respuesta
	var res struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Id      int    `json:"id_user"`
	}

	if token_jwt.Token_jwt == "" {
		reponse.WriteHeader(http.StatusBadRequest)
		res.Code = int(http.StatusBadRequest)
		res.Message = "No ha ingresado nada bobo"
		json.NewEncoder(reponse).Encode(&res)
		return
	}

	//toma los claims de ese token, en este caso solo el id
	id_user := services.ExtracClaimsJWT(token_jwt.Token_jwt)

	//si es cero el token no es valido
	if id_user == 0 {
		reponse.WriteHeader(http.StatusBadRequest)
		res.Code = int(http.StatusBadRequest)
		res.Message = "JWT invalido"
		json.NewEncoder(reponse).Encode(&res)
		return
	} else {
		//si es valido responde con el id
		reponse.WriteHeader(http.StatusAccepted)
		res.Code = int(http.StatusAccepted)
		res.Message = "JWT valido"
		res.Id = id_user
		json.NewEncoder(reponse).Encode(&res)
		return
	}

}

func UserById(reponse http.ResponseWriter, request *http.Request) {
	//tomamos de los parametos
	id_userJson := mux.Vars(request)
	id_user := id_userJson["id_user"]
	fmt.Println(id_user)

	//estructura del mensaje
	var res struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Id      int    `json:"id_user"`
		Email   string `json:"email"`
	}

	//verificar si viene nulo por alguna extraña razon
	if id_user == "" {
		fmt.Println("viene vacio")
		reponse.WriteHeader(http.StatusBadRequest)
		res.Code = int(http.StatusBadRequest)
		res.Message = "Entrada invalida"
		json.NewEncoder(reponse).Encode(&res)
		return
	} else {
		//si no es nulo buscarlo en la db
		var finduser models.UserAuth
		findUser := db.DB.First(&finduser, "id = ?", id_user)

		if errors.Is(findUser.Error, gorm.ErrRecordNotFound) {
			fmt.Println("No fue encontrado")

			//si no fue encontrado
			reponse.WriteHeader(http.StatusBadRequest)
			res.Code = int(http.StatusBadRequest)
			res.Message = "Usuario no encontrado"
			json.NewEncoder(reponse).Encode(&res)
			return

		} else {

			//responde con lo que encontró
			reponse.WriteHeader(http.StatusAccepted)
			res.Code = int(http.StatusAccepted)
			res.Id = int(finduser.ID)
			res.Message = "Usuario encontado"
			res.Email = finduser.Email
			json.NewEncoder(reponse).Encode(&res)
			return
		}

	}
}

//func RestorePassword() {}
