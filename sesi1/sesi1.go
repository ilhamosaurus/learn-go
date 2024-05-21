package main

import (
	"fmt"
)

func main() {
	var name string = "Ilham"
	var age int = 28
	var isMarried bool = true

	fmt.Printf("%s\n", name)
	fmt.Printf("%d\n", age)
	fmt.Printf("%t\n", isMarried)

	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", isMarried)
}
