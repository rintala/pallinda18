//Alarm Clock -- TASK 3
	//Solution with three different concurrent goroutines running and sending to a channel

package main

import (
	"fmt"
	"time"
)

func Remind(text string, delay time.Duration, outCh chan string){
		fmt.Println("Remind function called...\n")
		//NewTicker returns a new Ticker containing a channel that will send the
		//time with a period specified by the duration argument.
    newTimer := time.NewTicker(delay)

		//Run for-loop with select to identify what case/timer channel and return appropriate strings
		for{
			select{
				case <-newTimer.C:
					current := time.Now()
					currentString := fmt.Sprintf("Klockan är: "+"%d.%d.%d"+": "+"%s"+"\n", current.Hour(), current.Minute(), current.Second(), text)
					outCh<-currentString
			}
		}
}

func main() {
	p := fmt.Println
	outCh := make(chan string)

	p("------- Alarm Clock - Main initialized -------")
	start := time.Now()
	p("START TIME: ", start.Format("15:04:05 (2006-01-02)"))
	p("---------------------------------")

	go Remind("Dags att äta", 3*time.Second, outCh)
	go Remind("Dags att arbeta", 8*time.Second, outCh)
	go Remind("Dags att sova", 24*time.Second, outCh)

	for outp := range outCh{
		fmt.Print(outp)
	}

}
