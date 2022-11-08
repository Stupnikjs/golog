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

		c, errCookie := r.Cookie("token")

		if errCookie == nil {

			var tokenString string

			s := securecookie.New([]byte(os.Getenv("SECRET_COOKIE")), nil)

			// decode le cookie dans la valeur tokenString
			if errDecodeCookie := s.Decode("token", c.Value, &tokenString); errDecodeCookie == nil {

				// verifie le token
				token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("there's an error with the signing method")
					}

					return []byte(os.Getenv("SECRET_TOKEN")), nil

				})

				utils.ErrorHandler(err, errCookie)

				if token.Valid {

					claims, _ := token.Claims.(jwt.MapClaims)
					user := claims["user"].(string)
					id := claims["id"].(string)
					SetTokenInCookie(id, user, w)
					endpointHandler(w, r)
				} else {
					w.Write([]byte("Invalid user please try to log again"))
				}
			}
		}
	})
}
