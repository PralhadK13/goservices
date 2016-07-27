package handler

import (
	"encoding/json"
	"goserv/db"
	"goserv/models"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type UserController struct{}

func NewUserController() Handler {
	return new(UserController)
}

func (this *UserController) Set(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	data := new(db.DB)
	var usr model.User
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	if err = json.Unmarshal(b, &usr); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := data.RegisterUser(&usr); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

}

func (this *UserController) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	data := new(db.DB)

	var lgn model.LoginRequest
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	if err = json.Unmarshal(b, &lgn); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	resp, err := data.Login(&lgn)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	b, _ = json.MarshalIndent(resp, "", " ")
	w.Write(b)
}

func (this *UserController) Get(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	data := new(db.DB)
	vars := mux.Vars(r)
	id := vars["id"]
	resp, err := data.GetUser(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	b, _ := json.MarshalIndent(resp, "", " ")
	w.Write(b)
}
