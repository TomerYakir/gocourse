package main

import "fmt"

func isEven(val int) string {
	if val%2 == 0 {
		return "even"
	}
	return "odd"
}

func main() {
	sl := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, val := range sl {
		fmt.Println(val, "is", isEven(val))
	}
}
