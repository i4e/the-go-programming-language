package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"math/cmplx"
	"net/http"
	"strconv"
)

const (
	cells = 100 // number of grid cells

	maxFloat64 = 1.797693134862315708145274237317043567981e+308 // 2**1023 * (2**53 - 1) / 2**52
	minFloat64 = 4.940656458412465441765687928682213723651e-324 // 1 / 2**(1023 - 1 + 52)

	red  = 0x00ff0000
	blue = 0x000000ff
)

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

var width, height = 600.0, 320.0           // canvas size in pixels
var xyrange = 30.0                         // axis ranges (-xyrange..+xyrange)
var xyscale = float64(width) / 2 / xyrange // pixels per x or y unit
var zscale = float64(height) * 0.4         // pixels per z unit
var angle = math.Pi / 6                    // angle of x, y axes (=30°)

func handler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	if widthString, ok := r.Form["width"]; ok {
		width, _ = strconv.ParseFloat(widthString[0], 64)
	}

	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}

	err := png.Encode(w, img)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
