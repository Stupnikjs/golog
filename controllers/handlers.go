package controllers

import (
	"fmt"
	"net/http"

	"github.com/Stupnikjs/golog/utils"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	utils.SetHeader(w)
	fmt.Println("ici")
	// recuperer le paramettre de la requette pour recuperer info de la base de donnée

}
