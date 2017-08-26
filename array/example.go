package main

import "fmt"

var (
	start = rune(44032)
	end = rune(55204)
)

func HasConsonantSuffix(s string) bool {
	numEnds := 28
	result := false
	for _, r := range s {
		if start <= r && r < end {
			index := int(r-start)
			result = index%numEnds != 0
		}
	}
	return result
}

// func Example_array(slice []string) string {
// 	for _, value := range slice {
// 		fmt.Printf("%s는 맛있다.\n", value)
// 	}
// }

func main() {
	fruits := []string{"사과", "바나나", "멜론", "토마토", "수박"}
	for _, value := range fruits {
		result := HasConsonantSuffix(fruits)
		if result == true {
			fmt.Printf("%s은 맛있다.\n", value)
		} else {
			fmt.Printf("%s는 맛있다.\n", value)
		}
	}
}