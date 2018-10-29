package main

import "fmt"

func main() {
	fmt.Println(operation(17))
}

func operation(sum int) (x, y, z, a int) {
	x = sum * 4 / 9
	y = sum - x
	z = sum * sum
	a = x + y + sum
	return
}

/*
Naked return statements should be used only in short functions,
as with the example shown here.
They can harm readability in longer functions.
*/