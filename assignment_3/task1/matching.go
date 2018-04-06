//Matching Behaviour -- TASK 1

//http://www.nada.kth.se/~snilsson/concurrency/

package main

import (
	"fmt"
	"sync"
	"time"
)

// This programs demonstrates how a channel can be used for sending and
// receiving by any number of goroutines. It also shows how  the select
// statement can be used to choose one out of several communications.
func main() {
	people := []string{"Anna", "Bob", "Cody", "Dave", "Eva"}


	match := make(chan string,1) // Make room for one unmatched send.

	wg := new(sync.WaitGroup)
	//var wg sync.WaitGroup

	wg.Add(len(people))

	start := time.Now()

	for _, name := range people {
		go Seek(name, match, wg)
	}

	fmt.Println("Time:", time.Now().Sub(start))

	wg.Wait()

	//TAKES CARE OF UNEVEN NO OF PEOPLE
	select {
		//receiver created
		case name := <-match:
			fmt.Printf("No one received %s's message.\n", name)
		/*
		default:
			// There was no pending send operation.
			fmt.Println("DEFAULT")
		*/
	}
}

// Seek either sends or receives, whichever possible, a name on the match
// channel and notifies the wait group when done.
func Seek(name string, match chan string, /*wg sync.WaitGroup*/ wg *sync.WaitGroup) {
	select {
		case peer := <-match:
			fmt.Printf("%s sent a message to %s.\n", peer, name)
		
		case match <- name:
			// Wait for someone to receive my message.
			fmt.Printf("Waiting for receiver (adding name %s to match-channel)\n",name)
	}

	wg.Done()
}

/************************* CONCLUSIONS & OBSERVATIONS ************************************/
	/*


	*/

/**************************** QUESTIONS & ANSWERS ***************************************/
/*Hints & instructions:
		- Think about the order of the instructions and what happens with arrays of different lengths.
		- Explain what happens and why it happens if you make the following changes.
		- Try first to reason about it, and then test your hypothesis by changing and running the program.
*/

//What happens if you remove the go-command from the Seek call in the main function?
	/* 	- The function will not run in its own goroutine - i.e. one goroutine per name
		  and thus will the 5 instances of the Seek function not run concurrently
			- It will instead run in the same order each time, in a linear manner
	      prompting the same results and order every time it is being run

	*/

//What happens if you switch the declaration wg := new(sync.WaitGroup) to var wg sync.WaitGroup and the parameter wg *sync.WaitGroup to wg sync.WaitGroup?
	/*	- 'fatal error: all goroutines are asleep - deadlock!'
			- The parameter 'wg' will not be sent in as a pointer, rather a copy of the object itself. Since
		  several goroutines are being run, they use different memory spaces and the WaitGroup
          objects will be different and thus not work in the expected way - the WaitGroup.Done() will not return to outer wg object

	*/

//What happens if you remove the buffer on the channel match?
	/* 	- The channel 'match' will then block until
		- "Fatal error, all goroutines are asleep - deadlock!"
		-  'No one received Anna's message.' will never run since the second sending case match<-name
			will wait until receiver is active if no buffer exists - and the receiver will not be run until the case finished...
	*/

//What happens if you remove the default-case from the case-statement in the main function?
	/*		- The select statement in the main "takes care" of the last person
			- IF the no. of persons is uneven - the default will not be needed since the first case takes care of it
			- IF even no. of persons - the default is necessary to exit the select-statement - otherwise it will get stuck on blocking..  
	*/
