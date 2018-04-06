//GENERATOR - function that returns a channel

package main

import ("fmt"
		"time")

func main() {
	fmt.Println("Main initialized..")
	returnC := generator("string on channel")
	
	go func(){
		fmt.Println("INTIALIZE CLOSECHANNEL FUNCTION THAT WILL WAIT:...")
		time.Sleep(15*time.Second)
		fmt.Println("CLOSING CHANNEL, BYE...")
		close(returnC)
	}()

	for object := range returnC{
		fmt.Println("RECIEVING:", object)
	}
	
}

func generator(parameterString string) chan string{
	fmt.Println("Generator initialized..")
	ch := make(chan string)
	//initialize and declare unbuffered channel
	go func(){
		for i:=0 ; ; i++{
			fmt.Println("New goroutine initialized..")
			ch <- parameterString
			time.Sleep(5*time.Second)
		}
		
	}()

	return ch
}
