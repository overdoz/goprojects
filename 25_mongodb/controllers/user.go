package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"goprojects/25_mongodb/models"
	"net/http"
)

type UserController struct{}

func NewUserController() *UserController {
	return &UserController{}
}

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	u := models.User{
		Name: "James Bonde",
		Gender: "male",
		Age: 32,
		Id: p.ByName("id"),
	}

	//Marshal into JSON
	uj, _ := json.Marshal(u)

	//Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// composite literal - type and curly braces
	u := models.User{}

	// encode/decode for sending/ receiving JSON to/from a stream
	json.NewDecoder(r.Body).Decode(&u)

	// Change Id
	u.Id = "007"

	uj, _ := json.Marshal(u)

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "%s\n", uj)

}

func (uc UserController) DeleteUser (w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "write code to delete user\n")
}


