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


/* Yo,
I spawn a goroutine, light it up with go,
but main keep movin’, can’t wait for the show.
So I call wg.Add(1), tell the squad “hold tight,”
then I drop the truth bomb that keeps the sync right:
defer wg.Done() – that’s the holy decree,
run this line last, no matter what happens to me.
Panic? Return? Early bail like a G?
Still get that Done(), keep the counter clean.
Forget it once, main hang forever,
program turn zombie, no pulse, no endeavor.
Call it twice? Runtime slap you with the negative counter,
panic in your face like “who the hell allowed ya?”
defer my cleanup, defer my logs,
defer file.Close(), give zero f**ks about bugs.
LIFO stack, last in first out the gate,
defer on defer, watch ’em execute in reverse fate.
I’m the goroutine closer, the grim reaper of state,
when this func die, I still seal its fate.
defer wg.Done() – that’s the oath I swear,
keep the runtime happy, no deadlock despair.
So remember the rule when you’re deep in the code:
Add then defer Done(), that’s the Go dev road.
One line, one promise, one unbreakable pact—
defer’s got your back when the function go black.
Mic drop. */