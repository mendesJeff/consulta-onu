package controllers

import (
	"database/src/database"
	"database/src/repositories"
	"database/src/responses"
	"net/http"
	"strings"
)

func SearchOnu(w http.ResponseWriter, r *http.Request) {
	onuSerial := strings.ToLower(r.URL.Query().Get("onu"))

	db, error := database.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	repository := repositories.NewOnusRepository(db)
	onus, error := repository.SearchBySerial(onuSerial)
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusOK, onus)
}

func SearchByOlt(w http.ResponseWriter, r *http.Request) {
	oltIpOrOltName := strings.ToLower(r.URL.Query().Get("olt"))

	db, error := database.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	repository := repositories.NewOnusRepository(db)
	onus, error := repository.SearchByOlt(oltIpOrOltName)
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusOK, onus)
}
