package main

import (
	"fmt"
	"time"
)

func main() {
	// Create channels to communicate between goroutines
	c1 := make(chan string)
	c2 := make(chan string)

	// Use a ticker to control the 10-second duration
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	// Start the Ping goroutine first to ensure "Ping" is sent first
	go func() {
		c1 <- "Ping" // Immediately send the first "Ping"
		for {
			<-ticker.C // Wait for 1 second before sending "Ping"
			c1 <- "Ping"
		}
	}()

	// Start the Pong goroutine
	go func() {
		for {
			<-ticker.C // Wait for 1 second before sending "Pong"
			c2 <- "Pong"
		}
	}()

	// Alternate between "Ping" and "Pong" for 10 seconds
	for t := 0; t <= 10; t++ {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1, t, "s")
		case msg2 := <-c2:
			fmt.Println(msg2, t, "s")
		}
	}
}
