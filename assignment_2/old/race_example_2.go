package main

import "fmt"

func main() {
    i := 0
    
    //ch := make(chan int)

    go func() {
    	//i := 0
        i++ // write
        //ch<-i
    }()

    //theI := <- ch
    //theI++

    fmt.Println(i)	// concurrent read
    //fmt.Println(theI)
}