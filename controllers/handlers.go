package controllers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/Stupnikjs/golog/database"
	"github.com/Stupnikjs/golog/models"
	"github.com/Stupnikjs/golog/utils"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func IdFromRequest(r *http.Request) string {
	id := mux.Vars(r)

	return id["id"]
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	var result *models.User

	utils.SetHeader(w, "http://localhost:3000")
	id := IdFromRequest(r)

	client, ctx, cancel, err := database.Connect("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	defer database.Close(client, ctx, cancel)

	coll := client.Database("reactgo").Collection("users")

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Invalid id")
	}
	filter := bson.M{"_id": objectId}

	errFind := coll.FindOne(context.TODO(), filter).Decode(&result)
	utils.ErrorHandler(errFind)

	if errFind == nil {
		jsonResult, errJsonResult := json.Marshal(result)
		utils.ErrorHandler(errJsonResult)
		w.Write(jsonResult)
	}

}
