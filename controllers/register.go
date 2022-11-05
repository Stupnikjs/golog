package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Stupnikjs/golog/database"
	"github.com/Stupnikjs/golog/models"
	"github.com/Stupnikjs/golog/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {

	utils.SetHeader(w)

	marshal := models.User{}

	body, err := io.ReadAll(r.Body)
	utils.ErrorHandler(err)

	errDecode := json.Unmarshal(body, &marshal)
	utils.ErrorHandler(errDecode)

	marshal.Password = utils.HashtoHex(marshal.Password)
	mailFromReq := marshal.Email
	// Get Client, Context, CancelFunc and
	// err from connect method.
	client, ctx, cancel, err := database.Connect("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}

	coll := client.Database("reactgo").Collection("users")
	filter := bson.D{{Key: "email", Value: mailFromReq}}
	var result models.User

	err = coll.FindOne(context.TODO(), filter).Decode(&result)
	fmt.Println(err)
	if err == mongo.ErrNoDocuments {
		user := models.User{}
		user.Id = primitive.NewObjectID()
		user.Name = marshal.Name
		user.Email = marshal.Email
		user.Password = marshal.Password

		resultInsertOne, err := database.InsertOne(client, ctx, "reactgo", "users", user)
		// Release resource when the main
		// function is returned.
		utils.ErrorHandler(err)
		fmt.Printf("resultInsertOne: %v\n", resultInsertOne)

		defer database.Close(client, ctx, cancel)

		// Ping mongoDB with Ping method
		database.Ping(client, ctx)

		utils.ErrorHandler(err)

		json, errJson := json.Marshal(resultInsertOne)
		utils.ErrorHandler(errJson)
		w.Write(json)

	}

}
