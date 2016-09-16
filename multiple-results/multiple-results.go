package main

import (
	"fmt"
)

func swap(x, y string) (string, string) {
	return x + y, "again"
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
