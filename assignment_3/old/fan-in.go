package main

import ("fmt"
		"time"
		"math/rand"
		"sync"
)

type Message struct{
	str string
	wait chan bool
}

func main(){
	println("Main initialized..")
	
	/*
	channel1 := make(chan string)
	go boring("boringMess",channel1)
	for bor := range channel1{
		fmt.Println(bor)
	}
	*/
	
	//channel1 := make(chan string)
	//channel2 := make(chan string)

	channel1 := make(chan Message)
	channel2 := make(chan Message)

	fmt.Println("- goroutines initialized...")
	//go boring("boringmess", channel1)
	//go boring("boringtwoMess",channel2)

	fmt.Println("- fanIn of channels created..")
	outputChan := fanIn(channel1,channel2)

	//Initialize and declare wait channel to synch output
	waitForIt := make(chan bool)

	var wg sync.WaitGroup

	wg.Add(1)
	go func(){
		fmt.Println(" - GOROUTINE FOR RECEIVER INITIALIZED...")
		//for i:=0;i<5;i++{
		i:=0
		for {

			message1 := <-outputChan
			message2 := <-outputChan
			fmt.Println("OUT",i,": MESSAGE:",message1.str)
			fmt.Println("OUT",i,": MESSAGE:",message2.str)

			message1.wait <-true
			message2.wait <-true
			i++
			//fmt.Println("OUTPUT:",out)
		}
		wg.Done()
	}()

	//run boring functions
	go func(){
		for i:=0;i<10;i++{
			theString1 := fmt.Sprintf("message1: %d",i)
			theString2 := fmt.Sprintf("message2: %d",i)
			go boring(theString1, channel1, waitForIt)
			go boring(theString2, channel2, waitForIt)

			<-waitForIt
			fmt.Println("waiting for both to be done...",i)
			
			time.Sleep(time.Duration(rand.Intn(5))*time.Nanosecond)

		}
	}()

	wg.Wait()
	fmt.Println("DONE - closing down....")
	//time.Sleep(20*time.Second)
	//go boring("boringyess1", channel1)
	//go boring("boringtwoMess2",channel2)
	
	//waitForIt := make(chan bool)
	//waitForIt <- false

	//var msg1 Message
	/*go func(){
		msg1 := Message{"hi", waitForIt}
		channel1 <- msg1
	}()*/
	
	//fmt.Println("RECEIVER:",<-channel1)
	//channel1 <-msg1

	//<-waitForIt
	
	//time.Sleep(time.Second*10)

	
	
	//fmt.Println(<-lastResortCh)
	//waitForIt := make(chan bool)
	
	//msg1 := "message"
	//b := 0
	
	//msg1Object := Message{fmt.Sprintf("%s: %d", msg1, b), waitForIt}
	//lastResortCh <- msg1Object

	/*
	for i := 0; i<5;i++{
		msg1 := <-lastResortCh
		//fmt.Println(msg1.str)
		//msg2 := <-lastResortCh
		//fmt.Println(msg2.str)

		msg1.wait <-true
		//msg2.wait <-true

		

		//lastResortCh <- Message{"msg",waitForIt}
		//lastResortCh <- Message{fmt.Sprintf("%s: %d", msg2, i), waitForIt}
		time.Sleep(5*time.Second)
		<-waitForIt
	}
	*/

	/*
	for element := range lastResortCh{
		fmt.Println(element)
	}
	*/
	
}


//func boring(msg string, ch chan string){
func boring(msg string, ch chan Message, waitCh chan bool){
	//for i:=0; ;i++{
		//waitForIt := make(chan bool)
		//msg1 := Message{fmt.Sprintf("%s", msg), waitForIt}
		msg1 := Message{msg, waitCh}
		//msg2 := Message{"hi", waitCh}
		//msg1.str = msg
		//msg1.wait = false
		ch <- msg1
		//ch <- msg2

		//time.Sleep(time.Second*2)
	//}
}

//func fanIn(input1 <-chan string, input2 <-chan string) <-chan string {
func fanIn(input1, input2 <-chan Message) <-chan Message{
	outputCh := make(chan Message)

	go func(){
		for{
			outputCh <- <-input1
		}
	
	}()

	go func(){
		for c2 := range input2{
			outputCh <- c2
		}
	}()

	return outputCh
}