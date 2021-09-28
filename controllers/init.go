package controllers

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"mime/multipart"
	"tievs.com/km/db"
	"tievs.com/km/models"
	"tievs.com/km/utils"
	"time"
)

func PostFile(DB *mongo.Collection,file *multipart.FileHeader, fileName string, notes string) (interface{},error) {
	result, err := utils.PostUpload(file)
	if err != nil {
		return nil,err
	}

	var ctx, cancel = context.WithTimeout(context.Background(), 20*time.Second)
	var itemCreate models.ItemCreate
	var item models.Item
	var url models.Url
	timenow := time.Now()
	if fileName != "" {
		itemCreate.FileName = result.FileName
	}else {
		itemCreate.FileName= fileName
	}
	url.Url = result.Url
	itemCreate.Notes = notes
	itemCreate.CreatedAt = timenow
	itemCreate.UpdatedAt = timenow
	url.CreatedAt = timenow
	itemCreate.Urls = append(itemCreate.Urls,url)
	document, err := db.CreateDocument(ctx,DB, itemCreate)
	if err != nil {
		defer cancel()
		return nil,err
	}
	defer cancel()
	record,_:=json.Marshal(itemCreate)
	err = json.Unmarshal([]byte(record), &item)
	item.ID = document.(primitive.ObjectID)
	if err != nil {
		return nil,err
	}
	return item,nil
}

func PutFile(DB *mongo.Collection,file *multipart.FileHeader, fileName string, notes string, Id primitive.ObjectID) (interface{},error){
	var ctx, cancel = context.WithTimeout(context.Background(), 20*time.Second)
	timeNow := time.Now()
	result, err := utils.PutUpload(file,timeNow.String())
	if err != nil {
		defer cancel()
		return nil,err
	}
	//create URL model
	var url models.Url
	var setUpdate models.SetUpdate
	url.Url = result.Url
	url.CreatedAt = timeNow
	if fileName != "" {
		setUpdate.FileName = result.FileName
	}else {
		setUpdate.FileName= fileName
	}
	setUpdate.UpdatedAt = timeNow
	setUpdate.Notes = notes
	var objectID models.ObjectID
	objectID.ID = Id
	update := bson.M{
		"$set": setUpdate,
		"$push": bson.M{"urls":url},
	}
	_, err = db.UpdateDocument(ctx,DB,objectID, update)
	if err != nil {
		defer cancel()
		return nil,err
	}
	defer cancel()
	return url,nil
}

type DeleteMsg struct {
	Delete bool  `json:"delete" bson:"delete"`
	Errors []error `json:"errors" bson:"errors"`
	DelInterface interface{} `json:"del_interface" bson:"del_interface"`
}

func UpdateContent(DB *mongo.Collection, fileName string, notes string, Id primitive.ObjectID) (interface{}, error)  {
	var ctx, cancel = context.WithTimeout(context.Background(), 20*time.Second)
	timeNow := time.Now()
	var setUpdate models.SetUpdate
	setUpdate.UpdatedAt = timeNow
	setUpdate.Notes = notes
	setUpdate.FileName = fileName

	var objectID models.ObjectID
	objectID.ID = Id
	update := bson.M{
		"$set": setUpdate,
	}
	_, err := db.UpdateDocument(ctx,DB,objectID, update)
	if err != nil {
		defer cancel()
		return nil,err
	}
	defer cancel()
	return objectID,nil
}