//Many Senders; Many Receivers -- TASK 2

// Stefan Nilsson 2013-03-13

// This is a testbed to help you understand channels better.
package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

/******************** PROBLEM *******************************
	- Not all consumers have the time to finish
*************************************************************/


/******************** SOLUTION *******************************
	- Add a new WaitGroup object 'wgpc'
		- wgpc.Add(..)
		- wgpc.Wait()
		- wgpc.Done()
	- Set it up properly for the Consume function
	- Pass a pointer to the object as parameter into the Consume function
*************************************************************/

func main() {
	// Use different random numbers each time this program is executed.
	rand.Seed(time.Now().Unix())

	const strings = 32
	const producers = 4
	const consumers = 4

	before := time.Now()
	ch := make(chan string)
	wgp := new(sync.WaitGroup)

	//Declare and initialize new WaitGroup object 'wgpc'
	wgpc := new(sync.WaitGroup)

	wgp.Add(producers)
	for i := 0; i < producers; i++ {
		go Produce("p"+strconv.Itoa(i), strings/producers, ch, wgp)
	}

	//Add number of consumers (2) to wait for
	wgpc.Add(consumers)
	for i := 0; i < consumers; i++ {
		go Consume("c"+strconv.Itoa(i), ch, wgpc)
	}
	wgp.Wait() // Wait for all producers to finish.

	close(ch)
	fmt.Println("time:", time.Now().Sub(before))

	//Wait until all the consumers have sucessfully run wg.Done()
	wgpc.Wait()
}

// Produce sends n different strings on the channel and notifies wg when done.
func Produce(id string, n int, ch chan<- string, wg *sync.WaitGroup) {
	for i := 0; i < n; i++ {
		RandomSleep(100) // Simulate time to produce data.
		ch <- id + ":" + strconv.Itoa(i)
	}

	wg.Done()

}

// Consume prints strings received from the channel until the channel is closed.
func Consume(id string, ch <-chan string, wg *sync.WaitGroup) {
	for s := range ch {
		fmt.Println(id, "received", s)
		RandomSleep(100) // Simulate time to consume data.
	}

	//Communicates that the consume function is done running its processes
	wg.Done()
}

// RandomSleep waits for x ms, where x is a random number, 0 â‰¤ x < n,
// and then returns.
func RandomSleep(n int) {
	time.Sleep(time.Duration(rand.Intn(n)) * time.Millisecond)
}

/**************************** QUESTIONS & ANSWERS ***************************************/
//What happens if you switch the order of the statements wgp.Wait() and close(ch) in the end of the main function?
	/*	- We will close down the channel as soon as all goroutines have been initialized
		- This will mean we do not wait until the goroutines are done sending stuff on the channel before closing
		- Thus, we will receive a "panic: send on closed channel"
			- goroutine 5 [running]:
	*/

//What happens if you move the close(ch) from the main function and instead close the channel in the end of the function Produce?
	/*	- As soon as the first goroutine of Produce has run through until the end the channel will be closed
		- This means all the goroutines left will still try to send their values on the closed channel ch
		- Thus, we will receive a "panic: send on closed channel"
			- goroutine 6 [running]:
	*/

//What happens if you remove the statement close(ch) completely?
	/* 	- If there is no close statement of channel ch, the for loop in the Consume function will never end
		- The 'for s := range ch' will thus run forever - i.e. never get to the wg.Done() and thus the wgpc.Done() will wait forever
	*/

//What happens if you increase the number of consumers from 2 to 4?
	/*	- There will instead be 4 concurrent goroutines running to consume data (i.e. print it out from the channel)
		- Therefore the time will be faster, since number of consumers now are equal to number of producers => as fast as possible
	*/

//Can you be sure that all strings are printed before the program stops?
	/*	- Yes, since we added a waitgroup object handling the printing Consume function, with a Wait
		- This way, we make sure all strings that have been produced to the channel ch are printed out before shutdown
	/*
