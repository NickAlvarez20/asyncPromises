package main

import (
	"fmt"
	"sync"
	"time"
)


func basic(){
	fmt.Println("=== 1. Basic Goroutines + Proper Waiting ===")

	var wg sync.WaitGroup // Use sync.Waitgroup instead of time.Sleep

	wg.Add(1) // waiting for 1 goroutine
	go func(){
		defer wg.Done()
		fmt.Println("Hello from anonymous goroutine!")
		time.Sleep(2000 * time.Millisecond) // simulate work
		fmt.Println("Anonymous goroutine finished")
	} ()
	fmt.Println("Main is doing work while go routine runs...")
	wg.Wait()
	fmt.Println("Both done!\n")
}

func main() {
	// This starts a new goroutine (lightweight thread)
	basic()

	// Without a small sleep, the program might exit
	// before the goroutine has a chance to run
	time.Sleep(100 * time.Millisecond)

	fmt.Println("Hello from main!")
}

