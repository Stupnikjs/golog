package controllers

import (
	"fmt"
	"net/http"

	"github.com/Stupnikjs/golog/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/securecookie"
)

func VerifyJWT(endpointHandler func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("token")
		if err != nil {
			fmt.Printf("erreur dans la lecture du cookie %s", err)
		}
		var tokenString string

		s := securecookie.New([]byte("superscret"), nil)

		if err = s.Decode("token", c.Value, &tokenString); err == nil {
			fmt.Fprintln(w, tokenString)
			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("there's an error with the signing method")
				}
				return []byte("secret"), nil
			})
			utils.ErrorHandler(err)
			if token.Valid {
				fmt.Println(token)
				endpointHandler(w, r)

			}
		}
		/*
				if r.Header["Token"] != nil {

					token, errParseToken := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
						fmt.Println(token)
						_, ok := token.Method.(*jwt.SigningMethodECDSA)
						if !ok {
							w.WriteHeader(http.StatusUnauthorized)
							_, err := w.Write([]byte("You're Unauthorized!"))
							if err != nil {
								return nil, err

							}
						}
						return "", nil

					})

					if errParseToken != nil {
						w.WriteHeader(http.StatusUnauthorized)
						w.Write([]byte("You're Unauthorized due to error parsing the JWT"))

					}

					if token.Valid {
						endpointHandler(w, r)
					} else {
						w.WriteHeader(http.StatusUnauthorized)
						_, err := w.Write([]byte("You're Unauthorized due to invalid token"))
						if err != nil {
							return
						}
					}
				}

			})
		*/

	})
}
