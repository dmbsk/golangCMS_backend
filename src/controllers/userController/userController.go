package userController

import (
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"gopkg.in/mgo.v2/bson"

	. "../../models/userModel"
	. "../../dao/userDAO"
	. "../../respondFormating"
)

var dao = UsersDAO{}

func CreateUserEndPoint(w http.ResponseWriter, r *http.Request){
	defer r.Body.Close()
	var user UserModel
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	user.ID = bson.NewObjectId()
	if err := dao.Insert(user); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJson(w, http.StatusCreated, user)
}

func AllUsersEndPoint(w http.ResponseWriter, r *http.Request){
	defer r.Body.Close()
	users, err := dao.FindAll()
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJson(w, http.StatusOK, users)
}

func UserEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	users, err := dao.FindById(params["id"])
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	RespondWithJson(w, http.StatusOK, users)
}

func RemoveEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var user UserModel
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invailid request payload")
	}
	if err := dao.Delete(user); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
	}
	RespondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

func UpdateEndPoint(w http.ResponseWriter, r *http.Request){
	defer r.Body.Close()
	var user UserModel
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invailid request payload")
	}
	if err := dao.Update(user); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
	}
	RespondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}
