package main

import ("fmt")
func race() {
    ch := make(chan int)

    go func() {
    	n := 0	//local variable that is only visible to this goroutine
        n ++ // read, increment, write
        ch <- n //the data leaves one goroutine
    }()
    n := <-ch //.. and arrives safely in another
    n++ // conflicting access
    //<-wait
    fmt.Println(n) // Output: 2
}

func main(){
	race()
}