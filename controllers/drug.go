package controllers

import (
	"context"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"tievs.com/km/db"
	"tievs.com/km/models"
	"tievs.com/km/utils"
	"time"
)

func GetDrugFiles(c echo.Context) error {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var items [] models.Item
	DbOptions := options.Find()
	err := db.GetDocumentList(ctx,db.Drug,bson.M{}, &items, DbOptions)

	if err != nil {
		defer cancel()
		return echo.NewHTTPError(http.StatusBadRequest,err.Error())
	}
	defer cancel()
	return c.JSON(http.StatusOK, items)
}

func PostDrugFiles(c echo.Context) error {
	notes := c.FormValue("notes")
	fileName := c.FormValue("name")
	// Source
	file, err := c.FormFile("file")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest,err.Error())
	}

	drug, err := PostFile(db.Drug,file,fileName,notes)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest,err.Error())
	}
	return c.JSON(http.StatusOK, drug)
}

func UpdateDrugFiles(c echo.Context) error  {
	notes := c.FormValue("notes")
	fileName := c.FormValue("name")
	// Source
	file, err := c.FormFile("file")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest,err.Error())
	}
	Id, err := primitive.ObjectIDFromHex(c.Param("id"))
	url, err := PutFile(db.Drug,file,fileName,notes,Id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest,err.Error())
	}
	return c.JSON(http.StatusOK, url)
}

func UpdateDrugContent(c echo.Context) error  {
	notes := c.FormValue("notes")
	fileName := c.FormValue("name")
	Id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest,err.Error())
	}
	objectID, err := UpdateContent(db.Drug,fileName,notes,Id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest,err.Error())
	}
	return c.JSON(http.StatusOK,objectID)
}

func DeleteDrugFiles(c echo.Context) error {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	objectId, err := primitive.ObjectIDFromHex(c.Param("id"))
	var item models.Item
	err = db.GetDocument(ctx,db.Drug,bson.M{"_id": objectId}, &item)
	if err != nil {
		defer cancel()
		return echo.NewHTTPError(http.StatusBadRequest,err.Error())
	}

	errors := utils.Delete(&item)
	deleteV, err := db.DeleteDocument(ctx,db.Drug,bson.M{"_id": objectId})
	if err != nil {
		defer cancel()
		return echo.NewHTTPError(http.StatusBadRequest,err.Error())
	}
	var deleteMsg DeleteMsg
	if len(errors.E) != 0{
		deleteMsg.Errors = errors.E
	}
	deleteMsg.Delete = true
	deleteMsg.DelInterface = deleteV
	defer cancel()
	return c.JSON(http.StatusOK, deleteMsg)
}

func GetDrugFile(c echo.Context) error  {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	objectId, err := primitive.ObjectIDFromHex(c.Param("id"))
	var item models.Item
	err = db.GetDocument(ctx,db.Drug,bson.M{"_id": objectId}, &item)
	if err != nil {
		defer cancel()
		return echo.NewHTTPError(http.StatusBadRequest,err.Error())
	}
	defer cancel()
	return c.JSON(http.StatusOK, item)
}

