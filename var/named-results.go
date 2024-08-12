package main

import (
	"fmt"
)

func split(sum int) (y, x int) {
	x = sum * 4
	y = sum - x
	return
}

func main() {
	fmt.Println(split(2))
}
