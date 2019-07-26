package userDAO

import (
	"log"

	"../../models/userModel"
	"../dbDAO"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const collection = "users"


var db *mgo.Database

func Init(mainDao dbDAO.DAO) {
	db = mainDao.Connect(collection)
}

func (m *UsersDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	} else {
		println("Mongo db started without respond")
	}
	db = session.DB(m.Database)
}

func (m *UsersDAO) FindAll() ([]userModel.UserModel, error){
	var users []userModel.UserModel
	err := db.C(collection).Find(bson.M{}).All(&users)
	return users, err
}

func (m *UsersDAO) Insert(user userModel.UserModel) error {
	err := db.C(collection).Insert(&user)
	return err
}

func (m *UsersDAO) FindById(id string) (userModel.UserModel, error) {
	var user userModel.UserModel
	err := db.C(collection).FindId(bson.ObjectIdHex(id)).One(&user)
	return user, err
}

func (m *UsersDAO) FindByUsername(username string) (userModel.UserModel, error) {
	var user userModel.UserModel
	err := db.C(collection).Find(bson.M{"username": username}).One(&user)
	return user, err
}

func (m *UsersDAO) Delete(userModel userModel.UserModel) error {
	err := db.C(collection).Remove(&userModel)
	return err
}

func (m *UsersDAO) Update(userModel userModel.UserModel) error {
	err := db.C(collection).UpdateId(userModel.ID, userModel)
	return err
}