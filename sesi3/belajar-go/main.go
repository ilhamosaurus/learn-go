package main

import "belajar-go/helpers"

func main() {
	helpers.Greet()

	person := helpers.Person{}

	person.InvokeGreet()
}
