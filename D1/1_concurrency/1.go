package main

/*
1.1. Write a function printNumbers that prints numbers from 1 to 10.
1.2. Write a function printLetters that prints letters from 'a' to 'j'.
1.3. Use goroutines to concurrently run both functions.
*/

import (
	"fmt"
	"sync"
)

func printNumbers() {
	for i := 1; i <= 100; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
}

func printLetters() {
	for ch := 'a'; ch <= 'j'; ch++ {
		fmt.Printf("%c ", ch)
	}
	fmt.Println()
}

func main1() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		printNumbers()
	}()

	go func() {
		defer wg.Done()
		printLetters()
	}()

	wg.Wait()
	// time.Sleep(100 * time.Microsecond)
}
