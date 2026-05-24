package middleware

import (
	"log"
	"net/http"
	"social-plus/src/auth"
	"social-plus/src/responses"
)

//Logger apply messages in the terminal
func Logger(nextFunction http.HandlerFunc) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		nextFunction(w, r)
	}
}

//Authenticate verify if the request is authenticating...
func Authenticate( nextFunction http.HandlerFunc ) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if erro := auth.ValidateToken(r); erro != nil {
			responses.ERROR(w, http.StatusUnauthorized, erro)
			return
		}
		nextFunction(w, r)
	}
	
}