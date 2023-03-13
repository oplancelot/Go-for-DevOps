package main

import "fmt"

func main() {
	ch := make(chan string, 1)

	// Start three goroutines that read from the channel
	for i := 0; i < 3; i++ {
		go func() {
			msg := <-ch
			// fmt.Println("Received message:", i)

			fmt.Println("Received message:", msg)
		}()
	}
	// fmt.Println("Received message:")

	// Broadcast a message to all the goroutines
	// ch <- "Hello, World!"
	// fmt.Println("Received message:ch", ch)

}
