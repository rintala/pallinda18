package main

import ("fmt"
		"sync"
)

// I want this program to print "Hello world!", but it doesn't work.
func main() {
	var waitGroup sync.WaitGroup

    ch := make(chan string)
    
    //Adding 1 goroutine to wait for
    waitGroup.Add(1)
    
    go func (){
    	fmt.Println(<-ch)
    	
    	//Declaring that the goroutine has been completed
    	waitGroup.Done()
    }()

	ch <- "Hello world!"
    
    //Blocks goroutine that is calling it until it returns
    waitGroup.Wait()
}

//OLD STRUCTURE:
	/*package main

	import "fmt"

	// I want this program to print "Hello world!", but it doesn't work.
	func main() {
	    ch := make(chan string)
	    ch <- "Hello world!"
	    fmt.Println(<-ch)
	}
*/