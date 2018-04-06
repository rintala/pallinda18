package main

import ("fmt")

func main() {
	fmt.Println("Main initialized...")
	ch := make(chan int)
	//ch <- 1
	go func() {
        fmt.Println("received:", <-ch)
    }()

    ch <- 2
    
    //ch <- 1

    //fmt.Println(<-ch)

	//fmt.Println(<-ch)
	//close (ch)

}