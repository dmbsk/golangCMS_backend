package articleRouter

import (
	. "../../controllers/articleController"
	"github.com/gorilla/mux"
)

func Init(router *mux.Router) *mux.Router {
	router.HandleFunc("/article", AllArticlesEndPoint).Methods("GET", "OPTIONS")
	router.HandleFunc("/article", CreateArticleEndPoint).Methods("POST", "OPTIONS")
	router.HandleFunc("/article", RemoveEndPoint).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/article", UpdateEndPoint).Methods("PUT", "OPTIONS")

	router.HandleFunc("/article/{id}", ArticleEndPoint).Methods("GET", "OPTIONS")

	router.HandleFunc("/article/gallery", CreateGalleryEndPoint).Methods("POST", "OPTIONS")


	router.HandleFunc("/article/gallery/{id}/{imgName}", ImageEndPoint)

	return router
}
