// ref: ßhttp://maxima.zuisei.net/pg12.html

package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.06       // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	// for i := 0; i < cells; i++ {
	// 	for j := 0; j < cells; j++ {
	// 		// ax, ay := corner(i+1, j)
	// 		// bx, by := corner(i, j)
	// 		// cx, cy := corner(i, j+1)
	// 		// dx, dy := corner(i+1, j+1)
	// 		// ax, ay := donut(i+1, j)
	// 		// bx, by := donut(i, j)
	// 		// cx, cy := donut(i, j+1)
	// 		// dx, dy := donut(i+1, j+1)
	// 		fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
	// 			ax, ay, bx, by, cx, cy, dx, dy)
	// 	}
	// }
	for i := 99; i > -1; i-- {
		for j := 99; j > -1; j-- {
			ax, ay := bane(i+1, j)
			bx, by := bane(i, j)
			cx, cy := bane(i, j+1)
			dx, dy := bane(i+1, j+1)

			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := fa(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

func tamago(x, y float64) float64 {
	return math.Sin(x) + math.Cos(y)
}

func fa(x, y float64) float64 {
	return math.Sqrt(x*x+y*y) / 5
}

func donut(i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	s := xyrange * (float64(i)/cells - 0.5)
	t := xyrange * (float64(j)/cells - 0.5)

	x := (5 + 2*math.Cos(s)) * math.Cos(t)
	y := (5 + 2*math.Cos(s)) * math.Sin(t)
	z := math.Sin(s) * 0.15

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func bane(i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	s := xyrange * (float64(i)/cells - 0.5)
	t := xyrange * (float64(j)/cells - 0.5)

	x := (5 + math.Cos(s)) * math.Cos(t)
	z := (5 + math.Cos(s)) * math.Sin(t)
	y := math.Sin(s) + 0.6*t

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}
