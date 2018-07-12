package main

import (
	"fmt"
	"strconv"
	"time"
)

// pc[i] is the population count of i.
var pc [256]byte

func main() {
	a, _ := strconv.ParseUint("1101001101010101001010010100000010100101001010101000100101001", 2, 64)
	a = uint64(a)

	measureTime(PopCount, a)
	measureTime(PopCountLoop, a)
	measureTime(PopCountShift, a)
	measureTime(PopCountClear, a)
	measurTimeNTimes(PopCount, a, 100)
	measurTimeNTimes(PopCountLoop, a, 100)
	measurTimeNTimes(PopCountShift, a, 100)
	measurTimeNTimes(PopCountClear, a, 100)
}

func measureTime(fn func(uint64) int, a uint64) {
	start := time.Now()
	result := fn(a)
	time := time.Since(start).Seconds()
	fmt.Println(result, time)
}

func measurTimeNTimes(fn func(uint64) int, a uint64, n int) {
	var totalTime float64
	for i := 0; i < n; i++ {
		start := time.Now()
		fn(a)
		secs := time.Since(start).Seconds()
		totalTime += secs
	}
	fmt.Println(totalTime)
	return
}

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCountLoop(x uint64) int {
	p := 0
	for i := 0; i < 8; i++ {
		p += int(pc[byte(x>>(uint64(i)*8))])
	}
	return p
}

func PopCountShift(x uint64) int {
	p := int(x % 2)
	for i := 0; i < 64; i++ {
		x = x >> 1
		p += int(x % 2)
	}
	return p
}

func PopCountClear(x uint64) int {
	p := 0
	for x != 0 {
		x = x & (x - 1)
		p++
	}
	return p
}
