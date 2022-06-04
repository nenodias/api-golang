package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/nenodias/api-golang/database"
	"github.com/nenodias/api-golang/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
	var p []models.Personalidade
	database.DB.Find(&p)
	json.NewEncoder(w).Encode(p)
}

func RetornaUmaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var p models.Personalidade
	database.DB.First(&p, id)
	if p.Id != 0 {
		json.NewEncoder(w).Encode(p)
	} else {
		w.WriteHeader(404)
	}
}

func CriaUmaNovaPersonalidade(w http.ResponseWriter, r *http.Request) {
	var personalidade models.Personalidade
	err := json.NewDecoder(r.Body).Decode(&personalidade)
	if err != nil {
		w.WriteHeader(400)
		log.Fatal(err.Error())
	} else {
		database.DB.Create(&personalidade)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(201)
		json.NewEncoder(w).Encode(personalidade)

	}
}

func DeletaUmaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var p models.Personalidade
	database.DB.First(&p, id)
	if p.Id != 0 {
		database.DB.Delete(&p, id)
		w.WriteHeader(204)
	} else {
		w.WriteHeader(404)
	}
}

func EditaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err == nil {
		var p models.Personalidade
		database.DB.First(&p, id)
		json.NewDecoder(r.Body).Decode(&p)
		p.Id = id
		database.DB.Save(&p)
		json.NewEncoder(w).Encode(p)
	} else {
		w.WriteHeader(400)
	}
}
