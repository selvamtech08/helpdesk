package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/selvamtech08/helpdesk/controller"
	"github.com/selvamtech08/helpdesk/middleware"
	"github.com/selvamtech08/helpdesk/store"
)

func main() {

	// database call
	client := store.DB
	defer client.Disconnect(context.TODO())

	// router init
	mux := http.NewServeMux()

	mux.HandleFunc("POST /user/signup", controller.SignUp)
	mux.HandleFunc("POST /user/signin", controller.SignIn)

	mux.HandleFunc("GET /user/me", middleware.AuthorizedOnly(controller.ShowMe))
	mux.HandleFunc("PUT /user", middleware.AuthorizedOnly(controller.UpdateUser))
	mux.HandleFunc("DELETE /user", middleware.AdminOnly(controller.RemoveUser))

	mux.HandleFunc("POST /ticket", middleware.AuthorizedOnly(controller.CreateTicket))
	mux.HandleFunc("GET /tickets", middleware.AuthorizedOnly(controller.GetAllTicket))
	mux.HandleFunc("GET /ticket/{id}", middleware.AuthorizedOnly(controller.GetTicket))
	mux.HandleFunc("PUT /ticket/{id}", middleware.AuthorizedOnly(controller.UpdateTicket))

	mux.HandleFunc("GET /ticket/su/{id}", middleware.SuperUserOnly(controller.GetTicketForAnalysis))
	mux.HandleFunc("GET /tickets/su", middleware.SuperUserOnly(controller.GetAllTicketForAnalysis))
	mux.HandleFunc("PUT /ticket/su", middleware.SuperUserOnly(controller.UpdateTicketByAnalysis))

	fmt.Println("server running on http://localhost:8070")
	log.Fatalln(http.ListenAndServe("localhost:8070", middleware.Logger(mux)))
}
