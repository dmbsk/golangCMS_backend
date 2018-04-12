package userDAO

import (
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"log"

	"../../models/userModel"

)

const (
	COLLECTION = "users"
)

var db *mgo.Database

func (m *UsersDAO) Connect(){
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
	err := db.C(COLLECTION).Find(bson.M{}).All(&users)
	return users, err
}

func (m *UsersDAO) Insert(user userModel.UserModel) error {
	err := db.C(COLLECTION).Insert(&user)
	return err
}

func (m *UsersDAO) FindById(id string) (userModel.UserModel, error) {
	var user userModel.UserModel
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&user)
	return user, err
}

func (m *UsersDAO) Delete(userModel userModel.UserModel) error {
	err := db.C(COLLECTION).Remove(&userModel)
	return err
}

func (m *UsersDAO) Update(userModel userModel.UserModel) error {
	err := db.C(COLLECTION).UpdateId(userModel.ID, userModel)
	return err
}