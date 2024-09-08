package controller

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/selvamtech08/helpdesk/helper"
	"github.com/selvamtech08/helpdesk/model"
	"github.com/selvamtech08/helpdesk/store"
)

func CreateTicket(w http.ResponseWriter, r *http.Request) {
	userName := r.Header.Get("userName")
	var ticket *model.Ticket
	if err := json.NewDecoder(r.Body).Decode(&ticket); err != nil {
		helper.ErrResponse(w, http.StatusBadRequest, err)
		return
	}

	ticket.CreatedAt = time.Now()
	ticket.Author = userName
	ticket.DeadLine = time.Now().AddDate(0, 0, 3)
	if ticket.IssueType == "" {
		ticket.IssueType = "support"
	}
	if ticket.Priority == "" {
		ticket.IssueType = "low"
	}
	if ticket.Status == "" {
		ticket.Status = "YTC"
	}
	ticket, err := store.Ticket.NewTicket(ticket)
	if err != nil {
		helper.ErrResponse(w, http.StatusBadRequest, err)
		return
	}
	helper.SuccResponse(w, http.StatusOK, "ticket has been updated")
}

func GetTicket(w http.ResponseWriter, r *http.Request) {
	userName := r.Header.Get("userName")
	ticket, err := store.Ticket.GetTicketByName(userName)
	if err != nil {
		helper.ErrResponse(w, http.StatusBadRequest, err)
		return
	}
	helper.SuccResponse(w, http.StatusOK, ticket)
}

func GetAllTicket(w http.ResponseWriter, r *http.Request) {
	userName := r.Header.Get("userName")
	ticket, err := store.Ticket.GetAllTicketByName(userName)
	if err != nil {
		helper.ErrResponse(w, http.StatusBadRequest, err)
		return
	}
	helper.SuccResponse(w, http.StatusOK, ticket)
}

func UpdateTicket(w http.ResponseWriter, r *http.Request) {
	// TODO:
	helper.SuccResponse(w, http.StatusOK, "UpdateTicket")
}
