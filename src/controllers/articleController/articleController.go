package articleController

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"

	. "../../dao/articleDAO"
	. "../../models/articleModel"
	. "../../respondFormating"
)

var dao = ArticleDAO{}

func CreateArticleEndPoint(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	var article ArticleModel
	if err := json.NewDecoder(r.Body).Decode(&article); err != nil {
		fmt.Println(err)
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	article.ID = bson.NewObjectId()
	article.IsPublic = false
	article.Date = time.Now()

	var articleGallery Gallery
	articleGallery.Id = bson.NewObjectId()
	articleGallery.ArticleId = article.ID

	article.Gallery = articleGallery

	if err := dao.Insert(article); err != nil{
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithJson(w, http.StatusCreated, article)
}

func AllArticlesEndPoint(w http.ResponseWriter, r *http.Request){
	defer r.Body.Close()
	articles, err := dao.FindAll()
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJson(w, http.StatusOK, articles)
}

func ArticleEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	articles, err := dao.FindById(params["id"])
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	RespondWithJson(w, http.StatusOK, articles)
}

func RemoveEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var article ArticleModel
	if err := json.NewDecoder(r.Body).Decode(&article); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invailid request payload")
	}
	if err := dao.Delete(article); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
	}
	RespondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

func UpdateEndPoint(w http.ResponseWriter, r *http.Request){
	defer r.Body.Close()
	var article ArticleModel
	if err := json.NewDecoder(r.Body).Decode(&article); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invailid request payload")
	}
	if err := dao.Update(article); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
	}
	RespondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}
