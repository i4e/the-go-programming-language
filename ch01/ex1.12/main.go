package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
)

//!+main
var green = color.RGBA{0x00, 0x80, 0x80, 0xff}
var red = color.RGBA{0xff, 0x00, 0x00, 0xff}
var blue = color.RGBA{0x00, 0x00, 0xff, 0xff}
var palette = []color.Color{color.White, color.Black, green, red, blue}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
	greenIndex = 2
	redIndex   = 3
	blueIndex  = 4
)

func main() {
	// handler := func(w http.ResponseWriter, r *http.Request) {
	// 	lissajous(w)
	// }
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

//!+handler
// handler echoes the HTTP request.
func handler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	// for k, v := range r.Header {
	// 	fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	// }
	// fmt.Fprintf(w, "Host = %q\n", r.Host)
	// fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	cycles := 5   // number of complete x oscillator revolutions
	res := 0.001  // angular resolution
	size := 100   // image canvas covers [-size..+size]
	nframes := 64 // number of animation frames
	delay := 8    // delay between frames in 10ms units

	if cyclesString, ok := r.Form["cycles"]; ok {
		cycles, _ = strconv.Atoi(cyclesString[0])
	}

	if resString, ok := r.Form["res"]; ok {
		res, _ = strconv.ParseFloat(resString[0], 64)
	}

	if sizeString, ok := r.Form["size"]; ok {
		size, _ = strconv.Atoi(sizeString[0])
	}

	if nframesString, ok := r.Form["nframes"]; ok {
		nframes, _ = strconv.Atoi(nframesString[0])
	}

	if delayString, ok := r.Form["delay"]; ok {
		delay, _ = strconv.Atoi(delayString[0])
	}

	lissajous(w, cycles, res, size, nframes, delay)
	// lissajous(w)
}

func lissajous(out io.Writer, cycles int, res float64, size int, nframes int, delay int) {
	// const (
	// 	cycles  = 5     // number of complete x oscillator revolutions
	// 	res     = 0.001 // angular resolution
	// 	size    = 100   // image canvas covers [-size..+size]
	// 	nframes = 64    // number of animation frames
	// 	delay   = 8     // delay between frames in 10ms units
	// )

	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		// var r uint8 = 0x00
		// var g uint8 = 0x80
		// var b uint8 = 0xff
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5),
				uint8(int(t*10)*i%4+1))
			// img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5),
			// 	uint8(int(t*10)*i%3+2))
			// if r == 0xff {
			// 	r = 0x00
			// }
			// if g == 0xff {
			// 	g = 0x00
			// }
			// if b == 0x00 {
			// 	g = 0xff
			// }
			// r += 0x01
			// g += 0x40
			// b -= 0x20
			// img.Set(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5),
			// 	color.RGBA{g, b, r, 0xff})
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
