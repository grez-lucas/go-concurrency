package main

import (
	"fmt"
	"time"
)

func main() {
	s := goroutine("Sheep")
	c := goroutine("Cat")

	f := fanIn(s, c)

	for i := 0; i < 5; i++ {
		fmt.Println(<-f)
	}

}

// To count in LOCKSTEP
func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)

	// We redirect whatever comes in to both channels into a single channel
	go func() {
		for {
			c <- <-input1
		}
	}()
	go func() {
		for {
			c <- <-input2
		}
	}()

	return c // Return channel to the caller
}

func goroutine(message string) <-chan string {
	c := make(chan string)

	go func() {
		for i := 0; ; i++ {

			c <- fmt.Sprintf("%s %d", message, i)
			time.Sleep(time.Second * 1)
		}
	}()

	return c
}
