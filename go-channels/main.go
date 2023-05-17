/*
Channels are ways in which go routines can talk to each other
reading channel: chan<-
writing channel: <-chan
*/
package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Go Channels - joy-almeida.co/")
	wg := &sync.WaitGroup{}
	// creating a buffered channel of size 2
	ch := make(chan int, 2)
	wg.Add(2)
	// writing channel
	go func(wg *sync.WaitGroup, ch chan<- int) {
		ch <- 5
		close(ch)
		wg.Done()
	}(wg, ch)

	// reading channel
	go func(wg *sync.WaitGroup, ch <-chan int) {
		//recommended way
		x, ok := <-ch
		fmt.Println(ok)
		fmt.Println(x)

		wg.Done()
	}(wg, ch)

	wg.Wait()
}
