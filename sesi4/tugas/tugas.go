package main

import (
	"fmt"
	"sync"
)

func main() {
	// Membuat slice integer dari 1 hingga 10
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Menginisialisasi Waitgroup
	var wg sync.WaitGroup

	// Looping slice integer
	for _, number := range numbers {
		// Memanggil function IIFE sebagai go routine
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			fmt.Println("Number:", n)
		}(number)
	}

	// Menunggu seluruh go routine selesai
	wg.Wait()

	fmt.Println("Semua go routine telah selesai")
}
