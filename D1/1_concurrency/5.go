package main

/*
5.1. Create two channels: one for sending even numbers and another for sending odd numbers from 1 to 20.
5.2. Use a single goroutine that sends numbers to these channels based on whether they're even or odd.
5.3. In your main function, use the select statement to read from these channels, printing whether an even or odd number was received. Stop the operation after all numbers have been printed.Output : *hasil program bisa memiliki sequence yang berbeda.
*/

import (
	"fmt"
	"sync"
)

func sendNumbers(evenCh, oddCh chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 20; i++ {
		if i%2 == 0 {
			evenCh <- i
		} else {
			oddCh <- i
		}
	}
}

func main5() {
	evenCh := make(chan int)
	oddCh := make(chan int)
	var wg sync.WaitGroup

	wg.Add(1)
	go sendNumbers(evenCh, oddCh, &wg)

	for evenCh != nil && oddCh != nil {
		select {
		case num, ok := <-evenCh:
			if ok {
				fmt.Printf("Even: %d\n", num)
			}
		case num, ok := <-oddCh:
			if ok {
				fmt.Printf("Odd: %d\n", num)
			}
		}
	}

	wg.Wait()
}
