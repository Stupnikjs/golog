package controllers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Stupnikjs/golog/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/securecookie"
)

func VerifyJWT(endpointHandler func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		utils.SetHeader(w, "http://localhost:3000")

		c, err := r.Cookie("token")

		if err == nil {

			var tokenString string

			s := securecookie.New([]byte(os.Getenv("SECRET_COOKIE")), nil)

			if err = s.Decode("token", c.Value, &tokenString); err == nil {

				token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("there's an error with the signing method")
					}
					return []byte(os.Getenv("SECRET_TOKEN")), nil
				})
				utils.ErrorHandler(err)
				if token.Valid {

					endpointHandler(w, r)

				} else {
					w.Write([]byte("Invalid user please try to log again"))
				}
			}

		}

	})
}
