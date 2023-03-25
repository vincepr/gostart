package main

import (
	"context"
	"fmt"
	"time"
)
func main2(){
	ch := make(chan struct{})
	run := func(ctx context.Context){
		// we just create a channel, where we wait for the context cancel singal, 
		// then we loop n++ and print n every half second
		n := 1
		for{
			select{
			case <-ctx.Done():
				fmt.Println("recived graceful shutdown signal, shutting down")
				close(ch)
				return
			default:
				time.Sleep(time.Millisecond*500)
				fmt.Println(n)
				n++
			}
		}
	}

	// now we define our context and get our context and cancel()-function pair
	ctx, cancel := context.WithCancel(context.Background())
	// after a timer of 3 seconds we cancel our context by hand. By calling our cancel() 
	go func(){
		time.Sleep(time.Second *3)
		fmt.Println("goodbye, sending over a shutdown")
		cancel()
	}()
	go run(ctx)
	fmt.Println("only goroutines and channel left running")
	<-ch			// this blocks untill our channel is closed
	fmt.Println("main finished running")
}

// only goroutines and channel left running
// 1
// 2
// 3
// 4
// 5
// goodbye, sending over a shutdown
// 6
// recived graceful shutdown signal, shutting down
// main finished running