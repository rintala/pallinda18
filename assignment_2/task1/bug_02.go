//Bug 2 -- TASK 1

package main

import ("fmt"
		"sync"
)

//This program should go to 11, but sometimes it only prints 1 to 10.

/******************** PROBLEM *******************************
	- The main thread returns before all values are printed
	- Thus printout of 11 is sometimes missed
*************************************************************/


/******************** SOLUTION *******************************
	- Use sync libraries WaitGroup to signal when Print function is done
	- Declare wg.Wait() at end of main, which will wait until wg.Done()
*************************************************************/

func main(){
	ch := make(chan int)
	var wg sync.WaitGroup
	
	//Add one instance to wait for
	wg.Add(1)
	
	//Provide pointer to waitgroup object as parameter
	go Print(ch, &wg)
	
	for i := 1; i <= 11; i++ {
		ch <- i
	}
	
	close(ch)
	
	//Wait on wg.Done()
	wg.Wait()
		
	fmt.Println("WAIT IS DONE, CLOSING CHANNEL....")
	
}

//Print prints all numbers sent on the channel.
//The function returns when the channel is closed.
func Print(ch <-chan int, wg *sync.WaitGroup){
	for n := range ch {	//reads from channel until it's closed
		fmt.Println(n)
	}

	//Signal to waitgroup that printing process is done
	wg.Done()
	
}