//Bug 1 -- TASK 1 

package main

import ("fmt")

//I want this program to print "hello world!", but it doesnt work
func main(){
	ch := make(chan string)
	
	//start goroutine for sender
	go func(){
		ch <- "Hello world!"	
	}()

	//<-ch blocks until sender done
	fmt.Println(<-ch)
}

//------ BEFORE_FUNCTION -----------
/*
func main(){
	ch := make(chan string)
	ch <- "Hello world!"
	fmt.Println(<-ch)
}
*/
//----------------------------------