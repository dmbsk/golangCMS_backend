package authRouter

import (
	"github.com/gorilla/mux"

	. "../../controllers/authController"
)

func Init(router *mux.Router) *mux.Router{
	router.HandleFunc("/auth", CreateTokenEndPoint).Methods("POST")
	router.HandleFunc("/protected", ProtectedEndPoint).Methods("POST")

	router.HandleFunc("/auth-test", ValidateMiddleware(TestEndpoint)).Methods("GET")

	return router
}