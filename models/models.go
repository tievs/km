package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Item struct {
	ID        	primitive.ObjectID `json:"id" bson:"_id"`
	FileName	string			   `json:"file_name" bson:"file_name"`
	CreatedAt 	time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt 	time.Time          `json:"updated_at" bson:"updated_at"`
	Urls     	[]Url              `json:"urls" bson:"urls"`
	Notes 		string             `json:"notes" bson:"notes"`
}
type ItemCreate struct {
	FileName	string			   `json:"file_name" bson:"file_name"`
	CreatedAt 	time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt 	time.Time          `json:"updated_at" bson:"updated_at"`
	Urls     	[]Url          	   `json:"urls" bson:"urls"`
	Notes 		string             `json:"notes" bson:"notes"`
}

type File struct {
	FileName	string			   `json:"file_name" bson:"file_name"`
	Url     	string             `json:"url" bson:"url"`
}
type Url struct {
	Url		string 		`json:"url" bson:"url"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
}
type SetUpdate struct {
	FileName	string			   `json:"file_name" bson:"file_name"`
	UpdatedAt 	time.Time          `json:"updated_at" bson:"updated_at"`
	Notes 		string             `json:"notes" bson:"notes"`
}
type ObjectID struct {
	ID        	primitive.ObjectID `json:"id" bson:"_id"`
}