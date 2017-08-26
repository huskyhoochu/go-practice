package main

import "fmt"

func main() {
	fmt.Print("화씨를 입력하세요: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := (input - 32) * 5/9
	fmt.Println("섭씨 ", output, "도")
}