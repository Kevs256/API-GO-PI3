package models

type Documentacion struct {
	Titulo         string     `json:"titulo"`
	Uso            string     `json:"Uso"`
	NotaImportante string     `json:"NotaImportante"`
	URL            string     `json:"url"`
	Descripcion    string     `json:"descripcion"`
	Tecnologias    []string   `json:"tecnologias"`
	Cancion1       string     `json:"cancion1"`
	Cancion2       string     `json:"cancion2"`
	Cancion3       string     `json:"cancion3"`
	IP             string     `json:"ip"`
	Puerto         string     `json:"puerto"`
	RutaInicial    string     `json:"ruta_inicial"`
	Endpoints      []Endpoint `json:"endpoints"`
	Notas          []string   `json:"notas"`
}

type Endpoint struct {
	URL               string `json:"url_ruta"`
	Metodo            string `json:"metodo"`
	Body              string `json:"body"`
	Parametro         string `json:"parametro"`
	RespuestaCorrecta string `json:"respuestaCorrecta"`
	RespuestaMala1    string `json:"respuestaMala1"`
	RespuestaMala2    string `json:"respuestaMala2"`
	Descripcion       string `json:"descripcion"`
}
