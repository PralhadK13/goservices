// main
// Simple API server

package main

import (
	"fmt"
	"goserv/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	setroutes(router)
	port := ":8888"
	log.Printf("API server listening at %s", port)
	log.Fatal(http.ListenAndServe(port, router))
}

// Setup server routes
// Setting server datapoints handllers routes
func setroutes(router *mux.Router) {

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

}

func welcome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Welcome to Test API server.\n Health: Ok")
}
