package videoupload

import (
	"fmt"
	"io"
	"net/http"

	"github.com/Stupnikjs/golog/controllers"
	"github.com/Stupnikjs/golog/utils"
)

func PostVideo(w http.ResponseWriter, r *http.Request) {
	id := controllers.IdFromRequest(r)
	fmt.Println("coucou")
	utils.SetHeader(w, "http://localhost:3000")

	w.Header().Set("Content-Type", "video/mp4")
	body, err := io.ReadAll(r.Body)

	fmt.Println(body, err, id)

}
