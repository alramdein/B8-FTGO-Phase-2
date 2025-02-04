package main

/*
	3.1. Create a function produce that sends numbers from 1 to 10 on a channel.
	3.2. Create a function consume that reads from the channel and prints the numbers.
	3.3. Use goroutines to concurrently run both functions and demonstrate communication between them using the channel.
*/

import (
	"fmt"
	"sync"
	"time"
)

func produce(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 100; i++ {
		ch <- i
	}
	close(ch)
}

func consume(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for _ = range ch {
		// fmt.Printf("%d ", num)
	}
	fmt.Println()
}

func main4() {
	var wg sync.WaitGroup
	ch := make(chan int)

	start := time.Now()
	wg.Add(2)
	go produce(ch, &wg)
	go consume(ch, &wg)
	end := time.Now()

	wg.Wait()

	fmt.Println("diff: ", end.Sub(start))
}
