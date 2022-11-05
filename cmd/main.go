package main

// https://github.com/olliefr/docker-gs-ping to learn docker with golang

import (
	"net/http"

	"os"

	"github.com/Stupnikjs/golog/controllers"
	"github.com/gorilla/mux"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}

	r := mux.NewRouter()
	r.HandleFunc("/post", controllers.RegisterUser)

	r.HandleFunc("/login", controllers.LogUser)
	r.HandleFunc("/profile/{id}", controllers.VerifyJWT(controllers.GetUser))

	http.ListenAndServe(":"+port, r)

}
