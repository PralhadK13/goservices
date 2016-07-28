// main
// Simple API server

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Somnath004/goservices/handlers"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	s := setroutes(router)
	port := ":8888"
	log.Printf("API server listening at %s", port)
	s.Run(port)
}

// Setup server routes
// Setting server datapoints handllers routes
func setroutes(router *mux.Router) *negroni.Negroni {

	usr := handler.NewUserController()
	userLogin := new(handler.UserController)
	router.HandleFunc("/", welcome).Methods("GET")
	router.HandleFunc("/login", userLogin.Login).Methods("POST")
	router.HandleFunc("/user/{id}", usr.Get).Methods("GET")
	router.HandleFunc("/registration", usr.Set).Methods("POST")

	msg := handler.NewMsgController()
	//Messages
	router.HandleFunc("/getmessage/{id}", msg.Get).Methods("GET")
	router.HandleFunc("/setmessage", msg.Set).Methods("POST")
	cor := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "OPTIONS", "DELETE", "CONNECT"},
		AllowedHeaders: []string{"*"},
	})

	n := negroni.New(
		cor,
	)

	n.UseHandler(router)
	return n
}

func welcome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Welcome to Test API server.\n Health: Ok")
}
