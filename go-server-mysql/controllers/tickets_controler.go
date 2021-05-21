package controllers

import (
	"encoding/json"
	"log"
	"main/commons"
	"main/models"
	"net/http"

	"github.com/gorilla/mux"
)

func GetAll(writer http.ResponseWriter, request *http.Request) {
	tickets := []models.Ticket{}

	db := commons.GetConnetion()
	defer db.Close()

	db.Find(&tickets)
	json, _ := json.Marshal(tickets)
	commons.SendReponse(writer, http.StatusOK, json)
}

func Get(writer http.ResponseWriter, request *http.Request) {
	ticket := models.Ticket{}

	id := mux.Vars(request)["id"]

	db := commons.GetConnetion()
	defer db.Close()
	db.Find(&ticket, id)

	if ticket.ID > 0 {
		json, _ := json.Marshal(ticket)
		commons.SendReponse(writer, http.StatusOK, json)
	} else {
		commons.SendError(writer, http.StatusNotFound)
	}
}

func Save(writer http.ResponseWriter, request *http.Request) {
	ticket := models.Ticket{}

	db := commons.GetConnetion()
	defer db.Close()

	err := json.NewDecoder(request.Body).Decode(&ticket)

	if err != nil {
		log.Fatal(err)
		commons.SendError(writer, http.StatusBadRequest)
		return
	}

	err = db.Save(&ticket).Error

	if err != nil {
		log.Fatal(err)
		commons.SendError(writer, http.StatusInternalServerError)
		return
	}

	json, _ := json.Marshal(ticket)

	commons.SendReponse(writer, http.StatusCreated, json)
}

func Delete(writer http.ResponseWriter, request *http.Request) {
	ticket := models.Ticket{}

	db := commons.GetConnetion()
	defer db.Close()

	id := mux.Vars(request)["id"]

	db.Find(&ticket, id)

	if ticket.ID > 0 {
		db.Delete(ticket)
		commons.SendReponse(writer, http.StatusOK, []byte(`{}`))
	} else {
		commons.SendError(writer, http.StatusNotFound)
	}
}
