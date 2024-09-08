package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/selvamtech08/helpdesk/helper"
	"github.com/selvamtech08/helpdesk/store"
)

func AuthorizedOnly(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil {
			helper.ErrResponse(w, http.StatusUnauthorized, errors.New("access token missing, please loggin"))
			return
		}
		tokenString := cookie.String()
		tokenString = strings.TrimPrefix(tokenString, "token=")
		jwtToken, err := VerifyJWT(tokenString)
		if err != nil {
			helper.ErrResponse(w, http.StatusUnauthorized, errors.New("access couldn't granted"))
			return
		}
		userName, err := jwtToken.Claims.GetSubject()
		if err != nil {
			helper.ErrResponse(w, http.StatusUnauthorized, errors.New("couldn't identify the accessing user"))
			return
		}
		_, err = store.User.GetUserByName(userName)
		if err != nil {
			helper.ErrResponse(w, http.StatusUnauthorized, errors.New("invalid token, relogin again and try"))
			return
		}
		r.Header.Set("userName", userName)
		next.ServeHTTP(w, r)
	})
}

func AdminOnly(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil {
			helper.ErrResponse(w, http.StatusUnauthorized, errors.New("access token missing, please loggin"))
			return
		}
		tokenString := cookie.String()
		tokenString = strings.TrimPrefix(tokenString, "token=")
		jwtToken, err := VerifyJWT(tokenString)
		if err != nil {
			helper.ErrResponse(w, http.StatusUnauthorized, errors.New("access couldn't granted"))
			return
		}
		userRole, err := jwtToken.Claims.GetAudience()
		if err != nil || userRole[0] != "admin" {
			helper.ErrResponse(w, http.StatusUnauthorized, errors.New("you don't have access"))
			return
		}

		userName, err := jwtToken.Claims.GetSubject()
		if err != nil {
			helper.ErrResponse(w, http.StatusUnauthorized, errors.New("couldn't identify the accessing user"))
			return
		}
		r.Header.Set("userName", userName)
		next.ServeHTTP(w, r)
	})
}

func SuperUserOnly(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil {
			helper.ErrResponse(w, http.StatusUnauthorized, errors.New("access token missing, please loggin"))
			return
		}
		tokenString := cookie.String()
		tokenString = strings.TrimPrefix(tokenString, "token=")
		jwtToken, err := VerifyJWT(tokenString)
		if err != nil {
			helper.ErrResponse(w, http.StatusUnauthorized, errors.New("access couldn't granted"))
			return
		}
		userRole, err := jwtToken.Claims.GetAudience()
		if err != nil || userRole[0] != "superuser" && userRole[0] != "admin" {
			helper.ErrResponse(w, http.StatusUnauthorized, errors.New("you don't have access"))
			return
		}
		userName, err := jwtToken.Claims.GetSubject()
		if err != nil {
			helper.ErrResponse(w, http.StatusUnauthorized, errors.New("couldn't identify the accessing user"))
			return
		}
		r.Header.Set("userName", userName)
		r.Header.Set("role", userRole[0])
		next.ServeHTTP(w, r)
	})
}
