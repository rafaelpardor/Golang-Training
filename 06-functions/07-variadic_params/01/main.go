package main

import "fmt"

func main() {
	xs := []int{1, 2, 3}

	fmt.Println(add(1, 2, 3))
	fmt.Println(add(xs...))
}

func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}
