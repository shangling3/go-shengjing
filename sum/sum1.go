package main

import "fmt"

func Sum1(param ...int) int {
	sum := 0
	for i := range param {
		sum += i
	}
	return sum
}
func main() {
	fmt.Println(Sum1(1, 2, 3, 4))
}
