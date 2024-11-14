package main

func fanInWithSelect(input1, input2 <-chan string) <-chan string {

	c := make(chan string)

	go func() {
		// Note how both cases, we save the string to s and pass it to the c channel
		select {
		case s := <-input1: // In case we receive from input1...
			c <- s
		case s := <-input2: // In case we receive from input2...
			c <- s
		}
	}()

	return c

}
