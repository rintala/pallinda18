//Exercise: Fibonacci Closure -- TASK 2

/************************* TASK DESCRIPTION ************************************/
	/*	- Implement a fibonacci function that returns a function (a closure)
			- The closure should return successive fibonacci numbers (0,1,1,2,3,5,..)
	*/

package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.

func fibonacci() func() int {
	prev := 0
	current := 1

	fmt.Println(prev)
	fmt.Println(current)

	return func() int {
		//save the current value in a holder, in order to set it correctly to previous (after)
		currentHolder := current
		current = current + prev
		prev = currentHolder
		return current
	}
}

func main() {
	f := fibonacci()
	max := 10
	for i := 0; i < max-2; i++ {
		fmt.Println(f())
	}
}

/**************************** QUESTIONS & ANSWERS ***************************************/
//No questions
