package main

import ("fmt"
		"sync"
)


// This program should go to 11, but sometimes it only prints 1 to 10.
//Reason: The program and it's main exits before the Print routine is complete
//Could be resolved in an ugly way by for example setting a bunch of prints before

func main() {

    //waitGroup := new(sync.WaitGroup)
    var waitGroup sync.WaitGroup

    ch := make(chan int)
    
    waitGroup.Add(1)

    //have to send the memory adress of the waitGroup object since a new goroutine is started
    //in order for the correct waitGroup to be called in there

    go Print(ch, &waitGroup)  
    
	for i := 1; i <= 11; i++ {
		ch <- i
	}

	close(ch)
	waitGroup.Wait()
    
}

// Print prints all numbers sent on the channel.
// The function returns when the channel is closed.
func Print(ch <-chan int, wg *sync.WaitGroup) {
	//var waitGroup sync.WaitGroup
    for n := range ch { // reads from channel until it's closed
        fmt.Println(n)
    }
    wg.Done()

}