package controllers

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/Stupnikjs/golog/database"
	"github.com/Stupnikjs/golog/models"
	"github.com/Stupnikjs/golog/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/securecookie"
	"go.mongodb.org/mongo-driver/bson"
)

func LogUser(w http.ResponseWriter, r *http.Request) {

	utils.SetHeader(w)

	reqBody, errBody := ioutil.ReadAll(r.Body)

	var marshall, result models.User
	var mail string
	var response models.TokenIdResponse

	errMarshal := json.Unmarshal(reqBody, &marshall)

	utils.ErrorHandler(errBody)
	utils.ErrorHandler(errMarshal)

	client, ctx, cancel, err := database.Connect("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	defer database.Close(client, ctx, cancel)

	coll := client.Database("reactgo").Collection("users")
	filter := bson.D{{Key: "email", Value: marshall.Email}}

	errFind := coll.FindOne(context.TODO(), filter).Decode(&result)

	if errFind == nil {

		hasher := sha256.New()
		hashFromReq := hex.EncodeToString(hasher.Sum([]byte(marshall.Password)))

		fmt.Printf("password %s", result.Password)
		fmt.Printf("mail %s", mail)
		fmt.Printf("hash from req %s", hashFromReq)

		if hashFromReq == result.Password {

			token := jwt.New(jwt.SigningMethodHS256)

			claims := token.Claims.(jwt.MapClaims)
			claims["exp"] = time.Now().Add(10 * time.Minute)
			claims["authorized"] = true
			claims["user"] = marshall.Name
			claims["id"] = result.Id

			response.Token = token
			response.Id = result.Id.Hex()

			signedToken, err := token.SignedString([]byte("secret"))
			utils.ErrorHandler(err)

			fmt.Println(signedToken)
			s := securecookie.New([]byte("superscret"), nil)
			encoded, err := s.Encode("token", signedToken)

			utils.ErrorHandler(err)
			cookie := &http.Cookie{
				Name:     "token",
				Value:    encoded,
				HttpOnly: true,
				Path:     "/",
			}
			http.SetCookie(w, cookie)
			json.NewEncoder(w).Encode(response)

		} else {
			w.Write([]byte(" auhtentication failed "))
		}

	}

}
