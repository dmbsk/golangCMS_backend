package router

import (
	"net/http"

	"github.com/gorilla/mux"

	"./articleRouter"
	"./userRouter"
)

func Init() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		//w.Write([]byte("{\"Hello from the api\": \" and json\"}"))
	})
	router = userRouter.Init(router)
	router = articleRouter.Init(router)

	return router
}
