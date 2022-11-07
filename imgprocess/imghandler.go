package imgprocess

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/Stupnikjs/golog/utils"
)

func createImagePng(n int) *os.File {

	fmt.Println("calls")
	data := []int{10, 33, 73, 64}

	w, h := 500, len(data)*60+10
	r := image.Rect(0, 0, w, h)
	img := image.NewNRGBA(r)

	// tout le rectangle est blanc
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			img.Set(x, y, color.RGBA{255, 255, 255, 255})

		}
	}

	for x := 0; x < 300; x++ {
		for y := 0; y < 200; y++ {
			img.Set(x, 3*y, color.RGBA{180, 180, 250, 255})
		}

	}
	/*
		// graph
		for i, dp := range data {

			// 10, 33, 73, 64
			for x := i*60 + 10; x < (i+1)*60; x++ {

				for y := 100; y >= (100 - dp); y-- {
					img.Set(x, y, color.RGBA{180, 180, 250, 255})
				}
			}
		}
	*/

	file, err := ioutil.TempFile("images", "car-*.png")
	if err != nil {
		log.Fatal(err)
	}

	errPng := png.Encode(file, img)

	utils.ErrorHandler(err, errPng)

	return file
}

func GetImage(w http.ResponseWriter, r *http.Request) {
	utils.SetHeader(w, "http://localhost:3000")

	/*
		number := mux.Vars(r)["number"]
		numberInt, errAtoi := strconv.Atoi(number)
		utils.ErrorHandler(errAtoi)*/

	imgFile := createImagePng(46)

	bytes, err := os.ReadFile(imgFile.Name())

	b64img := base64.StdEncoding.EncodeToString(bytes)

	jsonImage, errMarshal := json.Marshal("data:image/png;base64 , " + b64img)
	utils.ErrorHandler(errMarshal, err)
	w.Write(jsonImage)
}
