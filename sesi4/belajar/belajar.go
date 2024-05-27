package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("Main execution started")

	var wg sync.WaitGroup

	go firstProcess(8, &wg)
	wg.Add(1)

	secondProcess(8)
	
	fmt.Println("No. of goroutines:", runtime.NumGoroutine())
	wg.Wait()
	// time.Sleep(2 * time.Second)
	fmt.Println("Main execution finished")
}

func firstProcess(index int, wg *sync.WaitGroup) {
	fmt.Println("First Process starting", index)
	for i := 1; i <= index; i++ {
		fmt.Println("i =", i)
	}

	fmt.Println("First Process finished", index)
	wg.Done()
}

func secondProcess(index int) {
	fmt.Println("Second Process starting", index)
	for j := 1; j <= index; j++ {
		fmt.Println("j =", j)
	}
	fmt.Println("Second Process finished", index)
}
