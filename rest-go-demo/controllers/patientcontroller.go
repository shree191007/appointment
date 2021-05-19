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


func GetallPatients (w http.ResponseWriter, r *http.Request){
	var patients []entity.Patient
	database.Connector.Find(&patients)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(patients)

}
func GetpatientByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	var patients entity.Patient
	database.Connector.First(&patients, key)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(patients)
}
//CreatePerson creates person
func Createpatient(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var patient entity.Patient
	json.Unmarshal(requestBody, &patient)

	database.Connector.Create(patient)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(patient)
}

//UpdatePersonByID updates person with respective ID
func UpdatepatientByID(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var patient entity.Doctor
	json.Unmarshal(requestBody, &patient)
	database.Connector.Save(&patient)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(patient)
}

//DeletPersonByID delete's person with specific ID
func DeletepatientByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	var patient entity.Patient
	id, _ := strconv.ParseInt(key, 10, 64)
	database.Connector.Where("id = ?", id).Delete(&patient)
	w.WriteHeader(http.StatusNoContent)
}
