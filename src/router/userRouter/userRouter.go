package userRouter

import (
	"github.com/gorilla/mux"
	. "../../controllers/userController"
)

func Init(router *mux.Router) *mux.Router {
	router.HandleFunc("/users", AllUsersEndPoint).Methods("GET")
	router.HandleFunc("/users", CreateUserEndPoint).Methods("POST")
	router.HandleFunc("/users", RemoveEndPoint).Methods("DELETE")
	router.HandleFunc("/users", UpdateEndPoint).Methods("PUT")

	router.HandleFunc("/users/{id}", UserEndPoint).Methods("GET")

	return router
}