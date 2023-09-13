package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 5) // Create a channel to pass ints
	for i := 0; i < 5; i++ {
		go cookingGopher(i, c) // Start a goroutine
	}

	for i := 0; i < 5; i++ {
		gopherID := <-c // Receive a value from a channel
		fmt.Println("gopher", gopherID, "finished the dish")
	} // All goroutines are finished at this point
}

/* Notice the channel as an argument */
func cookingGopher(id int, c chan int) {
	time.Sleep(10 * time.Second)
	fmt.Println("gopher", id, "started cooking")
	c <- id // Send a value back to main
}
