package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Welcome to Race Condition Groups")

	wg := &sync.WaitGroup{}
	mut := &sync.RWMutex{}
	var scores = []int{0}

	wg.Add(4)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Routine 1")
		mut.Lock()
		scores = append(scores, 1)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	// wg.Add(1)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Routine 2")
		mut.Lock()
		scores = append(scores, 2)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	// wg.Add(1)
	go func( wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Routine 3")
		mut.Lock()
		scores = append(scores, 3)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Routine 4")
		mut.RLock()
		fmt.Println("Scores:", scores)
		mut.RUnlock()
		wg.Done()
	}(wg, mut)
	wg.Wait()
	fmt.Println("Scores:", scores)
}
