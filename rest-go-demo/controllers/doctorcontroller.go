package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"rest-go-demo/database"
	"rest-go-demo/entity"
	"strconv"

	"github.com/gorilla/mux"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
func GetAllDoctors(w http.ResponseWriter, r *http.Request) {
	var doctors []entity.Doctor
	database.Connector.Find(&doctors)
	enableCors(&w)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(doctors)
}

func GetDoctorByID(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	vars := mux.Vars(r)
	key := vars["id"]

	var person entity.Doctor
	database.Connector.First(&person, key)
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(person)
}

//CreatePerson creates person
func CreateDoctor(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	requestBody, _ := ioutil.ReadAll(r.Body)
	var person entity.Doctor
	json.Unmarshal(requestBody, &person)

	database.Connector.Create(person)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(person)
}

//UpdatePersonByID updates person with respective ID
func UpdateDoctorByID(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	requestBody, _ := ioutil.ReadAll(r.Body)
	var person entity.Doctor
	json.Unmarshal(requestBody, &person)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(person)
}

//DeletPersonByID delete's person with specific ID
func DeleteDoctorByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	enableCors(&w)
	var person entity.Doctor

	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	id, _ := strconv.ParseInt(key, 10, 64)
	database.Connector.Where("id = ?", id).Delete(&person)
	w.WriteHeader(http.StatusNoContent)
}
