//Bug 2 -- TASK 1

package main

import ("fmt"
)

//This program should go to 11, but sometimes it only prints 1 to 10.

/******************** PROBLEM *******************************
	- The main thread returns before all values are printed
	- Thus printout of 11 is sometimes missed
*************************************************************/


/******************** SOLUTION *******************************
	- Use channel (done) to signal when Print function is done
	- Declare receiver at end of main, which will wait until value is sent
*************************************************************/

func main(){
	ch := make(chan int)
	done := make(chan bool)

	go Print(ch, done)

	for i := 1; i <= 11; i++ {
		ch <- i
	}
	
	close(ch)
	
	//Wait until a value is sent on the done channel
	//makes sure that main doesnt close before all values are printed
	<-done
	fmt.Println("WAIT IS DONE, CLOSING CHANNEL....")
}

//Print prints all numbers sent on the channel.
//The function returns when the channel is closed.
func Print(ch <-chan int, done chan <- bool){
	for n := range ch {	//reads from channel until it's closed
		fmt.Println(n)
	}

	//send true on done-channel in order to signal when the Print function is done
	done <-true

}