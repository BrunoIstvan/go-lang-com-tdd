package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/BrunoIstvan/go-rest-api/database"
	"github.com/BrunoIstvan/go-rest-api/models"
	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "<h1>Home Page</h1>")

}

func AllPersonalities(w http.ResponseWriter, r *http.Request) {

	var p []models.Personality

	database.DB.Find(&p)
	json.NewEncoder(w).Encode(p)

}

func GetOnePersonality(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	var p models.Personality
	database.DB.First(&p, id)

	// se encontrou o id
	if p.Id > 0 {
		json.NewEncoder(w).Encode(p)
	} else {
		w.WriteHeader(404)
		resp := make(map[string]string)
		resp["message"] = "Resource not found"
		json.NewEncoder(w).Encode(resp)
	}

}

func CreateAPersonality(w http.ResponseWriter, r *http.Request) {

	var p models.Personality

	json.NewDecoder(r.Body).Decode(&p)
	database.DB.Create(&p)
	json.NewEncoder(w).Encode(p)
}

func DeleteAPersonality(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	var p models.Personality
	database.DB.First(&p, id)

	// se encontrou o id
	if p.Id > 0 {
		database.DB.Delete(&p, id)
		json.NewEncoder(w).Encode(p)
	} else {
		w.WriteHeader(404)
		resp := make(map[string]string)
		resp["message"] = "Resource not found"
		json.NewEncoder(w).Encode(resp)
	}

}

func UpdateAPersonality(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	var p models.Personality
	database.DB.First(&p, id)

	// se encontrou o id
	if p.Id > 0 {
		json.NewDecoder(r.Body).Decode(&p)
		database.DB.Save(&p)
		json.NewEncoder(w).Encode(p)
	} else {
		w.WriteHeader(404)
		resp := make(map[string]string)
		resp["message"] = "Resource not found"
		json.NewEncoder(w).Encode(resp)
	}

}
