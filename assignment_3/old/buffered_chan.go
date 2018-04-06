package main
import ("fmt"
		"time")


func buffered_pinger(c chan string){
	for i:=0; i<10;i++{
		c <- "ping"
	}
	fmt.Println("pinger done")
}

func printer(c chan string){
	for {
		msg := <-c
		fmt.Println("printer received:",msg)
		//introduce a short delay to force the buffer to fill
		time.Sleep(time.Millisecond*100)
	}
}

func main() {
	// create a channel for communication
	var c chan string = make(chan string, 5)
	// create two concurrent threads
	go buffered_pinger(c)
	go printer(c)
	// runs until user input
	var input string
	fmt.Scanln(&input)
}