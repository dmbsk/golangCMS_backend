package articleModel

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type ArticleModel struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	Title       string        `bson:"title" json:"title"`
	Description string        `bson:"description" json:"description"`
	Content     string        `bson:"content" json:"content"`
	Date        time.Time     `bson:"date" json:"date"`
	Gallery     Gallery       `bson:"gallery" json:"gallery"`
	IsPublic    bool          `bson:"isPublic" json:"isPublic"`
	Author      string        `bson:"author" json:"author"`
}

type Gallery struct {
	Id          bson.ObjectId `bson:"_id" json:"id"`
	ArticleId   bson.ObjectId `bson:"_articleId" json:"articleId"`
	ImagesLinks []string      `bson:"imagesLinks" json:"imagesLinks"`
	Thumbnail   string        `bson:"thumbnailLink" json:"thumbnailLink"`
}
