package TagServices

import (
	"api/routes/db"
	TagDTO "api/routes/modules/tags/dtos"
	TagSchema "api/routes/modules/tags/schema"
	"fmt"
)

func Generate() (string, error) {
	fmt.Println("Hola, mundo!")
	return "elefate", nil
}

//set simple tags

func ConvertToSimpleTags(tags []TagSchema.Tag) ([]TagDTO.SimpleTagDTO, error) {
	var simpleTags []TagDTO.SimpleTagDTO
	for _, tag := range tags {
		simpleTag := TagDTO.SimpleTagDTO{
			IdTag:   int(tag.ID),
			NameTag: tag.Name,
		}
		simpleTags = append(simpleTags, simpleTag)
	}
	return simpleTags, nil
}

func CreateTag(nameTag string) (*TagSchema.Tag, error) {
	result := &TagSchema.Tag{Name: nameTag}
	dbResult := db.DB.Create(result)
	if dbResult.Error != nil {
		// Ocurrió un error al crear el tag
		return nil, dbResult.Error
	}
	return result, nil
}

func GetAllTags() ([]TagSchema.Tag, error) {
	var tags []TagSchema.Tag
	dbResult := db.DB.Find(&tags)
	if dbResult.Error != nil {
		// Ocurrió un error al tomar los tags
		return nil, dbResult.Error
	}
	return tags, nil
}
