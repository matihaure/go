package controller

import (
	"encoding/json"
	"net/http"

	"github.com/matihaure/model"
)

func dbtest(w http.ResponseWriter, r *http.Request) {

	records, err := model.DBTest()

	if err != nil {
		w.Write([]byte("NO DEVUELVE LAS BASE DE DATOS"))
		return
	}
	json.NewEncoder(w).Encode(records)

	w.WriteHeader(http.StatusOK)

}

func enginedbtest(w http.ResponseWriter, r *http.Request) {

	db, err := model.Connect()

	if err != nil {
		w.Write([]byte("ERROR DE CONECCION CON LA BASE"))
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	json.NewEncoder(w).Encode(db)

	w.WriteHeader(http.StatusOK)

}
