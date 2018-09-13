package main

import "fmt"

func main() {
	values := []int{1, 2, 3, 4}
	maxValue, err := max(values...)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("max:", maxValue)

	maxValue, err = max()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("max:", maxValue)

	fmt.Println("max:", max2(values[0], values[1:]...))

	minValue, err := min(values...)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("min", minValue)

	minValue, err = min()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("min:", minValue)

	fmt.Println("min:", min2(values[0], values[1:]...))
}

func max(vals ...int) (int, error) {
	if len(vals) <= 0 {
		return 0, fmt.Errorf("no values")
	}
	maxVal := vals[0]
	for _, val := range vals[1:] {
		if val > maxVal {
			maxVal = val
		}
	}
	return maxVal, nil
}

func min(vals ...int) (int, error) {
	if len(vals) <= 0 {
		return 0, fmt.Errorf("no values")
	}
	minVal := vals[0]
	for _, val := range vals[1:] {
		if val < minVal {
			minVal = val
		}
	}
	return minVal, nil
}

func max2(f int, vals ...int) int {
	maxVal := f
	for _, val := range vals[1:] {
		if val > maxVal {
			maxVal = val
		}
	}
	return maxVal
}

func min2(f int, vals ...int) int {
	minVal := f
	for _, val := range vals[1:] {
		if val < minVal {
			minVal = val
		}
	}
	return minVal
}
