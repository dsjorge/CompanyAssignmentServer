package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"company/dbsql"
	"company/models"

	"github.com/julienschmidt/httprouter"
)

var dbInstance = dbsql.InitDb()

func handleErrorsResponseWrite(err error, w http.ResponseWriter) bool {
	result := true
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return result
	}

	return false
}

func GetAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	result, err := dbsql.GetAll(dbInstance)
	if handleErrorsResponseWrite(err, w) {
		return
	}
	js, err := json.Marshal(result)
	if handleErrorsResponseWrite(err, w) {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func GetAllResume(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	result, err := dbsql.GetAllResume(dbInstance)
	if handleErrorsResponseWrite(err, w) {
		return
	}
	js, err := json.Marshal(result)
	if handleErrorsResponseWrite(err, w) {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func GetById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	idIn := p.ByName("id")
	id, err := strconv.Atoi(idIn)
	if handleErrorsResponseWrite(err, w) {
		return
	}
	result, err := dbsql.GetById(dbInstance, id)
	if handleErrorsResponseWrite(err, w) {
		return
	}
	js, err := json.Marshal(result)
	if handleErrorsResponseWrite(err, w) {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func AddCompany(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	company := models.Company{}
	json.NewDecoder(r.Body).Decode(&company)
	err := dbsql.AddCompany(dbInstance, company)
	if handleErrorsResponseWrite(err, w) {
		return
	}
	w.WriteHeader(201)
}

func UpdateCompany(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	company := models.Company{}
	json.NewDecoder(r.Body).Decode(&company)
	err := dbsql.UpdateCompany(dbInstance, company)
	if handleErrorsResponseWrite(err, w) {
		return
	}
	w.WriteHeader(201)
}

func GetAvailableOwners(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	idIn := p.ByName("id")
	id, err := strconv.Atoi(idIn)
	if handleErrorsResponseWrite(err, w) {
		return
	}
	result, err := dbsql.GetAvailableOwners(dbInstance, id)
	if handleErrorsResponseWrite(err, w) {
		return
	}
	js, err := json.Marshal(result)
	if handleErrorsResponseWrite(err, w) {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
