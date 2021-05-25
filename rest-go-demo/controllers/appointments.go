package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"gorm.io/gorm/clause"
	"net/http"
	"rest-go-demo/database"
	"rest-go-demo/entity"
	"strconv"
)
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	setupResponse(&w, req)
	if (*req).Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	// process the request...
}
func GetAllappointments(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	var appointments []entity.Appointment
	database.Connector.Preload(clause.Associations).Find(&appointments)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(appointments)
}

func DeleteAppointemntsByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	var appointment entity.Appointment
	id, _ := strconv.ParseInt(key, 10, 64)
	database.Connector.Where("id = ?", id).Delete(&appointment)
	w.WriteHeader(http.StatusNoContent)
}


func CreateAppointment(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	var appointment entity.Appointment
	decoder := json.NewDecoder(r.Body)
	err:=decoder.Decode(&appointment)
	if err != nil {
		panic(err)
	}
	if dbc := database.Connector.Create(&appointment); dbc.Error != nil {
		// Create failed, do something e.g. return, panic etc.
		panic(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(appointment)
}