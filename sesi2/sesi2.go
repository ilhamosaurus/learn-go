package main

import "fmt"

func main() {
	tinggi := 7

	for baris := 1; baris <= tinggi; baris++ {
		spasi := tinggi - baris

		for i := 1; i <= spasi; i++ {
			fmt.Print(" ")
		}

		bintang := baris*2 - 1
		for j := 1; j <= bintang; j++ {
			fmt.Print("*")
		}

		fmt.Println()
	}
}
