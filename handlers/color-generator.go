package handlers

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"net/http"

	"github.com/gorilla/mux"
)

func ParseHexColor(s string) (c color.RGBA, err error) {
	c.A = 0xff
	switch len(s) {
	case 7:
		_, err = fmt.Sscanf(s, "#%02x%02x%02x", &c.R, &c.G, &c.B)
	case 4:
		_, err = fmt.Sscanf(s, "#%1x%1x%1x", &c.R, &c.G, &c.B)
		// Double the hex digits:
		c.R *= 17
		c.G *= 17
		c.B *= 17
	default:
		err = fmt.Errorf("invalid length, must be 7 or 4")

	}
	return
}

func GenerateColor(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	size := 200

	upLeft := image.Point{0, 0}
	lowRight := image.Point{size, size}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	color, _ := ParseHexColor("#" + params["color"])

	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			img.Set(x, y, color)
		}
	}

	png.Encode(w, img)

}
