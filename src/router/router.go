package router

import (
	"github.com/gorilla/mux"

	"./userRouter"
	"./authRouter"
	)

func Init() *mux.Router {
	router := mux.NewRouter()
	router = authRouter.Init(router)
	router = userRouter.Init(router)

	return router
}
