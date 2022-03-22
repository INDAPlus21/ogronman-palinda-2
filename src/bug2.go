package main

import (
	"fmt"
	"time"
	"sync"
)

var wait sync.WaitGroup
func main() {
	ch := make(chan int)
	wait.Add(1)
	go Print(ch)
	for i := 1; i <= 11; i++ {
		ch <- i
	}
	close(ch)
	wait.Wait()
}

// Print prints all numbers sent on the channel.
// The function returns when the channel is closed.
func Print(ch <-chan int) {
	for n := range ch { // reads from channel until it's closed
		time.Sleep(10 * time.Millisecond) // simulate processing time
		fmt.Println(n)
	}
	wait.Done()
}

// This program should go to 11, but it seemingly only prints 1 to 10.
// The problem with this code is that the main function or loop finishes before the print function has time to finish
// This means that the program exits before the print function can finish
// The solution to this is to make a waitgroup that waits for all goroutines/channels to finish before actually exiting the program
func obsolete_main() {
	ch := make(chan int)
	go obsolete_Print(ch)
	for i := 1; i <= 11; i++ {
		ch <- i
	}
	close(ch)
}

// Print prints all numbers sent on the channel.
// The function returns when the channel is closed.
func obsolete_Print(ch <-chan int) {
	for n := range ch { // reads from channel until it's closed
		time.Sleep(10 * time.Millisecond) // simulate processing time
		fmt.Println(n)
	}
}