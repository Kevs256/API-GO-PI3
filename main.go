package main

import (
	"api/routes/db"
	CheckpointSchema "api/routes/modules/checkpoints/schema"
	ReportSchema "api/routes/modules/reports/schema"
	RoutesSchema "api/routes/modules/routes/schema"
	TagSchema "api/routes/modules/tags/schema"
	TagxCheckSchema "api/routes/modules/tagxcheck/schema"

	CheckpointRoutes "api/routes/modules/checkpoints/routes"
	ReportRoutes "api/routes/modules/reports/routes"
	RoutesRoutes "api/routes/modules/routes/routes"
	TagRoutes "api/routes/modules/tags/routes"
	TagxCheckRoutes "api/routes/modules/tagxcheck/routes"

	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"github.com/joho/godotenv"
)

func main() {

	//variable globales
	if err := godotenv.Load("./env"); err != nil {
		log.Fatal(err)
	}

	//tomamos variables de entorno
	IP_DB := os.Getenv("IP_DB")
	PORT_DB := os.Getenv("PORT_DB")
	USER_DB := os.Getenv("USER_DB")
	PASS_DB := os.Getenv("PASS_DB")
	NAME_DB := os.Getenv("NAME_DB")

	//vamos a conectar con la base de datos
	db.DBConnection(IP_DB, PORT_DB, USER_DB, PASS_DB, NAME_DB)

	//vamos a crear las tablas
	db.DB.AutoMigrate(ReportSchema.Report{})
	db.DB.AutoMigrate(TagxCheckSchema.TagxCheck{})

	db.DB.AutoMigrate(CheckpointSchema.CheckPoint{})
	db.DB.AutoMigrate(TagSchema.Tag{})
	db.DB.AutoMigrate(RoutesSchema.Route{})

	//acá se cre un objeto ruta del modulo mux
	router := mux.NewRouter()

	//importacion de las rutas:

	// se crean las primeras rutas
	//la funcion handlefunc lo que hace es recibir dos parametros
	//el primero es la ruta a la cual se va a dirigir
	//el segundo recibe la funcion de lo que va a responder
	//responde con una funcion

	//CHECKPOINTS
	router.HandleFunc("/checkpoint/", CheckpointRoutes.Test).Methods("GET")
	router.HandleFunc("/checkpoint/createCheckpoint", CheckpointRoutes.CreateCheckpoint).Methods("POST")
	router.HandleFunc("/checkpoint/uploadImageVideo", CheckpointRoutes.UploadFileCheckPoint).Methods("POST")

	//REPORTS
	router.HandleFunc("/report/", ReportRoutes.Test).Methods("GET")

	//ROUTES
	router.HandleFunc("/routes/", RoutesRoutes.Test).Methods("GET")
	router.HandleFunc("/routes/createRoute", RoutesRoutes.CreateRoute).Methods("POST")

	//TAGS
	router.HandleFunc("/tag/", TagRoutes.Test).Methods("GET")
	router.HandleFunc("/tag/createTag", TagRoutes.CreateTag).Methods("POST")
	router.HandleFunc("/tag/getTags", TagRoutes.GetAllTags).Methods("GET")

	//TAGXCHECK
	router.HandleFunc("/tagxCheck/", TagxCheckRoutes.Test).Methods("GET")

	//inicializamos el servidor
	//recibe el puerto y el router inicializador
	http.ListenAndServe(":8000", router)
}
