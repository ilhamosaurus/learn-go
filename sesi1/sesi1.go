package main

import (
	"fmt"
)

func main() {
	var name string = "Ilham"
	var age int = 42
	var isIntrovert bool = true

	fmt.Printf("%s\n", name)
	fmt.Printf("%d\n", age)
	fmt.Printf("%t\n", isIntrovert)

	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", isIntrovert)
}
