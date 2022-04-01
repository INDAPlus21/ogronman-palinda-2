package main

import "fmt"

func main() {
	ch := make(chan string)
	go func() { // Now the codes goes to the function with a goroutine
		ch <- "Hello world!"
		close(ch)
	}()
	fmt.Println(<-ch) //The print function waits for the channel to recieve information, which it gets parallell in the func
}

// I want this program to print "Hello world!", but it doesn't work.
/*
	The problem with this code is that the channel needs a goroutine to continue/receive it, which means that the channel is waiting for something that never happens
*/
func obsolete_main() {
	ch := make(chan string)
	ch <- "Hello world!"
	fmt.Println(<-ch)
}
