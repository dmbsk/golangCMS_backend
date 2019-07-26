package articleDAO

import (
	"../dbDAO"

	"../../models/articleModel"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const collection = "articles"

var db *mgo.Database

func Init(mainDao dbDAO.DAO) {
	db = mainDao.Connect(collection)
}

func (m *ArticleDAO) FindAll() ([]articleModel.ArticleModel, error){
	var article []articleModel.ArticleModel
	err := db.C(collection).Find(bson.M{}).All(&article)
	return article, err
}

func (m *ArticleDAO) Insert(article articleModel.ArticleModel) error {
	err := db.C(collection).Insert(&article)
	return err
}

func (m *ArticleDAO) FindById(id string) (articleModel.ArticleModel, error) {
	var article articleModel.ArticleModel
	err := db.C(collection).FindId(bson.ObjectIdHex(id)).One(&article)
	return article, err
}

func (m *ArticleDAO) Delete(articleModel articleModel.ArticleModel) error {
	err := db.C(collection).Remove(&articleModel)
	return err
}

func (m *ArticleDAO) Update(articleModel articleModel.ArticleModel) error {
	err := db.C(collection).UpdateId(articleModel.ID, articleModel)
	return err
}