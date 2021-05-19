package main

import (
	"log"
	"net/http"
	"rest-go-demo/controllers"
	"rest-go-demo/database"
	"rest-go-demo/entity"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql" //Required for MySQL dialect
)

func main() {
	initDB()
	log.Println("Starting the HTTP server on port 8081")

	router := mux.NewRouter().StrictSlash(true)
	initaliseHandlers(router)
	log.Fatal(http.ListenAndServe(":5000", router))
}

func initaliseHandlers(router *mux.Router) {
	router.HandleFunc("/doctors", controllers.CreateDoctor).Methods("POST")
	router.HandleFunc("/doctors", controllers.GetAllDoctors).Methods("GET")
	router.HandleFunc("/doctors/{id}", controllers.GetDoctorByID).Methods("GET")
	router.HandleFunc("/doctors/{id}", controllers.UpdateDoctorByID).Methods("PUT")
	router.HandleFunc("/doctors/{id}", controllers.DeleteDoctorByID).Methods("DELETE")
	router.HandleFunc( "/patients",controllers.GetallPatients).Methods("GET")
	router.HandleFunc("/patients/{id}",controllers.GetpatientByID).Methods("GET")
	router.HandleFunc("/patients",controllers.Createpatient).Methods("POST")
	router.HandleFunc("patients/{id}",controllers.UpdatepatientByID).Methods("PUT")
	router.HandleFunc("/patients/{id}",controllers.DeletepatientByID).Methods("DELETE")

}

func initDB() {
	config :=
		database.Config{
			ServerName: "localhost:3306",
			User:       "root",
			Password:   "11111111",
			DB:         "app",
		}

	connectionString := database.GetConnectionString(config)
	err := database.Connect(connectionString)
	if err != nil {
		panic(err.Error())
	}
	database.Migrate(&entity.Doctor{},&entity.Patient{},&entity.Role{},&entity.Privilege{},&entity.RolePrivilege{})




}
