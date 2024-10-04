package server

import (
	//"backend/app/internals/handlers/auth"
	"backend/app/internals/handlers/auth"
	"backend/app/internals/models"
	"backend/app/internals/utils"
	"log"
	"net/http"
)

func LoggerMiddleware(next RouteHandler) RouteHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s %s", r.Method, r.URL.Path)
		next(w, r)
	}
}

func AuthenticationMiddleware(next RouteHandler) RouteHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		isAuthenticated := auth.CheckToken(w, r)

		if !isAuthenticated {
			data := models.Data{
				Error: "Not cookie",
			}
			utils.SendJSON(w, http.StatusUnauthorized, data)
			return
		}

		next(w, r)
	}
}

func CORSMiddleware(next RouteHandler) RouteHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next(w, r)
	}
}
