package main

import (
	"fmt"
	"time"
	"sync"
	//"math"
)

func Remind(text string, delay time.Duration){
	fmt.Println("\nRemind function called...")
	fmt.Println(text)
	//fmt.Println(delay)
	//c := make(chan time.Time)

	newTicker := time.NewTicker(delay)
	<-newTicker.C

	fmt.Println("NEWTICKER:",newTicker)
	
	//current := time.Now()
	//fmt.Print("Klockan är: ", current.Hour(),".", current.Minute(),".",current.Second(),": ",text,"\n")
	
    for current := range newTicker.C {
        fmt.Print("Klockan är: ", current.Hour(),".", current.Minute(),".",current.Second(),": ",text,"\n")
    }
    /*
    threeHrTimer := time.NewTicker(3*time.Second)
    eightHrTimer := time.NewTicker(8*time.Second)
    twentyfourHrTimer := time.NewTicker(24*time.Second)

    for {
    	select {
    		case <-threeHrTimer.C:
    			fmt.Println("HI THIS RUNS EVERY THREE HOURS")
    		case <-eightHrTimer.C:
    			fmt.Println("8hrs")
    		case <-twentyfourHrTimer.C:
    			fmt.Println("24hrs")
    	}
    }
    */
    /*for c := time.Tick(1 * time.Minute); ; {
    // do work here
    select {
    case <-c:
        continue
    case <-cancel:
        return
    }
}*/


    /*
    ticker := time.NewTicker(5 * time.Second)
	quit := make(chan struct{})
	fmt.Println("\nRemind function called...")
	go func() {
	    for {
	       select {
	        case <- ticker.C:
	        	fmt.Println("\nRemind function called.1..")
	            fmt.Println(text)
	        case <- quit:
	        	fmt.Println("\nRemind function called...")
	            ticker.Stop()
	            return
	        }
	    }
	 }()
	 */

	/*select{
		case delay%3 == 0:
			threeHrTimer := time.NewTimer(3*time.Second)
			<-threeHrTimer.C
			current := time.Now()
			fmt.Print("Klockan är: ", current.Hour(),".", current.Minute(),".",current.Second(),": ",text,"\n")
		

		case delay%8 == 0:
			eightHrTimer := time.NewTimer(8*time.Second)
			<-eightHrTimer.C
			current := time.Now()
			fmt.Print("Klockan är: ", current.Hour(),".", current.Minute(),".",current.Second(),": ",text,"\n")
	}
	*/

}

//use time.Duration(timeint) to convert int => time Duration

func main() {
	p := fmt.Println
	
	p("------- Alarm Clock - Main initialized -------")
	var wg sync.WaitGroup
	

	//start := time.Now()
	current := time.Now()
	fmt.Print("Klockan är: ", current.Hour(),":", current.Minute(),".",current.Second(),"\n")
	p("---------------------------------")
	//diff := current.Sub(start)
	//fmt.Println("DIFF:",diff)
	p("START TIME:",current)
	
	//add one to waitgroup object for each new goroutine started
	wg.Add(1)
	go Remind("Dags att äta",3*time.Second)
	wg.Add(1)
	go Remind("Dags att arbeta",8*time.Second)
	wg.Add(1)
	go Remind("Dags att sova",24*time.Second)
	
	wg.Wait()
	//p(currTime.Format("15:04:05 (2006-01-02)"))
	//p("TIME DIFF: ",currTime.Sub(startTime))
	

}
