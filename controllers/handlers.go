package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Stupnikjs/golog/utils"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	utils.SetHeader(w, "http://localhost:3000")
	fmt.Println("ici est la ")
	// recuperer le paramettre de la requette pour recuperer info de la base de donnée

	json.NewEncoder(w).Encode([]byte("merde ca va fonctionner "))

}
