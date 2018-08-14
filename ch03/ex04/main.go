package main

import (
	"fmt"
	"io"
	"log"
	"math"
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

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func surface(out io.Writer) {
	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", int(width), int(height))

	maxHeight, minHeight := maxMin(f)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, az := corner(i+1, j)
			bx, by, bz := corner(i, j)
			cx, cy, cz := corner(i, j+1)
			dx, dy, dz := corner(i+1, j+1)

			height := (az + bz + cz + dz) / 4
			color(height, maxHeight, minHeight)
			corners := [8]float64{ax, ay, bx, by, cx, cy, dx, dy}

			isWrongPolygon := false
			for _, corner := range corners {
				isWrongPolygon = isWrongPolygon || isInfinityOrNaN(corner)
			}
			if !isWrongPolygon {
				fmt.Fprintf(out, "<polygon points='%g,%g %g,%g %g,%g %g,%g' fill='#%06x'/>\n",
					ax, ay, bx, by, cx, cy, dx, dy, color(height, maxHeight, minHeight))
			}
		}
	}
	fmt.Fprintf(out, "</svg>")
}

func color(height, maxHeight, minHeight float64) uint {
	normalizedHeight := (height - minHeight) / (maxHeight - minHeight)
	scaledHeight := uint((1 - normalizedHeight) * 255)
	c := (red - scaledHeight<<16) + (scaledHeight)
	return c
}

func color2(height, maxHeight, minHeight float64) uint {
	normalizedHeight := (height - minHeight) / (maxHeight - minHeight)
	scaledHeight := uint(normalizedHeight * 511)
	var c uint
	if scaledHeight < 256 {
		c = blue + scaledHeight<<16
	} else {
		c = red + blue - (scaledHeight - 255)
	}
	return c
}

func isInfinityOrNaN(f float64) bool {
	return math.IsInf(f, 0) || math.IsNaN(f)
}

func corner(i, j int) (float64, float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x, y)

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z
}

func maxMin(f func(x, y float64) float64) (float64, float64) {
	max := minFloat64
	min := maxFloat64
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			x := xyrange * (float64(i)/cells - 0.5)
			y := xyrange * (float64(j)/cells - 0.5)
			z := f(x, y)

			if z > max {
				max = z
			}
			if z < min {
				min = z
			}
		}
	}
	return max, min
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
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

	if heightString, ok := r.Form["height"]; ok {
		height, _ = strconv.ParseFloat(heightString[0], 64)
	}

	xyscale = width / 2 / xyrange // pixels per x or y unit
	zscale = height * 0.4         // pixels per z unit
	angle = math.Pi / 6

	w.Header().Add("Content-Type", "image/svg+xml")
	surface(w)
}
