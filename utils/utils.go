package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"
)

func ErrorHandler(e ...error) {
	for _, error := range e {
		if error != nil {
			fmt.Println(error)
		}
	}

}

func HashtoHex(t string) string {
	hasher := sha256.New()
	hashedPswrd := hasher.Sum([]byte(t))
	result := hex.EncodeToString(hashedPswrd)
	return result
}

func SetHeader(w http.ResponseWriter, url string) {
	w.Header().Set("Access-Control-Allow-Origin", url)
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "GET,HEAD,OPTIONS,POST,PUT")

	w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Access-Control-Allow-Credentials, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")

}
