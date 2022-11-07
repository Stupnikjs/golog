package imgprocess

import (
	"fmt"
	"net/http"

	svg "github.com/ajstarks/svgo"
)

func CreateSVG(w http.ResponseWriter) {

	w.Header().Set("Content-Type", "image/svg+xml")
	s := svg.New(w)
	s.Start(500, 500)
	s.Circle(250, 20, 125, "fill:none;stroke:black")
	s.End()
	fmt.Println(s)
}
