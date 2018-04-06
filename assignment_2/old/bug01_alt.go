package main

import ("fmt")

// I want this program to print "Hello world!", but it doesn't work.

//Alternative solution:
/*
    Having the goroutine block reading from channel,
    and have the goroutine write to the channel once its done
*/

func main() {
    ch := make(chan string)
    
    go func (){
    	ch <- "Hello world!"
    }()
	
    fmt.Println(<-ch)
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