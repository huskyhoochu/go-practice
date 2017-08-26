package main

import "fmt"

func average(arg []float64) float64 {
	total := 0.0
	for _, v := range arg {
		total += v
	}
	return total / float64(len(arg))
}

func main() {
	class := []float64{98, 93, 77, 82, 83}

	fmt.Println(average(class))
}