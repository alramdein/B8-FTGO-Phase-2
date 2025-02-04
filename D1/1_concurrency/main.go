package main

/*

6.1. Implement a scenario where you can pass errors between goroutines using channels.
For example, when trying to send a number greater than 20 to the channels
 6.2. Handle this error in the main function by printing it out.Output : * error 21 dan 22 harus selalu di print setelah seluruh number di proses.


*/

import (
	"errors"
	"fmt"
)

func main() {
	dataCh := make(chan int)
	errCh := make(chan error)

	go func() {
		numbers := []int{5, 15, 25, 10, 30} // Sample numbers

		for _, num := range numbers {
			if num > 20 {
				errCh <- errors.New(fmt.Sprintf("error: %d is greater than 20", num))
			} else {
				dataCh <- num
			}
		}
		close(dataCh)
		close(errCh)
	}()

	// Consumer Goroutine
	for {
		select {
		case num, ok := <-dataCh:
			if !ok {
				dataCh = nil // Mark channel as nil to avoid blocking select
			} else {
				fmt.Println("Received number:", num)
			}
		case err, ok := <-errCh:
			if !ok {
				errCh = nil // Mark channel as nil to avoid blocking select
			} else {
				fmt.Println("Received error:", err)
			}
		}

		// Break when both channels are closed
		if dataCh == nil && errCh == nil {
			break
		}
	}
}
