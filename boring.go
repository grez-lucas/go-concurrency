package main

import (
	"fmt"
	"math/rand"
	"time"
)

func boring(msg string, c chan string) {
	for i := 0; ; i++ {

		c <- fmt.Sprintf("%s %d", msg, i)
		time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	}
}

func main() {
	// Get the first 5 messages, then leave...
	c := make(chan string)
	go boring("boring", c)
	for i := 0; i < 5; i++ {
		// <- c WAITS for a value to be sent
		fmt.Printf("You say: %q\n", <-c) // Receive what was put into the channel
	}
	fmt.Println("You're boring, I'm leaving!")
}
