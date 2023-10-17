package TagRoutes

import (
	TagDTO "api/routes/modules/tags/dtos"
	TagServices "api/routes/modules/tags/services"
	"encoding/json"
	"net/http"
)

// la funcion es como una funcion flecha y recibe 2 parametros
// request y response, del modulo http, el de toda la vida
// ahora mandamos por el response un .write
func Test(reponse http.ResponseWriter, request *http.Request) {
	TagServices.Generate()
	reponse.Write([]byte("HOLA MUNDO, tags"))
}

func CreateTag(reponse http.ResponseWriter, request *http.Request) {
	//declaramos una variable tipo TagDTO.ReqTagCreateDTO para recibir
	var createTagDTO TagDTO.ReqTagCreateDTO
	//decodificamos el body del request
	json.NewDecoder(request.Body).Decode(&createTagDTO)
	//declaramos una variable tipo TagDTO.ResTagCreateDTO para responder
	var resTagDTO TagDTO.ResTagCreateDTO
	//si el nombre del tag es vacio respondemos con un status 400 y un mensaje
	if createTagDTO.NameTag == "" {
		reponse.WriteHeader(http.StatusBadRequest)
		resTagDTO.StatusCode = int(http.StatusBadRequest)
		resTagDTO.Message = "El nombre del tag es requerido"
		json.NewEncoder(reponse).Encode(&resTagDTO)
		return
	}
	//si el nombre del tag es distinto de vacio, se crea el tag
	createdTag, error := TagServices.CreateTag(createTagDTO.NameTag)
	if error != nil {
		reponse.WriteHeader(http.StatusBadRequest)
		resTagDTO.StatusCode = int(http.StatusBadRequest)
		resTagDTO.Message = "Ha ocurrido un error al crear el objeto" + error.Error()
		json.NewEncoder(reponse).Encode(&resTagDTO)
		return
	}
	//si el tag se crea correctamente respondemos con un status 200 y el tag
	reponse.WriteHeader(http.StatusOK)
	resTagDTO.StatusCode = int(http.StatusOK)
	resTagDTO.Message = "Tag creado correctamente"
	resTagDTO.IdTag = int(createdTag.ID)
	resTagDTO.NameTag = createdTag.Name
	json.NewEncoder(reponse).Encode(&resTagDTO)
	return
}

func GetAllTags(reponse http.ResponseWriter, request *http.Request) {
	var resTagDTO TagDTO.ResTagGetAllDTO
	allTags, error := TagServices.GetAllTags()
	if error != nil {
		reponse.WriteHeader(http.StatusBadRequest)
		resTagDTO.StatusCode = int(http.StatusBadRequest)
		resTagDTO.Message = "Ha ocurrido un error al crear el objeto" + error.Error()
		json.NewEncoder(reponse).Encode(&resTagDTO)
		return
	}

	reponse.WriteHeader(http.StatusOK)
	resTagDTO.StatusCode = int(http.StatusOK)
	resTagDTO.Message = "Tags obtenidos con exito"
	simpleTags, error := TagServices.ConvertToSimpleTags(allTags)
	resTagDTO.Tags = simpleTags

	json.NewEncoder(reponse).Encode(&resTagDTO)
	return
}
