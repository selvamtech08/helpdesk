package controller

import (
	"encoding/json"
	"net/http"

	"github.com/selvamtech08/helpdesk/helper"
	"github.com/selvamtech08/helpdesk/model"
	"github.com/selvamtech08/helpdesk/store"
)

func GetTicketForAnalysis(w http.ResponseWriter, r *http.Request) {
	// userName := r.Header.Get("userName")
	id := r.PathValue("id")
	ticket, err := store.Ticket.GetTicketByID(id)
	if err != nil {
		helper.ErrResponse(w, http.StatusBadRequest, err)
		return
	}
	helper.SuccResponse(w, http.StatusOK, ticket)
}

func GetAllTicketForAnalysis(w http.ResponseWriter, r *http.Request) {
	userName := r.Header.Get("userName")
	ticket, err := store.Ticket.GetAllTicketByName(userName)
	if err != nil {
		helper.ErrResponse(w, http.StatusBadRequest, err)
		return
	}
	helper.SuccResponse(w, http.StatusOK, ticket)
}

func UpdateTicketByAnalysis(w http.ResponseWriter, r *http.Request) {
	superUserName := r.Header.Get("userName")
	var ticket model.UpdateTicket
	if err := json.NewDecoder(r.Body).Decode(&ticket); err != nil {
		helper.ErrResponse(w, http.StatusBadRequest, helper.ErrBadInputRequest)
		return
	}
	if ticket.ID == "" {
		helper.ErrResponse(w, http.StatusBadRequest, helper.ErrTicketIDMissing)
		return
	}
	err := store.Ticket.UpdateTicket(superUserName, ticket)
	if err != nil {
		helper.ErrResponse(w, http.StatusBadRequest, err)
		return
	}
	helper.SuccResponse(w, http.StatusOK, ticket)
}
