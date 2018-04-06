package main

import (
		"fmt"
		"time"
		"sync"
)

func main(){

	fmt.Println("Main initialized..")
	
	AFP := make(chan string)
	println(AFP)
	//AFP<-"LATEST IN OMGNEWS.COM"
	var wg sync.WaitGroup
	wg.Add(1)

	
	go func(){
		fmt.Println("NEW GOROUTINE INITITALIZED")
		//println("RECIEVED",<-AFP)
		//fmt.Println(<-AFP)

		for alive:=true; alive; {
			timer := time.NewTimer(time.Hour)
			fmt.Println("LOOP")
			//AFP<-"LATEST IN OMGNEWS.COM"

			select {
				case news := <-AFP:
					fmt.Println("CASE1",news)
					timer.Stop()
					alive = false
				
				case <- time.After(time.Second):
					
					fmt.Println("CASE2", "No news in an hour.")
				/*default:
					println("DEFAULT")
				*/
			}
		}

		wg.Done()
	}()

	go func(){
		for now := range time.Tick(time.Second*2){
			fmt.Println(now, statusUpdate())
		}
	}()

	time.Sleep(5*time.Second)
	AFP<-"LATEST IN OMGNEWS.COM"

	fmt.Println("AFTER ",<-AFP)
	wg.Wait()

}

func statusUpdate() string {
	return "UPDATED STATUS"
}

