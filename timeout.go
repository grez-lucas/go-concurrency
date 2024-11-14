package main

import (
	"fmt"
	"time"
)

func main() {
	c := goroutine("Joe")
	for {
		select {
		case s := <-c:
			fmt.Println(s)
			// time.After returns a channel that blocks for the specified duration
		case <-time.After(1 * time.Second):
			fmt.Println("You're too slow.")
		}
	}
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
