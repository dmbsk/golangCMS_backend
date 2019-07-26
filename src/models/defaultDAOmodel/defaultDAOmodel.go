package defaultDAOmodel

import "gopkg.in/mgo.v2/bson"

type DefaultDAOmodel struct {
	ID bson.ObjectId `bson:"_id" json:"id"`
}
