package controllers

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/Stupnikjs/golog/database"
	"github.com/Stupnikjs/golog/models"
	"github.com/Stupnikjs/golog/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/securecookie"
	"go.mongodb.org/mongo-driver/bson"
)

func getToken(result *models.User) string {

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(10 * time.Minute)
	claims["authorized"] = true
	claims["user"] = result.Name
	claims["id"] = result.Id

	signedToken, err := token.SignedString([]byte(os.Getenv("SECRET_TOKEN")))
	utils.ErrorHandler(err)
	return signedToken

}

func getCoookie(content string) (*http.Cookie, error) {
	s := securecookie.New([]byte(os.Getenv("SECRET_COOKIE")), nil)

	encoded, err := s.Encode("token", content)

	cookie := &http.Cookie{
		Name:     "token",
		Value:    encoded,
		HttpOnly: true,
		Path:     "/",
		Secure:   false,
	}
	return cookie, err
}

func LogUser(w http.ResponseWriter, r *http.Request) {
	var marshall, result models.User

	utils.SetHeader(w, "http://localhost:3000")
	reqBody, errBody := ioutil.ReadAll(r.Body)

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

		if hashFromReq == result.Password {

			signedToken := getToken(&result)

			fmt.Println(signedToken)

			cookie, errCookie := getCoookie(signedToken)
			utils.ErrorHandler(errCookie)

			http.SetCookie(w, cookie)
			json.NewEncoder(w).Encode(result.Id)

		} else {
			w.Write([]byte(" auhtentication failed "))
		}

	}

}
