package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/selvamtech08/helpdesk/helper"
	"github.com/selvamtech08/helpdesk/middleware"
	"github.com/selvamtech08/helpdesk/model"
	"github.com/selvamtech08/helpdesk/store"
)

func setCookie(accessToken string) *http.Cookie {
	return &http.Cookie{
		Name:  "token",
		Value: accessToken,
		Path:  "/",
		// Domain:   "", // Optional
		MaxAge:   3600, // Optional: Set the cookie's expiration time in seconds
		Secure:   true, // Optional: Set to true for HTTPS-only cookies
		HttpOnly: true, // Optional: Set to true to prevent JavaScript access
	}
}

// register new user in the system
func SignUp(w http.ResponseWriter, r *http.Request) {
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		helper.ErrResponse(w, http.StatusBadRequest, helper.ErrBadInputRequest)
		return
	}
	defer r.Body.Close()

	// password must provided
	if user.Password == "" || user.Name == "" || user.Email == "" {
		helper.ErrResponse(w, http.StatusBadRequest, helper.ErrBadInputRequest)
		return
	}

	hashPassword, err := store.User.HashPassword(user.Password)
	if err != nil {
		helper.ErrResponse(w, http.StatusInternalServerError, helper.ErrDBUpdateFailed)
		return
	}
	user.Password = hashPassword
	user.CreatedAt = time.Now()
	user.IsActive = true
	user.Role = "user"
	err = store.User.New(user)
	if err != nil {
		helper.ErrResponse(w, http.StatusInternalServerError, err)
		return
	}
	helper.SuccResponse(w, http.StatusCreated, "user added successfully!")
}

// sign in to access the api service
func SignIn(w http.ResponseWriter, r *http.Request) {
	var creden model.LogIn
	if err := json.NewDecoder(r.Body).Decode(&creden); err != nil {
		helper.ErrResponse(w, http.StatusBadRequest, err)
		return
	}
	defer r.Body.Close()

	user, err := store.User.GetUserByName(creden.UserName)
	if err != nil {
		helper.ErrResponse(w, http.StatusInternalServerError, helper.ErrCredentialNotMatch)
		return
	}
	if err := store.User.VeriftPassword(creden.Password, user.Password); err != nil {
		helper.ErrResponse(w, http.StatusInternalServerError, helper.ErrCredentialNotMatch)
		return
	}
	accessToken, err := middleware.GenerateJWT(user.Name, user.Role)
	if err != nil {
		helper.ErrResponse(w, http.StatusInternalServerError, err)
		return
	}
	http.SetCookie(w, setCookie(accessToken))
	helper.SuccResponse(w, http.StatusOK, fmt.Sprintf("Hi %s, have a nice day!", user.Name))
}

// show current user information
func ShowMe(w http.ResponseWriter, r *http.Request) {
	name := r.Header.Get("userName")
	user, err := store.User.GetUserByName(name)
	if err != nil {
		helper.ErrResponse(w, http.StatusInternalServerError, err)
		return
	}
	user.Password = ""
	helper.SuccResponse(w, http.StatusOK, user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	name := r.Header.Get("userName")
	var updateUser model.UpdateUser
	if err := json.NewDecoder(r.Body).Decode(&updateUser); err != nil {
		log.Println("UpdateUser: json decode failed")
		helper.ErrResponse(w, http.StatusBadRequest, err)
		return
	}
	err := store.User.UpdateUser(name, updateUser)
	if err != nil {
		log.Println("UpdateUser: db update failed")
		helper.ErrResponse(w, http.StatusBadRequest, err)
		return
	}
	helper.SuccResponse(w, http.StatusOK, updateUser)
}

func RemoveUser(w http.ResponseWriter, r *http.Request) {
	helper.SuccResponse(w, http.StatusOK, "RemoveUser")
}
