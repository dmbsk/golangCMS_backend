package dbDAO

import (
	"log"

	"gopkg.in/mgo.v2"
)

func (m *DAO) Connect(sessionName string) *mgo.Database {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	} else {
		println( sessionName + " started without error")
	}
	return session.DB(m.Database)
}


//func (m *DAO) FindAll() ([]struct{}, error){
//	var users []struct{}
//	err := db.C(COLLECTION).Find(bson.M{}).All(&users)
//	return users, err
//}
//
//func (m *DAO) Insert(user struct{}) error {
//	err := db.C(COLLECTION).Insert(&user)
//	return err
//}
//
//func (m *DAO) FindById(id string) (struct{}, error) {
//	var user struct{}
//	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&user)
//	return user, err
//}
//
//func (m *DAO) Delete(model struct{}) error {
//	err := db.C(COLLECTION).Remove(&model)
//	return err
//}
//
//func (m *DAO) Update(model struct{}) error {
//	type defaultModel struct {
//		ID bson.ObjectId `bson:"_id" json:"id"`
//		model struct{}
//	}
//	err := db.C(COLLECTION).UpdateId(defaultModel.ID, model)
//	return err
//}