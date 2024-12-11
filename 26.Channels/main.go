package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Understanding Channels")

	myChannel := make(chan int)
	wg := &sync.WaitGroup{}

	// myChannel <- 1
	// fmt.Println(<-myChannel)
	wg.Add(2)
	// receive only channel
	go func(ch <-chan int, wg *sync.WaitGroup) {
		fmt.Println("Routine 1")
		val, isChannelOpen := <-ch
		fmt.Println(val, isChannelOpen)
		fmt.Println(<-ch)
		fmt.Println(<-ch)
		wg.Done()
	}(myChannel, wg)
	// send only channel
	go func(ch chan<- int, wg *sync.WaitGroup) {
		fmt.Println("Routine 2")
		ch <- 0
		close(ch)
		// ch <- 2
		// ch <- 4
		wg.Done()
	}(myChannel, wg)
	wg.Wait()
}
