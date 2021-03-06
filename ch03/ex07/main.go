package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
		dx                     = (xmax - xmin) / width
		dy                     = (ymax - ymin) / height
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			c := superSample(x, y, dx, dy)
			img.Set(px, py, c)
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func superSample(x, y, dx, dy float64) color.Color {
	subPixels := []color.Color{
		newton(complex(x-dx, y-dy)),
		newton(complex(x-dx, y+dy)),
		newton(complex(x+dx, y-dy)),
		newton(complex(x+dx, y+dy))}

	return avg(subPixels)
}

func avg(colors []color.Color) color.Color {
	var r, g, b, a uint16
	n := len(colors)
	for _, c := range colors {
		r_, g_, b_, a_ := c.RGBA()
		r += uint16(r_ / uint32(n))
		g += uint16(g_ / uint32(n))
		b += uint16(b_ / uint32(n))
		a += uint16(a_ / uint32(n))
	}
	return color.RGBA64{r, g, b, a}
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			r := contrast * n
			g := 127 - contrast*n/2
			b := 255 - contrast*n
			return color.RGBA{r, g, b, 255}
		}
	}
	return color.Black
}

//!-

// Some other interesting functions:

func acos(z complex128) color.Color {
	v := cmplx.Acos(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{192, blue, red}
}

func sqrt(z complex128) color.Color {
	v := cmplx.Sqrt(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{128, blue, red}
}

// f(x) = x^4 - 1
//
// z' = z - f(z)/f'(z)
//    = z - (z^4 - 1) / (4 * z^3)
//    = z - (z - 1/z^3) / 4
func newton(z complex128) color.Color {
	const iterations = 37
	const contrast = 7
	for i := uint8(0); i < iterations; i++ {
		z -= (z - 1/(z*z*z)) / 4

		switch {
		case cmplx.Abs((1+0i)-z) < 1e-6:
			return color.RGBA{0, 0, contrast * i, 0xff}
		case cmplx.Abs((-1+0i)-z) < 1e-6:
			return color.RGBA{0, contrast * i, 0, 0xff}
		case cmplx.Abs((0+1i)-z) < 1e-6:
			return color.RGBA{contrast * i, 0, 0, 0xff}
		case cmplx.Abs((0-1i)-z) < 1e-6:
			return color.RGBA{contrast * i, contrast * i, contrast * i, 0xff}
		}
	}
	return color.Black
}
