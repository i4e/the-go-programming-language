package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func echo1() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func echo2() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func echo3() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

func measurTime(function func()) float64 {
	start := time.Now()
	function()
	secs := time.Since(start).Seconds()
	return secs
}

func measurTimeNTimes(function func(), n int) float64 {
	var totalTime float64
	for i := 0; i < n; i++ {
		start := time.Now()
		function()
		secs := time.Since(start).Seconds()
		totalTime += secs
	}
	return totalTime
}

func main() {
	// FIXME 11.4節 Benchmark関数 を使う
	secs1 := measurTime(echo1)
	secs2 := measurTime(echo2)
	secs3 := measurTime(echo3)
	secs1_1000times := measurTimeNTimes(echo1, 1000)
	secs2_1000times := measurTimeNTimes(echo2, 1000)
	secs3_1000times := measurTimeNTimes(echo3, 1000)
	fmt.Println("execute once")
	fmt.Println("  echo1: ", secs1, ", echo2:", secs2, ", echo3:", secs3)
	fmt.Println("execute 1000 times")
	fmt.Println("  echo1: ", secs1_1000times, ", echo2:", secs2_1000times, ", echo3:", secs3_1000times)
}
