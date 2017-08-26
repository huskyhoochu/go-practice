package main

import "fmt"

func main() {
	for i := 1; i <=10; i++ {
		if i % 2 == 0 {
			fmt.Println(i, "짝수입니다.")
		} else {
			fmt.Println(i, "홀수입니다.")
		}
	}
}