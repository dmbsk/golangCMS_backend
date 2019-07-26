package userModel

import (
	"gopkg.in/mgo.v2/bson"
)

type UserModel struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	Username    string        `bson:"username" json:"username"`
	Password    string        `bson:"password" json:"password"`
	Description string        `bson:"description" json:"description"`
	Role        string        `bson:"role" json:"role"`
}
