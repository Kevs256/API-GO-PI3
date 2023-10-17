package TagDTO

//simple tag dto
type SimpleTagDTO struct {
	IdTag   int    `json:"IdTag"`
	NameTag string `json:"NameTag"`
}

//dtos usados para la creacion de tags
type ReqTagCreateDTO struct {
	NameTag string `json:"NameTag"`
}

type ResTagCreateDTO struct {
	StatusCode int    `json:"statusCode"`
	IdTag      int    `json:"IdTag"`
	NameTag    string `json:"NameTag"`
	Message    string `json:"Message"`
}

//dtos usados para la obtencion total de tags
type ResTagGetAllDTO struct {
	StatusCode int            `json:"statusCode"`
	Tags       []SimpleTagDTO `json:"Tags"`
	Message    string         `json:"Message"`
}

//
