package userController

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"

	. "../../dao/userDAO"
	. "../../models/userModel"
	. "../../respondFormating"
)

var dao = UsersDAO{}

func CreateUserEndPoint(w http.ResponseWriter, r *http.Request){
	defer r.Body.Close()
	var user UserModel
	buf, bodyErr := ioutil.ReadAll(r.Body)
	if bodyErr != nil {
		log.Print("bodyErr ", bodyErr.Error())
		http.Error(w, bodyErr.Error(), http.StatusInternalServerError)
		return
	}

	rdr1 := ioutil.NopCloser(bytes.NewBuffer(buf))
	rdr2 := ioutil.NopCloser(bytes.NewBuffer(buf))
	log.Printf("BODY: %q", rdr1)
	r.Body = rdr2
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if user, _ := dao.FindByUsername(user.Username); user.Username != "" {
		RespondWithError(w, http.StatusConflict, "Username is taken")
		return
	}
	//hashedPassword, hashError := HashAndCheckPassword(user.Password)

	//if hashError != nil{
	//	RespondWithError(w, http.StatusInternalServerError, hashError.Error())
	//	return
	//}
	//
	//user.Password = hashedPassword
	if user.Role == "" {
		user.Role = "reader"
	}
	user.ID = bson.NewObjectId()

	if err := dao.Insert(user); err != nil{
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
		return
	}
	if err := dao.Delete(user); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

func UpdateEndPoint(w http.ResponseWriter, r *http.Request){
	defer r.Body.Close()
	var user UserModel
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invailid request payload")
		return
	}
	if err := dao.Update(user); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

func UserSignIn(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var user UserModel
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invailid request payload")
		return
	}
	dbUser, err := dao.FindByUsername(user.Username)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Wrong username")
		return
	}
	if dbUser.Password == user.Password{
		RespondWithJson(w, http.StatusOK, dbUser)
		return
	}
	RespondWithError(w, http.StatusUnauthorized, "Wrong password or username")
	return

}