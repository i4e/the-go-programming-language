#!/bin/bash
go build gopl.io/ch3/mandelbrot
go build main.go
./mandelbrot | ./main >mandelbrot.gif
./mandelbrot | ./main >mandelbrot.jpg
./mandelbrot | ./main >mandelbrot.png
rm mandelbrot
rm main